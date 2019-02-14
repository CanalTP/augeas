package augeas

import (
	"net/http"
	"strconv"

	"github.com/CanalTP/augeas/model"
	"github.com/CanalTP/augeas/serializer"
	"github.com/gin-gonic/gin"
)

func GetCarParksHanlder(dm *DataManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, serializer.SerializeCarParks(dm.GetAllCarParks()))
	}
}

func GetCarParkByIDHanlder(dm *DataManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("car_park_id")
		c.JSON(http.StatusOK, serializer.SerializeCarParks(dm.GetCarParkByID(&id)))
	}
}

type durationParams struct {
	Lon                float64
	Lat                float64
	N                  uint64
	WalkingSpeed       float64
	MaxParkingDuration uint64
}

func getParams(c *gin.Context) (*durationParams, error) {
	lon, err := strconv.ParseFloat(c.Query("lon"), 64)
	if err != nil {
		return nil, err
	}
	lat, err := strconv.ParseFloat(c.Query("lat"), 64)
	if err != nil {
		return nil, err
	}
	n, err := strconv.ParseUint(c.DefaultQuery("n", "5"), 10, 32)
	if err != nil {
		return nil, err
	}
	walkingSpeed, err := strconv.ParseFloat(c.DefaultQuery("walking_speed", "1.11"), 64)
	if err != nil {
		return nil, err
	}
	maxParkingDuration, err := strconv.ParseUint(c.DefaultQuery("max_parking_duration", "1200"), 10, 64)
	if err != nil {
		return nil, err
	}

	return &durationParams{lon, lat, n, walkingSpeed, maxParkingDuration}, nil
}

func GetParkDurationHandler(dm *DataManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		params, err := getParams(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		targetPoint := model.Coordinate{[2]float64{params.Lon, params.Lat}}
		neighbours := dm.GetNearestCarPark(&targetPoint, params.N)

		c.JSON(http.StatusOK, serializer.SerializeDurations(&targetPoint, params.WalkingSpeed, params.MaxParkingDuration, neighbours))
	}
}
