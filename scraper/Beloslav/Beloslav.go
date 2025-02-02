package beloslav

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
	BeloslavURL = "http://www.beloslav.org/news-1-1.html"
	POIName     = "Beloslav"
	Latitude    = "43.187310"
	Longitude   = "27.704150"
)

// Beloslav : Scraper specific
func Beloslav(db *sql.DB) {
	res, err := http.Get(BeloslavURL)
	defer res.Body.Close()

	helpers.CheckErr(err)
	helpers.StatusCodeChecker(res.StatusCode, BeloslavURL)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	helpers.CheckErr(err)

	var event helpers.GenericScraperEvent
	var poi helpers.POIInit

	selection1 := doc.Find(".section-news-block .new-items")
	selection1.Find(".new-item").Each(func(i int, s *goquery.Selection) {
		dateSplit := strings.Split(strings.TrimSpace(s.Find(".new-date-text").Text()), ".")
		day := dateSplit[0]
		month := dateSplit[1]
		year := dateSplit[2]
		dateFixed := fmt.Sprintf("%v/%v/%v", year, month, day)
		event.Date, err = dateparse.ParseAny(dateFixed)

		poi.Name = POIName
		poi.Latitude = Latitude
		poi.Longitude = Longitude

		if err == nil {
			event.Title = strings.TrimSpace(s.Find(".new-item-caption").Text())
			if !dbwrapper.DBEventExists(dateFixed, event.Title, poi.Name, db) {
				dbwrapper.DBAddEvent(dateFixed, event.Title, poi, db)
			}
		} else {
			helpers.CheckErrNonFatal(err)
		}
	})

}
