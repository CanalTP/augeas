package augeas

import (
	"log"
	"math"

	"github.com/CanalTP/augeas/model"
	"github.com/hongshibao/go-kdtree"
)

type DataManager struct {
	carParks        []model.CarPark
	minParkDuration uint64
	kdTree          *kdtree.KDTree
}

func NewDataManager(carParks []model.CarPark, minParkDuration uint64) *DataManager {
	dm := DataManager{}

	dm.carParks = carParks
	dm.minParkDuration = minParkDuration
	// Build the KdTree
	points := make([]kdtree.Point, 0)

	for idx := range carParks {
		points = append(points, &carParks[idx])
	}
	dm.kdTree = kdtree.NewKDTree(points)

	log.Print("Finish building data")

	return &dm
}

func (dm *DataManager) GetAllCarParks() []model.CarPark {
	return dm.carParks
}

func (dm *DataManager) GetCarParkByID(id string) []model.CarPark {
	ret := make([]model.CarPark, 0)
	for _, v := range dm.carParks {
		if v.ID == id {
			ret = append(ret, v)
		}
	}
	return ret
}

func (dm *DataManager) GetNearestCarPark(targetPoint *model.Coordinate, n uint64, walkingSpeed float64, maxParkingDuration uint64) []model.CarPark {
	neighbours := dm.kdTree.KNN(targetPoint, int(n))
	ret := []model.CarPark{}
	log.Printf("%d car parks have been found by KNN", len(neighbours))
	for _, n := range neighbours {
		distance := targetPoint.Distance(n)
		// Duration = min_park_duration + walking_duration
		duration := uint64(distance*math.Sqrt(2)/walkingSpeed) + dm.minParkDuration
		if duration > maxParkingDuration {
			continue
		}

		// Downcasting
		p := n.(*model.CarPark)

		// Copy
		newPark := *p
		newPark.DistanceToTarget = uint64(distance)
		newPark.ParkDuration = duration

		ret = append(ret, newPark)
	}
	log.Printf("%d car parks have been found", len(ret))
	return ret
}
