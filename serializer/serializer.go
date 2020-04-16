package serializer

import (
	"github.com/CanalTP/augeas/model"
)

type carParkResponse struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Lon          float64 `json:"lon"`
	Lat          float64 `json:"lat"`
	Total        uint64  `json:"total_places,default=0"`
	Available    uint64  `json:"available,default=0"`
	Occupied     uint64  `json:"occupied,default=0"`
	AvailablePRM uint64  `json:"available_PRM,default=0"`
	OccupiedPRM  uint64  `json:"occupied_PRM,default=0"`
}

type parkZoneResponse struct {
	Name    string        `json:"name"`
	GeoJSON model.GeoJSON `json:"geojson"`
}

type carParksResponse struct {
	CarParks []carParkResponse `json:"car_parks"`
}

type durationResponse struct {
	CarPark  *carParkResponse  `json:"car_park,omitempty"`
	ParkZone *parkZoneResponse `json:"park_zone,omitempty"`
	Distance uint64            `json:"distance,default=0"`
	Duration uint64            `json:"duration,default=0"`
}

type durationsResponse struct {
	Durations []durationResponse `json:"durations"`
}

func SerializeCarParks(parks []model.CarPark) carParksResponse {
	ret := make([]carParkResponse, len(parks))
	for i, p := range parks {
		ret[i] = carParkResponse{
			ID:           p.ID,
			Name:         p.Name,
			Lon:          p.Lon(),
			Lat:          p.Lat(),
			Total:        p.Total,
			Available:    p.Available,
			Occupied:     p.Occupied,
			AvailablePRM: p.AvailablePRM,
			OccupiedPRM:  p.OccupiedPRM,
		}
	}
	return carParksResponse{
		CarParks: ret,
	}
}

func SerializeDurations(parks []model.CarPark, zones []model.ParkZone) durationsResponse {

	ret := []durationResponse{}
	for _, p := range parks {
		ret = append(ret, durationResponse{
			CarPark: &carParkResponse{
				ID:           p.ID,
				Name:         p.Name,
				Lon:          p.Lon(),
				Lat:          p.Lat(),
				Total:        p.Total,
				Available:    p.Available,
				Occupied:     p.Occupied,
				AvailablePRM: p.AvailablePRM,
				OccupiedPRM:  p.OccupiedPRM,
			},
			Distance: p.DistanceToTarget,
			Duration: p.ParkDuration,
		})
	}
	for _, z := range zones {
		ret = append(ret, durationResponse{
			ParkZone: &parkZoneResponse{
				Name:    z.Name,
				GeoJSON: z.GeoJSON,
			},
			Distance: z.DistanceToTarget,
			Duration: z.ParkDuration,
		})
	}
	return durationsResponse{
		Durations: ret,
	}
}
