package dbwrapper

import (
	sql "database/sql"
	fmt "fmt"
	// Required by the driver
	_ "github.com/lib/pq"
	helpers "github.com/weekendplanner/helpers"
	time "time"
)

// DSN constructor data
const (
	DbUser     = "postgres"
	DbPassword = "postgres"
	DbName     = "weekendplanner"
	DbHost     = "localhost"
	DbPort     = "5432"
)

// DBOpen : returns db connection
func DBOpen() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		DbUser, DbPassword, DbName, DbHost, DbPort)
	db, err := sql.Open("postgres", dbinfo)
	helpers.CheckErr(err)
	return db
}

// DBClose : closes db connection
func DBClose(db *sql.DB) {
	db.Close()
}

// EventExists : check if event exists
func EventExists(date string, title string, poiName string, db *sql.DB) bool {
	var exists bool
	var query string
	query = fmt.Sprintf("SELECT exists(SELECT id FROM events WHERE date = '%s' and event = '%s' and poi_name = '%s')", date, title, poiName)
	err := db.QueryRow(query).Scan(&exists)
	helpers.SQLCheckErr(err)
	return exists
}

// DBInsert : increments persistence
func DBInsert(date string, title string, poiName string, db *sql.DB) {
	stmt, err := db.Prepare("INSERT INTO events(created_on, updated_on, date, event, poi_name) VALUES($1,$2,$3,$4, $5)")
	helpers.SQLCheckErr(err)

	_, err1 := stmt.Exec(time.Now(), time.Now(), date, title, poiName)
	helpers.SQLCheckErr(err1)
}

// DBGetPOI : Gets all the data from db for the given POI
func DBGetPOI(poi string, db *sql.DB) []helpers.GenericPOIEvent {
	var eventsList []helpers.GenericPOIEvent
	var event helpers.GenericPOIEvent
	var ScanDate time.Time
	sqlStatement := "SELECT date, event, poi_name FROM events WHERE date = $1 AND poi_name = $2 ORDER BY date ASC;"
	rows, err := db.Query(sqlStatement, time.Now(), poi)
	helpers.SQLCheckErr(err)
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&ScanDate, &event.Event, &event.PoiName)
		helpers.CheckErr(err)
		year, month, day := ScanDate.Date()
		monthStr := fmt.Sprintf("%v", month)
		eventsList = append(eventsList, helpers.GenericPOIEvent{Day: day, Month: monthStr, Year: year, Event: event.Event, PoiName: event.PoiName})
	}
	return eventsList
}

// DBGetLastEvents : Gets latest 4 events
func DBGetLastEvents(db *sql.DB) []helpers.GenericEvent {
	var eventsList []helpers.GenericEvent
	var event helpers.GenericEvent
	var ScanDate time.Time
	sqlStatement := "SELECT date, event, poi_name FROM events ORDER BY created_on DESC LIMIT 5;"
	rows, err := db.Query(sqlStatement)
	helpers.SQLCheckErr(err)
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&ScanDate, &event.Event, &event.PoiName)
		helpers.CheckErr(err)
		eventsList = append(eventsList, helpers.GenericEvent{Date: ScanDate.Format("2006-01-02"), Event: event.Event, PoiName: event.PoiName})
	}
	return eventsList
}
