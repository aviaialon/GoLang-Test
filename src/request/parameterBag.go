package request

import (
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
var _otherSiteId 	int
var _currency 		string

type ParameterBag struct {
	SiteId        	int    	`json:"site_id"`
	OtherSiteId    	int	`json:"other_site_id"`
	Currency    	string	`json:"currency"`
}

/**
 * Constructor
 */
func NewParameterBag(r *http.Request) *ParameterBag {
	_request      = r
	_siteId       = getSiteId();
	_otherSiteId  = getOtherSiteId();
	_currency     = getCurrency();
	_parameterBag = &ParameterBag{
		SiteId: 	_siteId,
		OtherSiteId:	_otherSiteId,
		Currency: 	_currency,
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

	return fmt.Sprintf("%v", value)
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

	return strings.TrimSpace(strings.ToUpper(map[bool]string{true: currency, false: CURRENCY_CODE_CAD} [currency == CURRENCY_CODE_USD]))
}
