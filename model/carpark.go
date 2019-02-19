package model

type CarPark struct {
	Coordinate
	ID               string
	Name             string
	Total            uint64
	Available        uint64
	Occupied         uint64
	AvailablePRM     uint64
	OccupiedPRM      uint64
	DistanceToTarget uint64
	ParkDuration     uint64
}
