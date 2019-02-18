package poi_parser

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/CanalTP/augeas/model"
)

func newCarPark(record []string, carParkType string) (*model.CarPark, error) {
	poiType := record[6]
	if poiType != carParkType {
		return nil, nil
	}
	id := record[0]
	name := record[1]
	poiLat, err := strconv.ParseFloat(record[4], 64)
	if err != nil {
		return nil, err
	}
	poiLon, err := strconv.ParseFloat(record[5], 64)
	if err != nil {
		return nil, err
	}
	return &model.CarPark{
		Coordinate:   model.Coordinate{[2]float64{poiLon, poiLat}},
		ID:           id,
		Name:         name,
		Total:        0,
		Available:    0,
		Occupied:     0,
		AvailablePRM: 0,
		OccupiedPRM:  0,
	}, nil
}

func ParsePoi(poiFile string, carParkType string, csvComma string) []model.CarPark {
	f, err := os.Open(poiFile)
	defer f.Close()

	if err != nil {
		log.Panicln(err)
		panic(err)
	}
	csvr := csv.NewReader(f)
	csvr.Comma = ([]rune(csvComma))[0]

	if err != nil {
		log.Panicln(err)
		panic(err)
	}

	carParks := make([]model.CarPark, 0)

	// Skip the first line (Header)
	csvr.Read()

	log.Printf("Parsing poi file: %s with carParcType: %s", poiFile, carParkType)
	for {
		row, err := csvr.Read()

		if err != nil {
			if err != io.EOF {
				log.Panicln(err)
				continue
			} else {
				break
			}
		}
		cp, err := newCarPark(row, carParkType)
		if err != nil {
			log.Panicln(err)
			continue
		}
		if cp == nil {
			continue
		}
		carParks = append(carParks, *cp)
	}
	log.Printf("Finishing reading. %d car parks have been found", len(carParks))

	return carParks
}
