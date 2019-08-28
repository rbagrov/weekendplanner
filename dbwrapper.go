package dbwrapper

import (
	sql "database/sql"
	fmt "fmt"
	// Required by the driver
	_ "github.com/lib/pq"
	helpers "github.com/weekendplanner/helpers"
	settings "github.com/weekendplanner/settings"
	time "time"
)

// DBOpen : returns db connection
func DBOpen() *sql.DB {
	settings := settings.Settings()
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		settings.DbUser, settings.DbPassword, settings.DbName, settings.DbHost, settings.DbPort)
	db, err := sql.Open("postgres", dbinfo)
	helpers.CheckErr(err)
	return db
}

// DBClose : closes db connection
func DBClose(db *sql.DB) {
	db.Close()
}

// DBEventExists : check if event exists
func DBEventExists(date string, title string, poiName string, db *sql.DB) bool {
	var exists bool
	var query string
	query = fmt.Sprintf("SELECT exists(SELECT id FROM events JOIN poi ON(events.poi_id = poi.id) WHERE events.date = '%s' and events.event_title = '%s' and poi.name = '%s')", date, title, poiName)
	err := db.QueryRow(query).Scan(&exists)
	helpers.SQLCheckErr(err)
	return exists
}

// DBPOIExists : check if poi exists
func DBPOIExists(poiName string, db *sql.DB) bool {
	var exists bool
	var query string
	query = fmt.Sprintf("SELECT exists(SELECT id FROM poi WHERE name = '%s')", poiName)
	err := db.QueryRow(query).Scan(&exists)
	helpers.SQLCheckErr(err)
	return exists
}

// DBGetPOIId : returns POI id
func DBGetPOIId(poiName string, db *sql.DB) helpers.POI {
	statement := "SELECT id FROM poi WHERE name = $1"
	rows, err := db.Query(statement, poiName)
	helpers.SQLCheckErr(err)
	defer rows.Close()

	var poi helpers.POI

	for rows.Next() {
		err := rows.Scan(&poi.ID)
		helpers.CheckErr(err)
	}
	return poi
}

// DBPOIAdd : adds POI
func DBPOIAdd(poiName string, db *sql.DB) {
	stmt, err := db.Prepare("INSERT INTO poi(created_on, updated_on, name) VALUES($1, $2, $3)")
	helpers.SQLCheckErr(err)

	_, err1 := stmt.Exec(time.Now(), time.Now(), poiName)
	helpers.SQLCheckErr(err1)
}

// DBAddEvent : increments persistence
func DBAddEvent(date string, title string, poiName string, db *sql.DB) {
	if !DBPOIExists(poiName, db) {
		DBPOIAdd(poiName, db)
	}

	poi := DBGetPOIId(poiName, db)
	stmt, err := db.Prepare("INSERT INTO events(created_on, updated_on, date, event_title, poi_id) VALUES($1,$2,$3,$4,$5)")
	helpers.SQLCheckErr(err)

	_, err1 := stmt.Exec(time.Now(), time.Now(), date, title, poi.ID)
	helpers.SQLCheckErr(err1)
}

// DBGetPOI : Gets all the data from db for the given POI
func DBGetPOI(poi string, db *sql.DB) []helpers.GenericPOIEvent {
	var eventsList []helpers.GenericPOIEvent
	var event helpers.GenericPOIEvent
	var ScanDate time.Time
	sqlStatement := "SELECT events.date, events.event_title, poi.name FROM events JOIN poi ON (events.poi_id = poi.id) WHERE events.date = $1 AND poi.name LIKE '%' || $2 || '%' ORDER BY events.date ASC;"
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
	sqlStatement := "SELECT events.date, events.event_title, poi.name FROM events JOIN poi ON (events.poi_id = poi.id) ORDER BY events.created_on DESC LIMIT 5;"
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
