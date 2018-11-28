package request

import (
	"utils/uniqid"
	"crypto/md5"
	"encoding/hex"
	"net/http"
	"strings"
	"strconv"
	"utils"
	"fmt"
)

const SITE_ID_JUSTFLY   = 1
const SITE_ID_FLIGHTHUB = 4
const CURRENCY_CODE_USD = "USD"
const CURRENCY_CODE_CAD = "CAD"

var _parameterBag       *ParameterBag
var _request 		*http.Request
var _siteId 		int
var _query 		string
var _currency 		string

type ParameterBag struct {
	SiteId        	int    	`json:"site_id"`
	OtherSiteId    	int	`json:"other_site_id"`
	Currency    	string	`json:"currency"`
	SearchString  	string	`json:"search_query"`
	SearchId    	string	`json:"search_id"`
}

/**
 * Constructor
 */
func NewFromRequest(r *http.Request) *ParameterBag {
	_request      = r
	_siteId       = getSiteId()
	_currency     = getCurrency();

	loadQueryString()

	_parameterBag = &ParameterBag{
		SiteId: 	_siteId,
		Currency: 	_currency,
		SearchString:   _query,
		OtherSiteId:	getOtherSiteId(),
		SearchId: 	getSearchId(),
	}

	return _parameterBag
}

/**
 * Get all parameters
 * @return ParameterBag
 */
func All() (*ParameterBag) {
	return _parameterBag
}

/**
 * Get a request param
 *
 * @param string - request variable name
 * @param string (optional) - default value
 * @return string
 */
func getRequestParam(paramName, defaultValue string) (string) {
	value := strings.Join(_request.URL.Query()[paramName], "")

	if (value == "" && defaultValue != "") {
		value = defaultValue
	}

	return strings.TrimSpace(strings.ToUpper(fmt.Sprintf("%v", value)))
}

/**
 * Get the site id
 *
 * @return int
 */
func getSiteId() (_siteId int) {
	_siteId, err := strconv.Atoi(getRequestParam("site_id", ""))

	if (err != nil || _siteId == 0) {
		_siteId = SITE_ID_JUSTFLY;
	}

	exists, _ := utils.InArray(_siteId, []int{SITE_ID_JUSTFLY, SITE_ID_FLIGHTHUB})
	if (exists == false) {
		_siteId = SITE_ID_JUSTFLY;
	}

	return _siteId
}
/**
 * Get Other site id
 *
 * @return int
 */
func getOtherSiteId() (int) {
	return map[bool]int{true: SITE_ID_FLIGHTHUB, false: SITE_ID_JUSTFLY} [_siteId == SITE_ID_JUSTFLY]
}

/**
 * Get the request currency
 *
 * @return string
 */
func getCurrency() (string) {
	var currency string = getRequestParam("currency", CURRENCY_CODE_CAD);

	return map[bool]string{true: currency, false: CURRENCY_CODE_CAD} [currency == CURRENCY_CODE_USD]
}

func loadQueryString()  {
	var prefix, getParam  		 string   = "seg%v_", ""
	var segIndex          		 int      = 0
	var segmentValueIndexes 	[]string  = []string{"to", "from", "date", "time"}
	var searchParamIndexes	 	[]string  = []string{"num_adults", "num_children", "num_infants", "num_infants_lap", "preferred_carrier_code", "seat_class", "flexible_date", "nearby_airports", "non_stop", "no_penalties", "target_id"}
	var _params 			[]string

	// Add segment params
	for (segIndex <= 10) {
		if (getRequestParam(fmt.Sprintf(prefix + "from", segIndex), "-1") != "-1") {
			for i := 0; i < len(segmentValueIndexes); i++ {
				getParam = fmt.Sprintf(prefix + "%v", segIndex, segmentValueIndexes[i])
				_params = append(_params, fmt.Sprintf("%v=%v&", getParam, getRequestParam(getParam, "")))
			}
		}

		segIndex += 1
	}

	// Add other params
	for i := 0; i < len(searchParamIndexes); i++ {
		getParam = fmt.Sprintf(prefix + "%v", segIndex, searchParamIndexes[i])
		_params = append(_params, fmt.Sprintf("%v=%v&", getParam, getRequestParam(getParam, "")))
	}

	// Add static params
	_params = append(_params, fmt.Sprintf("site_id=%v&", _siteId))
	_params = append(_params, fmt.Sprintf("currency=%v&", _currency))
	_params = append(_params, fmt.Sprintf("uniqid=%v&", uniqid.New(uniqid.Params{"", false})))

	_query = strings.Join(_params, "")
}

/**
 * Generates the search id
 *
 * @see \Mv_Ota_Fare_Package_Search::getSearchId
 * @return string
 */
func getSearchId() (string) {
	return md5Value("fare_fetching_packages:" + strings.ToLower(_query))
}

func md5Value(argData string) (string) {
	hasher := md5.New()
	hasher.Write([]byte(argData))

	return hex.EncodeToString(hasher.Sum(nil))
}