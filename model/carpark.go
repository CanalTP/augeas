package model

type CarPark struct {
	Coordinate
	ID           string
	Name         string
	Total        uint64
	Available    uint64
	Occupied     uint64
	AvailablePRM uint64
	OccupiedPRM  uint64
	// Internal attribute, used to simplify serializer interfaces
	DistanceToTarget uint64
	// Same as above
	ParkDuration uint64
}
