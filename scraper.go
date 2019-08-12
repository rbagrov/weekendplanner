package main

import (
	sql "database/sql"
	dbwrapper "github.com/weekendplanner/dbwrapper"
	BalchikScraper "github.com/weekendplanner/scraper/Balchik"
	BeloslavScraper "github.com/weekendplanner/scraper/Beloslav"
	RuseScraper "github.com/weekendplanner/scraper/Ruse"
)

var db *sql.DB

// ScrapeStarter starts scraping
func main() {
	db := dbwrapper.DBOpen()
	defer dbwrapper.DBClose(db)
	BeloslavScraper.Beloslav(db)
	RuseScraper.Ruse(db)
	BalchikScraper.Balchik(db)
}
