package poi_parser

import (
	"reflect"
	"testing"

	"github.com/CanalTP/augeas/model"
)

func TestParsePoi(t *testing.T) {
	type args struct {
		poiFile     string
		carParkType string
		csvComma    string
	}
	tests := []struct {
		name string
		args args
		want []model.CarPark
	}{
		{
			"simple parsing",
			args{"./fixtures/poi.txt", "amenity:parking", ";"},
			[]model.CarPark{
				{Coordinate: model.Coordinate{Coords: [2]float64{2.285156, 48.872505}}, ID: "937854398", Name: "Étoile - Foch"},
				{Coordinate: model.Coordinate{Coords: [2]float64{2.291498, 48.873689}}, ID: "937950603", Name: "Étoile - Foch"},
				{Coordinate: model.Coordinate{Coords: [2]float64{2.284068, 48.872286}}, ID: "939658365", Name: "Étoile - Foch"},
				{Coordinate: model.Coordinate{Coords: [2]float64{2.299415, 48.874107}}, ID: "838076170", Name: "Étoile Friedland"},
				{Coordinate: model.Coordinate{Coords: [2]float64{2.300731, 48.874457}}, ID: "838076561", Name: "Étoile Friedland"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParsePoi(tt.args.poiFile, tt.args.carParkType, tt.args.csvComma); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParsePoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
