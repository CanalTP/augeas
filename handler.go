package augeas

import (
	"fmt"
	"net/http"

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
		c.JSON(http.StatusOK, serializer.SerializeCarParks(dm.GetCarParkByID(id)))
	}
}

type durationParams struct {
	Lon             float64 `form:"lon"`
	Lat             float64 `form:"lat"`
	N               uint64  `form:"n,default=5"`
	WalkingSpeed    float64 `form:"walking_speed,default=1.11"`
	MaxParkDuration uint64  `form:"max_park_duration,default=1200"`
}

func getParams(c *gin.Context) (*durationParams, error) {
	var params durationParams
	err := c.Bind(&params)
	if err != nil {
		return nil, err
	}
	fmt.Println(params)
	return &params, nil
}

func GetParkDurationHandler(dm *DataManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		params, err := getParams(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		targetPoint := model.Coordinate{Coords: [2]float64{params.Lon, params.Lat}}
		nearestCarParks := dm.GetNearestCarPark(&targetPoint, params.N, params.WalkingSpeed, params.MaxParkDuration)
		// We cannot find car parks that meet all criterias, so we just tell the user
		// to park in a park zone.
		// Since there is no park zones defined in the data so far, we just creat one around the target point
		parkZones := make([]model.ParkZone, 0)
		if len(nearestCarParks) == 0 {
			parkZones = append(parkZones, model.NewParkZone("Area 51",
				params.MaxParkDuration, []model.Coordinate{
					{Coords: [2]float64{targetPoint.Lon() + 0.004, targetPoint.Lat() + 0.002}},
					{Coords: [2]float64{targetPoint.Lon() - 0.004, targetPoint.Lat() + 0.002}},
					{Coords: [2]float64{targetPoint.Lon() - 0.004, targetPoint.Lat() - 0.002}},
					{Coords: [2]float64{targetPoint.Lon() + 0.004, targetPoint.Lat() - 0.002}},
					{Coords: [2]float64{targetPoint.Lon() + 0.004, targetPoint.Lat() + 0.002}},
				}))
		}
		c.JSON(http.StatusOK, serializer.SerializeDurations(nearestCarParks, parkZones))
	}
}
