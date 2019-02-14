package serializer

import (
	"math"

	"github.com/CanalTP/augeas/model"
)

type carParkResposne struct {
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
	CarParks []*carParkResposne `json:"car_parks"`
}

type durationResponse struct {
	CarPark  carParkResposne `json:"car_park"`
	Distance int             `json:"distance"`
	Duration int             `json:"duration"`
}

type durationsResponse struct {
	Durations []*durationResponse `json:"durations"`
}

func SerializeCarParks(parks []*model.CarPark) carParksResponse {
	ret := make([]*carParkResposne, len(parks))
	for i, p := range parks {
		ret[i] = &carParkResposne{
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

func SerializeDurations(target *model.Coordinate, speed float64, maxParkingDuration uint64, parks []*model.CarPark) durationsResponse {
	ret := make([]*durationResponse, len(parks))
	for i, p := range parks {
		d := target.Distance(p)
		ret[i] = &durationResponse{
			carParkResposne{
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
			int(d),
			int(math.Min(d*1.414/speed, float64(maxParkingDuration))),
		}
	}
	return durationsResponse{
		Durations: ret,
	}
}
