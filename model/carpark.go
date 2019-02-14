package model

type CarPark struct {
	Coordinate
	ID           string
	Name         string
	Total        int
	Available    int
	Occupied     int
	AvailablePRM int
	OccupiedPRM  int
}
