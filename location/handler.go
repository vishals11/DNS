package location

import (
	"encoding/json"
	"net/http"

	"DroneNavigationSystem/location/model"
)

type LocateDataBank interface {
	GetDataBankLocation(w http.ResponseWriter, r *http.Request)
}

type locateDataBank struct {
	s Service
}

func NewLocateDataBankHandler(s Service) *locateDataBank{
	return &locateDataBank{s:s}
}

// GetDataBankLocation - Request and response handler
func (l *locateDataBank) GetDataBankLocation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var request model.LocateDataBankRequest
	var err error
	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var response model.LocateDataBankResponse
	response.Loc, err = l.s.GetLocation(request)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
