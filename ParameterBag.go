package main

import (
	"net/http"
	"strings"
	"strconv"
	"reflect"
)

const SITE_ID_JUSTFLY   = 1
const SITE_ID_FLIGHTHUB = 4

var _request *http.Request
var _siteId int

type ParameterBag struct {
	SiteId        	int    	`json:"site_id"`
	OtherSiteId    	int	`json:"other_site_id"`
}

func NewParameterBag(r *http.Request) *ParameterBag {
	_request = r
	_siteId  = getSiteId();

	return &ParameterBag{
		SiteId: _siteId,
		OtherSiteId: getOtherSiteId(),
	}
}

func getSiteId() (_siteId int) {
	_siteId, err := strconv.Atoi(strings.Join(_request.URL.Query()["site_id"], ""))

	if (err != nil || _siteId == 0) {
		_siteId = SITE_ID_JUSTFLY;
	}

	exists, _ := inArray(_siteId, []int{SITE_ID_JUSTFLY, SITE_ID_FLIGHTHUB})
	if (exists == false) {
		_siteId = SITE_ID_JUSTFLY;
	}

	return _siteId
}

func getOtherSiteId() (int) {
	return map[bool]int{true: SITE_ID_FLIGHTHUB, false: SITE_ID_JUSTFLY} [_siteId == SITE_ID_JUSTFLY]
}

func inArray(needle interface{}, haystack interface{}) (exists bool, index int) {
	exists = false
	index = -1

	switch reflect.TypeOf(haystack).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(haystack)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(needle, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}

	return
}