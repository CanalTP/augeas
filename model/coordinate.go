package model

import (
	"math"

	"github.com/hongshibao/go-kdtree"
)

type Coordinate struct {
	Coords [2]float64 `json:"t"`
}

func (c *Coordinate) Lon() float64 {
	return c.Coords[0]
}

func (c *Coordinate) Lat() float64 {
	return c.Coords[1]
}

func (c *Coordinate) Dim() int {
	return 2
}

func (c *Coordinate) GetValue(dim int) float64 {
	return c.Coords[dim]
}

func (c *Coordinate) Distance(other kdtree.Point) float64 {
	degToRad := 0.01745329238
	earthRadius := 6372797.560856

	lonDelta := (c.GetValue(0) - other.GetValue(0)) * degToRad
	lonH := math.Pow(math.Sin(lonDelta*0.5), 2)
	latDelta := (c.GetValue(1) - other.GetValue(1)) * degToRad
	latH := math.Pow(math.Sin(latDelta*0.5), 2)
	a := math.Cos(c.GetValue(1)*degToRad) * math.Cos(other.GetValue(1)*degToRad)
	return earthRadius * 2 * math.Asin(math.Sqrt(latH+a*lonH))
}

func (c *Coordinate) PlaneDistance(val float64, dim int) float64 {
	var other [2]float64

	copy(other[0:], c.Coords[0:])
	other[dim] = val

	return c.Distance(&Coordinate{Coords: other})
}
