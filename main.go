package main

import (
	"fmt"
	"log"
	"net/http"

	"DroneNavigationSystem/router"
)

func main() {
	r := router.InitializeRouter()
	fmt.Println("Server listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
