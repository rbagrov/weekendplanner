package main

import (
	sql "database/sql"
	fmt "fmt"
	dbwrapper "github.com/weekendplanner/dbwrapper"
	BalchikScraper "github.com/weekendplanner/scraper/Balchik"
	RuseScraper "github.com/weekendplanner/scraper/Ruse"
)

var db *sql.DB

// ScrapeStarter starts scraping
func main() {
	db := dbwrapper.DBOpen()
	defer dbwrapper.DBClose(db)
	BalchikScraper.Balchik(db)
	RuseScraper.Ruse(db)
	fmt.Println(dbwrapper.DBGetPOI("Balchik", db))
}
