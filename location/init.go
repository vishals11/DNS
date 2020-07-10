package location

import (
	"log"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

// InitializeLocationAPI - read sectorID from ENV variable and call the API
func InitializeLocationAPI(r *mux.Router)  {
	env := os.Getenv("SECTOR_ID")
	sectorID, err := strconv.ParseFloat(env, 64)
	if err != nil{
		log.Fatal("Environment variable missing....")
		return
	}
	service := NewService(sectorID)
	handler := NewLocateDataBankHandler(service)
	r.HandleFunc("/v1/dns", handler.GetDataBankLocation).Methods("POST")
}
