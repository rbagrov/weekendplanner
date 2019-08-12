package helpers

import (
	sql "database/sql"
	log "log"
)

// GenericEvent struct that should be used from dbwrapper
type GenericEvent struct {
	Date    string
	Event   string
	PoiName string
}

// GenericPOIEvent struct that should be used from dbwrapper
type GenericPOIEvent struct {
	Day     int
	Month   string
	Year    int
	Event   string
	PoiName string
}

// CheckErr : Logs fatal if error exists
func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// SQLCheckErr : Checks if specific sql error has been raised
func SQLCheckErr(err error) {
	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
	}
}

// StatusCodeChecker : Checks http response code
func StatusCodeChecker(code int, url string) {
	if code != 200 {
		log.Fatal("Code %d error for: %s", code, url)
	}
}
