package helpers

import (
	sql "database/sql"
	fmt "fmt"
	log "log"
	strconv "strconv"
	time "time"
)

// POI struct
type POI struct {
	ID int
}

// GenericScraperEvent struct used in scrapers
type GenericScraperEvent struct {
	Date  time.Time
	Title string
}

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

// CheckErrNonFatal : Logs non fatal errors
func CheckErrNonFatal(err error) {
	if err != nil {
		log.Println(err)
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

// FirstOfCurrentMonth just for VT for now
func FirstOfCurrentMonth(delimiter string) string {
	today := time.Now()
	month := strconv.Itoa(int(today.Month()))

	if len(month) == 1 {
		month = fmt.Sprintf("0%v", month)
	}

	return fmt.Sprintf("%v%v%v%v01", today.Year(), delimiter, month, delimiter)
}
