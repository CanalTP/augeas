package model

type Properties struct {
	Whatever string `json:"whatever"`
}

type Geometry struct {
	Type        string         `json:"type"`
	Coordinates [][][2]float64 `json:"coordinates"`
}

type Feature struct {
	Type       string     `json:"type"`
	Properties Properties `json:"properties"`
	Geometry   Geometry   `json:"geometry"`
}

type GeoJSON struct {
	Type     string    `json:"type"`
	Features []Feature `json:"features"`
}
