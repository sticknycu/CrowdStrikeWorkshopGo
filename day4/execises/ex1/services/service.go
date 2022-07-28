package services

import (
	"github.com/emicklei/go-restful/v3"
	"log"
	"net/http"
	"rest-http/v/gateways"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) StartWebService() {
	ws := new(restful.WebService)
	restful.Add(ws)

	api := gateways.NewAPI()
	api.RegisterRoutes(ws)

	log.Printf("Started serving on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
