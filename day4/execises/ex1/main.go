package main

import (
	"rest-http/v/services"
)

func main() {
	s := services.NewService()
	s.StartWebService()
}
