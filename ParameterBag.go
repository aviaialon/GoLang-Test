package main

import (
	"log"
	"net/http"
	"strconv"
	"reflect"
	"github.com/gorilla/mux"
)

const SITE_ID_JUSTFLY   = 1
const SITE_ID_FLIGHTHUB = 4

var _request *http.Request

type ParameterBag struct {
	Test			string `Avi`
	SiteId        	int
	OtherSiteId    	int
}

func NewParameterBag(r *http.Request) *ParameterBag {
	_request = r
	return &ParameterBag{
		OtherSiteId: getSiteId(),
		SiteId: getSiteId(),
	}
}

func getSiteId() (_siteId int) {
	_siteId, err := strconv.Atoi(mux.Vars(_request)["site_id"]) // Seems like the third argument is not honored? i always get an int64 dafuq??
log.Print(mux.Vars(_request)["site_id"])
	if (err != nil || _siteId == 0) {
		_siteId = SITE_ID_JUSTFLY;
	}

	exists, _ := inArray(_siteId, []int{SITE_ID_JUSTFLY, SITE_ID_FLIGHTHUB})
	if (exists == false) {
		_siteId = SITE_ID_JUSTFLY;
	}

	return _siteId
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