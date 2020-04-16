package model

type ParkZone struct {
	Name         string
	ParkDuration uint64
	GeoJSON      GeoJSON
	// Internal attribute, used to simplify serializer interfaces
	DistanceToTarget uint64
}

func NewParkZone(name string, duration uint64, coordinates []Coordinate) ParkZone {
	geoJSONCoords := make([][2]float64, len(coordinates))

	for i, c := range coordinates {
		geoJSONCoords[i] = c.Coords
	}
	return ParkZone{
		Name:         name,
		ParkDuration: duration,
		GeoJSON: GeoJSON{
			Type: "FeatureCollection",
			Features: []Feature{
				Feature{
					Type: "Feature",
					Properties: Properties{
						Whatever: "Whatever, Who cares",
					},

					Geometry: Geometry{
						Type:        "Polygon",
						Coordinates: [][][2]float64{geoJSONCoords},
					},
				},
			},
		},
		DistanceToTarget: 0,
	}
}
