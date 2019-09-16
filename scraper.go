package main

import (
	sql "database/sql"
	dbwrapper "github.com/weekendplanner/dbwrapper"
	BalchikScraper "github.com/weekendplanner/scraper/Balchik"
	BeloslavScraper "github.com/weekendplanner/scraper/Beloslav"
	RuseScraper "github.com/weekendplanner/scraper/Ruse"
	VelikoTurnovoScraper "github.com/weekendplanner/scraper/VelikoTurnovo"
)

var db *sql.DB

// ScrapeStarter starts scraping
func main() {
	db := dbwrapper.DBOpen()
	defer dbwrapper.DBClose(db)
	go BeloslavScraper.Beloslav(db)
	go RuseScraper.Ruse(db)
	go BalchikScraper.Balchik(db)
	go VelikoTurnovoScraper.VelikoTurnovo(db)
}
