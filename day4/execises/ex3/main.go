package main

import "http-rest/v/services"

func main() {
	s := services.NewService()
	s.StartWebService()
}
