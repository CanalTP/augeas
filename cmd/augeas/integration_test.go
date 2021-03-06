package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/CanalTP/augeas"
	"github.com/CanalTP/augeas/model"
	"github.com/CanalTP/augeas/poi_parser"
	"github.com/CanalTP/augeas/serializer"
)

var MIN_PARK_DURATION uint64 = 300

func performRequest(r http.Handler, method, path string) (int, *httptest.ResponseRecorder) {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w.Code, w
}

func compareCarParks(t *testing.T, code int, response *httptest.ResponseRecorder, wantedCode int, wantedCarParks []model.CarPark) bool {
	if code != wantedCode {
		t.Errorf("Response Code = %d, Wanted Code %d", code, wantedCode)
		return false
	}

	marshalled, _ := json.Marshal(serializer.SerializeCarParks(wantedCarParks))

	if response.Body.String() != string(marshalled) {
		t.Errorf("Response: %s", response.Body.String())
		t.Errorf("Wanted: %s", string(marshalled))
		return false
	}

	return true
}

func TestCarParks(t *testing.T) {

	carParks := poi_parser.ParsePoi("./fixtures/poi.txt", "amenity:parking", ";")
	dm := augeas.NewDataManager(carParks, MIN_PARK_DURATION)

	router := SetupRouter(dm)

	type args struct {
		router http.Handler
		method string
		path   string
	}
	tests := []struct {
		name           string
		args           args
		wantedCode     int
		wantedCarParks []model.CarPark
	}{
		{
			"test/v0/car_parks",
			args{router, "GET", "/v0/car_parks"},
			http.StatusOK,
			[]model.CarPark{
				{Coordinate: model.Coordinate{Coords: [2]float64{2.285156, 48.872505}}, ID: "937854398", Name: "Étoile - Foch"},
				{Coordinate: model.Coordinate{Coords: [2]float64{2.291498, 48.873689}}, ID: "937950603", Name: "Étoile - Foch"},
				{Coordinate: model.Coordinate{Coords: [2]float64{2.284068, 48.872286}}, ID: "939658365", Name: "Étoile - Foch"},
				{Coordinate: model.Coordinate{Coords: [2]float64{2.299415, 48.874107}}, ID: "838076170", Name: "Étoile Friedland"},
				{Coordinate: model.Coordinate{Coords: [2]float64{2.300731, 48.874457}}, ID: "838076561", Name: "Étoile Friedland"},
			},
		},
		{
			"test /v0/car_parks/id",
			args{router, "GET", "/v0/car_parks/838076561"},
			http.StatusOK,
			[]model.CarPark{
				{Coordinate: model.Coordinate{Coords: [2]float64{2.300731, 48.874457}}, ID: "838076561", Name: "Étoile Friedland"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if code, response := performRequest(tt.args.router, tt.args.method, tt.args.path); !compareCarParks(t, code, response, tt.wantedCode, tt.wantedCarParks) {
				t.Errorf("TestCarParks %v %v , want %v", tt.args.method, tt.args.path, tt.wantedCarParks)
			}
		})
	}
}

func compareParkDurations(t *testing.T, code int, response *httptest.ResponseRecorder,
	targetPoint model.Coordinate, speed float64, maxParkDuration uint64, wantedCode int, wantedCarParks []model.CarPark, wantedParkZones []model.ParkZone) bool {
	if code != wantedCode {
		t.Errorf("Response Code = %d, Wanted Code %d", code, wantedCode)
		return false
	}

	marshalled, _ := json.Marshal(serializer.SerializeDurations(wantedCarParks, wantedParkZones))

	if response.Body.String() != string(marshalled) {
		t.Errorf("Response: %s", response.Body.String())
		t.Errorf("Wanted: %s", string(marshalled))
		return false
	}

	return true
}

func TestParkingDuration(t *testing.T) {

	carParks := poi_parser.ParsePoi("./fixtures/poi.txt", "amenity:parking", ";")
	dm := augeas.NewDataManager(carParks, MIN_PARK_DURATION)

	router := SetupRouter(dm)

	type args struct {
		router          http.Handler
		method          string
		lon             float64
		lat             float64
		walkingSpeed    float64
		n               uint64
		maxParkDuration uint64
	}
	tests := []struct {
		name            string
		args            args
		wantedCode      int
		wantedCarParks  []model.CarPark
		wantedParkZones []model.ParkZone
	}{
		{
			"test /v0/park_duration?lon=2.300731&lat=48.874457&n=5&max_park_duration=500",
			args{router, "GET", 2.300731, 48.874457, 1.11, 5, 1200},
			http.StatusOK,
			[]model.CarPark{
				{Coordinate: model.Coordinate{Coords: [2]float64{2.300731, 48.874457}}, ID: "838076561", Name: "Étoile Friedland", DistanceToTarget: 0, ParkDuration: 300},
				{Coordinate: model.Coordinate{Coords: [2]float64{2.299415, 48.874107}}, ID: "838076170", Name: "Étoile Friedland", DistanceToTarget: 103, ParkDuration: 432},
				{Coordinate: model.Coordinate{Coords: [2]float64{2.291498, 48.873689}}, ID: "937950603", Name: "Étoile - Foch", DistanceToTarget: 680, ParkDuration: 1167},
			},
			[]model.ParkZone{},
		},
		{
			"test /v0/park_duration?lon=2.300731&lat=48.874457&n=1&max_park_duration=99999",
			args{router, "GET", 2.300731, 48.874457, 1.11, 1, 99999},
			http.StatusOK,
			[]model.CarPark{
				{Coordinate: model.Coordinate{Coords: [2]float64{2.300731, 48.874457}}, ID: "838076561", Name: "Étoile Friedland", DistanceToTarget: 0, ParkDuration: 300},
			},
			[]model.ParkZone{},
		},
		{
			"test /v0/park_duration?lon=2.300731&lat=48.874457&n=1&max_park_duration=99999",
			args{router, "GET", 2.300731, 48.874457, 1.11, 1, 250},
			http.StatusOK,
			[]model.CarPark{},
			[]model.ParkZone{
				model.NewParkZone("Area 51",
					uint64(250), []model.Coordinate{
						{Coords: [2]float64{2.304731, 48.876457}},
						{Coords: [2]float64{2.296731, 48.876457}},
						{Coords: [2]float64{2.296731, 48.872457}},
						{Coords: [2]float64{2.304731, 48.872457}},
						{Coords: [2]float64{2.304731, 48.876457}},
					}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := fmt.Sprintf("/v0/park_duration?lon=%f&lat=%f&n=%d&max_park_duration=%d&walking_speed=%f",
				tt.args.lon, tt.args.lat, tt.args.n, tt.args.maxParkDuration, tt.args.walkingSpeed)
			if code, response := performRequest(tt.args.router, tt.args.method, path); !compareParkDurations(t, code, response,
				model.Coordinate{Coords: [2]float64{tt.args.lon, tt.args.lat}},
				tt.args.walkingSpeed, tt.args.maxParkDuration, http.StatusOK, tt.wantedCarParks, tt.wantedParkZones) {
				t.Errorf("TestParkingDuration %v, want %v %v", tt.args.method, tt.wantedCarParks, tt.wantedParkZones)
			}
		})
	}
}
