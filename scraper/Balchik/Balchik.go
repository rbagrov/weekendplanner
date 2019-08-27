package balchik

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
	BalchikURL = "http://www.balchik.bg/bg/news"
	POIName    = "Balchik"
)

// Balchik : Scraper specific
func Balchik(db *sql.DB) {
	res, err := http.Get(BalchikURL)
	defer res.Body.Close()

	helpers.CheckErr(err)
	helpers.StatusCodeChecker(res.StatusCode, BalchikURL)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	helpers.CheckErr(err)

	var event helpers.GenericScraperEvent

	doc.Find(".box_news_left .news_info").Each(func(i int, s *goquery.Selection) {
		split := strings.Split(strings.TrimSpace(s.Find(".news_title").Text()), ".")
		day := split[0]
		month := split[1]
		year := split[2]
		dateFixed := fmt.Sprintf("%v/%v/%v", year, month, day)
		event.Date, err = dateparse.ParseAny(dateFixed)
		if err == nil {
			event.Title = strings.TrimSpace(s.Find(".news_title_more").Text())

			if !dbwrapper.DBEventExists(dateFixed, event.Title, POIName, db) {
				dbwrapper.DBAddEvent(dateFixed, event.Title, POIName, db)
			}
		} else {
			helpers.CheckErrNonFatal(err)
		}
	})

}
