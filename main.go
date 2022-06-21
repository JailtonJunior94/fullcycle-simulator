package main

import (
	"log"

	"github.com/jailtonjunior94/fullcycle-simulator/application/route"
)

func main() {
	route := route.Route{
		ID:       "1",
		ClientID: "1",
	}

	if err := route.LoadPositions(); err != nil {
		log.Println(err)
	}

	json, err := route.ExportJsonPositions()
	if err != nil {
		log.Println(err)
	}

	log.Println(json[0])
}
