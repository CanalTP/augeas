package augeas

import (
	"log"

	"github.com/CanalTP/augeas/model"
	"github.com/hongshibao/go-kdtree"
)

type DataManager struct {
	carParksMap map[string]*model.CarPark
	carParksVec []*model.CarPark
	kdTree      *kdtree.KDTree
}

func NewDataManager(carParks []*model.CarPark) *DataManager {
	dm := DataManager{}

	dm.carParksVec = carParks

	dm.carParksMap = map[string]*model.CarPark{}
	// Build the search map
	for _, v := range carParks {
		dm.carParksMap[v.ID] = v
	}

	// Build the KdTree
	points := make([]kdtree.Point, 0)

	for _, v := range carParks {
		points = append(points, v)
	}
	dm.kdTree = kdtree.NewKDTree(points)

	log.Print("Finish building data")

	return &dm
}

func (dm *DataManager) GetAllCarParks() []*model.CarPark {
	return dm.carParksVec
}

func (dm *DataManager) GetCarParkByID(id *string) []*model.CarPark {
	park, ok := dm.carParksMap[*id]
	ret := make([]*model.CarPark, 0)
	if ok {
		ret = append(ret, park)
	}
	return ret
}

func (dm *DataManager) GetNearestCarPark(targetPoint *model.Coordinate, n uint64) []*model.CarPark {
	neighbours := dm.kdTree.KNN(targetPoint, int(n))
	ret := make([]*model.CarPark, len(neighbours))
	log.Printf("%d car parks have been found", len(neighbours))
	for idx, n := range neighbours {
		p := n.(*model.CarPark)
		ret[idx] = p
	}
	return ret
}
