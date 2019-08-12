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
	time "time"
)

// URL for specific town
const (
	BalchikURL = "http://www.balchik.bg/bg/news"
	POIName    = "Balchik"
)

// Event struct
type Event struct {
	date  time.Time
	title string
}

// Balchik : Scraper specific
func Balchik(db *sql.DB) {
	res, err := http.Get(BalchikURL)
	defer res.Body.Close()

	helpers.CheckErr(err)
	helpers.StatusCodeChecker(res.StatusCode, BalchikURL)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	helpers.CheckErr(err)

	var event Event

	doc.Find(".box_news_left .news_info").Each(func(i int, s *goquery.Selection) {
		split := strings.Split(strings.TrimSpace(s.Find(".news_title").Text()), ".")
		day := split[0]
		month := split[1]
		year := split[2]
		dateFixed := fmt.Sprintf("%v/%v/%v", year, month, day)
		event.date, err = dateparse.ParseAny(dateFixed)
		helpers.CheckErr(err)

		event.title = strings.TrimSpace(s.Find(".news_title_more").Text())

		if !dbwrapper.EventExists(dateFixed, event.title, POIName, db) {
			dbwrapper.DBInsert(dateFixed, event.title, POIName, db)
		}
	})

}
