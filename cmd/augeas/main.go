package main

import (
	"flag"

	"github.com/CanalTP/augeas"
	"github.com/CanalTP/augeas/poi_parser"
	"github.com/gin-gonic/gin"
)

func main() {

	poiFile := flag.String("poi", "", "poi.txt file's path")
	carParkType := flag.String("car_park_type", "amenity:parking", "car park's poi type")
	csvComma := flag.String("csv_comma", ";", "csv delimiter")

	flag.Parse()

	carParks := poi_parser.ParsePoi(poiFile, carParkType, csvComma)

	dm := augeas.NewDataManager(carParks)

	router := gin.Default()

	router.GET("/v0/car_parks", augeas.GetCarParksHanlder(dm))
	router.GET("/v0/car_parks/:car_park_id", augeas.GetCarParkByIDHanlder(dm))
	router.GET("/v0/park_duration", augeas.GetParkDurationHandler(dm))

	router.Run(":1337")
}
