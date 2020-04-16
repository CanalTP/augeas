package main

import (
	"flag"

	"github.com/CanalTP/augeas"
	"github.com/CanalTP/augeas/poi_parser"
	"github.com/gin-gonic/gin"
)

func SetupRouter(dm *augeas.DataManager) *gin.Engine {
	router := gin.Default()

	router.GET("/v0/car_parks", augeas.GETCarParksHanlder(dm))
	router.GET("/v0/car_parks/:car_park_id", augeas.GETCarParkByIDHanlder(dm))
	router.GET("/v0/park_duration", augeas.GETParkDurationHandler(dm))
	router.POST("/v0/park_duration", augeas.POSTParkDurationHandler(dm))

	return router
}

func main() {

	poiFile := flag.String("poi", "", "poi.txt file's path")
	minParkDuration := flag.Int("min_park_duration", 300, "minimun park duration")
	carParkType := flag.String("car_park_type", "amenity:parking", "car park's poi type")
	csvComma := flag.String("csv_comma", ";", "csv delimiter")

	flag.Parse()

	carParks := poi_parser.ParsePoi(*poiFile, *carParkType, *csvComma)

	dm := augeas.NewDataManager(carParks, uint64(*minParkDuration))

	router := SetupRouter(dm)

	if err := router.Run(":1337"); err != nil {
		panic(err)
	}
}
