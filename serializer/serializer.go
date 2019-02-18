package serializer

import (
	"math"

	"github.com/CanalTP/augeas/model"
)

type carParkResponse struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Lon          float64 `json:"lon"`
	Lat          float64 `json:"lat"`
	Total        int     `json:"total_places"`
	Available    int     `json:"available"`
	Occupied     int     `json:"occupied"`
	AvailablePRM int     `json:"available_PRM"`
	OccupiedPRM  int     `json:"occupied_PRM"`
}

type carParksResponse struct {
	CarParks []carParkResponse `json:"car_parks"`
}

type durationResponse struct {
	CarPark  carParkResponse `json:"car_park"`
	Distance uint64          `json:"distance"`
	Duration uint64          `json:"duration"`
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

func SerializeDurations(target *model.Coordinate, speed float64, maxParkingDuration uint64, parks []model.CarPark) durationsResponse {
	ret := make([]durationResponse, 0)
	for _, p := range parks {
		distance := target.Distance(&p)
		duration := uint64(distance * math.Sqrt(2) / speed)
		if duration > maxParkingDuration {
			continue
		}
		ret = append(ret, durationResponse{
			carParkResponse{
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
			uint64(distance),
			duration,
		})
	}
	return durationsResponse{
		Durations: ret,
	}
}
