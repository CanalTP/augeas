package model

import (
	"math"
	"testing"

	"github.com/hongshibao/go-kdtree"
)

// https://www.geodatasource.com/distance-calculator
func TestCoordinate_Distance(t *testing.T) {
	type fields struct {
		Coords [2]float64
	}
	type args struct {
		other kdtree.Point
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		{
			"simple geographical distance",
			fields{[2]float64{2.333333, 48.866666}},
			args{&Coordinate{[2]float64{2.3610305, 48.87077002}}},
			2077,
		},
		{
			"same point",
			fields{[2]float64{2.333333, 48.866666}},
			args{&Coordinate{[2]float64{2.333333, 48.866666}}},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Coordinate{
				Coords: tt.fields.Coords,
			}
			if got := c.Distance(tt.args.other); math.Round(got) != tt.want {
				t.Errorf("Coordinate.Distance() = %v, want %v", math.Round(got), tt.want)
			}
		})
	}
}

func TestCoordinate_PlaneDistance(t *testing.T) {
	type fields struct {
		Coords [2]float64
	}
	type args struct {
		val float64
		dim int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		{
			"plane projection on Dim 0",
			fields{[2]float64{2.333333, 48.866666}},
			args{2.3, 0},
			2439,
		},
		{
			"plane projection on Dim 1",
			fields{[2]float64{2.333333, 48.866666}},
			args{48.8, 1},
			7415,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Coordinate{
				Coords: tt.fields.Coords,
			}
			if got := c.PlaneDistance(tt.args.val, tt.args.dim); math.Round(got) != tt.want {
				t.Errorf("Coordinate.PlaneDistance() = %v, want %v", math.Round(got), tt.want)
			}
		})
	}
}
