package velikoturnovo

import (
	sql "database/sql"
	fmt "fmt"
	goquery "github.com/PuerkitoBio/goquery"
	dateparse "github.com/araddon/dateparse"
	dbwrapper "github.com/weekendplanner/dbwrapper"
	helpers "github.com/weekendplanner/helpers"
	log "log"
	http "net/http"
	strings "strings"
	//time "time"
)

// URL for specific town
const (
	POIName = "VelikoTurnovo"
)

// VelikoTurnovo : Scraper specific
func VelikoTurnovo(db *sql.DB) {
	VelikoTurnovoURL := fmt.Sprintf("https://www.veliko-tarnovo.bg/bg/events/?month=%v", helpers.FirstOfCurrentMonth("-"))
	res, err := http.Get(VelikoTurnovoURL)
	defer res.Body.Close()

	helpers.CheckErr(err)
	helpers.StatusCodeChecker(res.StatusCode, VelikoTurnovoURL)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	helpers.CheckErr(err)

	var event helpers.GenericScraperEvent

	doc.Find(".news-list .events").Each(func(i int, s *goquery.Selection) {
		bothDates := strings.TrimSpace(s.Find(".events-info-date").Text())
		StartSplit := strings.Split(strings.Split(bothDates[0:16], " ")[0], ".")
		EndSplit := strings.Split(strings.Split(bothDates[16:32], " ")[0], ".")

		start, errStart := dateparse.ParseAny(fmt.Sprintf("%v/%v/%v", StartSplit[2], StartSplit[1], StartSplit[0]))

		end, errEnd := dateparse.ParseAny(fmt.Sprintf("%v/%v/%v", EndSplit[2], EndSplit[1], EndSplit[0]))

		if errStart == nil && errEnd == nil {

			event.Title = strings.TrimSpace(s.Find("h5").Text())

			for i := 0; i <= int(end.Sub(start).Hours()/24); i++ {
				date := start.AddDate(0, 0, i)

				dateString := fmt.Sprintf("%v/%v/%v", date.Year(), int(date.Month()), date.Day())

				if !dbwrapper.EventExists(dateString, event.Title, POIName, db) {
					dbwrapper.DBInsert(dateString, event.Title, POIName, db)
				}

			}

		} else {
			log.Println("Error in date parsing for VelikoTurnovo")
		}
	})
}
