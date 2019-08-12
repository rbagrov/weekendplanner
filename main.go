package main

import (
	gin "github.com/gin-gonic/gin"
	dbwrapper "github.com/weekendplanner/dbwrapper"
	http "net/http"
)

func main() {
	db := dbwrapper.DBOpen()
	defer dbwrapper.DBClose(db)

	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	router.Static("css", "templates/css")
	router.Static("images", "templates/images")
	router.Static("scripts", "templates/scripts")

	router.GET("/", func(c *gin.Context) {
		events := dbwrapper.DBGetLastEvents(db)

		c.HTML(
			http.StatusOK,
			"index.html",
			gin.H{
				"events": events,
			},
		)
	})

	router.POST("/", func(c *gin.Context) {
		poi := c.PostForm("poi")
		poiEvents := dbwrapper.DBGetPOI(poi, db)
		events := dbwrapper.DBGetLastEvents(db)

		c.HTML(
			http.StatusOK,
			"index.html",
			gin.H{
				"events":    events,
				"poiEvents": poiEvents,
			},
		)

	})

	router.Run(":8090")
}
