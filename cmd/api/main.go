package main

import (
	"log"

	"github.com/iamRosty/api-skeleton/internal/app/api"
)

var ()

func init() {

}

func main() {
	log.Printf("Ok")
	// server instance init
	server := api.New()
	// api server start
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
