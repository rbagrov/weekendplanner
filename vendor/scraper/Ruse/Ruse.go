package ruse

import (
	sql "database/sql"
	fmt "fmt"
	goquery "github.com/PuerkitoBio/goquery"
	dateparse "github.com/araddon/dateparse"
	dbwrapper "github.com/weekendplanner/dbwrapper"
	helpers "github.com/weekendplanner/helpers"
	http "net/http"
	strings "strings"
	time "time"
)

// URL for specific town
const (
	RuseURL = "http://free-spirit-city.eu/aktualno-za-sedmicata"
	POIName = "Ruse"
)

// Event struct
type Event struct {
	date  time.Time
	title string
}

// Ruse : Scraper specific
func Ruse(db *sql.DB) {
	res, err := http.Get(RuseURL)
	defer res.Body.Close()

	helpers.CheckErr(err)
	helpers.StatusCodeChecker(res.StatusCode, RuseURL)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	helpers.CheckErr(err)

	var event Event

	doc.Find(".photos_common .news_box").Each(func(i int, s *goquery.Selection) {
		split := strings.Split(strings.TrimSpace(s.Find(".news_date_list").Text()), ".")
		day := split[0]
		month := split[1]
		year := split[2]
		dateFixed := fmt.Sprintf("%v/%v/%v", year, month, day)
		event.date, err = dateparse.ParseAny(dateFixed)
		helpers.CheckErr(err)

		event.title = strings.TrimSpace(s.Find(".news_title").Text())

		if !dbwrapper.EventExists(dateFixed, event.title, POIName, db) {
			dbwrapper.DBInsert(dateFixed, event.title, POIName, db)
		}
	})

}
