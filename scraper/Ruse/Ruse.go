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
)

// URL for specific town
const (
	RuseURL = "http://free-spirit-city.eu/aktualno-za-sedmicata"
	POIName = "Ruse"
)

// Ruse : Scraper specific
func Ruse(db *sql.DB) {
	res, err := http.Get(RuseURL)
	defer res.Body.Close()

	helpers.CheckErr(err)
	helpers.StatusCodeChecker(res.StatusCode, RuseURL)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	helpers.CheckErr(err)

	var event helpers.GenericScraperEvent

	doc.Find(".photos_common .news_box").Each(func(i int, s *goquery.Selection) {
		split := strings.Split(strings.TrimSpace(s.Find(".news_date_list").Text()), ".")
		day := split[0]
		month := split[1]
		year := split[2]
		dateFixed := fmt.Sprintf("%v/%v/%v", year, month, day)
		event.Date, err = dateparse.ParseAny(dateFixed)
		if err == nil {
			event.Title = strings.TrimSpace(s.Find(".news_title").Text())

			if !dbwrapper.EventExists(dateFixed, event.Title, POIName, db) {
				dbwrapper.DBInsert(dateFixed, event.Title, POIName, db)
			}
		} else {
			helpers.CheckErrNonFatal(err)
		}
	})

}
