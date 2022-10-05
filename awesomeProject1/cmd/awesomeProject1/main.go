package main

import (
	"awesomeProject1/internal/api"
	"log"
)

func main() {
	log.Println("Applicaton start!")
	api.StartServer()
	log.Println("Application Terminated!")
}
