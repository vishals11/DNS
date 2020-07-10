package location

import (
	"fmt"
	"strconv"

	"DroneNavigationSystem/location/model"
)

type Service interface {
	GetLocation(request model.LocateDataBankRequest) (float64, error)
}

type service struct {
	SectorID float64
}

// NewService - Return calling service
func NewService(id float64) *service{
	return &service{SectorID:id}
}

// GetLocation - Returns exact location of a data-bank to upload the data
func (s *service) GetLocation(request model.LocateDataBankRequest) (float64, error) {
	x, err := strconv.ParseFloat(request.X, 64)
	if err != nil {
		return 0, fmt.Errorf("Error while parsing")
	}
	y, err := strconv.ParseFloat(request.Y, 64)
	if err != nil {
		return 0, fmt.Errorf("Error while parsing")
	}
	z, err := strconv.ParseFloat(request.Z, 64)
	if err != nil {
		return 0, fmt.Errorf("Error while parsing")
	}
	vel, err := strconv.ParseFloat(request.Vel, 64)
	if err != nil {
		return 0, fmt.Errorf("Error while parsing")
	}

	return x * s.SectorID + y * s.SectorID + z * s.SectorID + vel, nil
}