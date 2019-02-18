package augeas

import (
	"log"

	"github.com/CanalTP/augeas/model"
	"github.com/hongshibao/go-kdtree"
)

type DataManager struct {
	carParks []model.CarPark
	kdTree   *kdtree.KDTree
}

func NewDataManager(carParks []model.CarPark) *DataManager {
	dm := DataManager{}

	dm.carParks = carParks

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

func (dm *DataManager) GetNearestCarPark(targetPoint *model.Coordinate, n uint64) []model.CarPark {
	neighbours := dm.kdTree.KNN(targetPoint, int(n))
	ret := make([]model.CarPark, len(neighbours))
	log.Printf("%d car parks have been found", len(neighbours))
	for idx, n := range neighbours {
		p := n.(*model.CarPark)
		ret[idx] = *p
	}
	return ret
}
