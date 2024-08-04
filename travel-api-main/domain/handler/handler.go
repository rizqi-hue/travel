package handler

import (
	"github.com/travel-api-main/domain/services"
	"google.golang.org/grpc"
)

type Server struct {
	grpc           *grpc.Server
	CountryService services.CountryService
	TravelService  services.TravelService
}

func InitServer(grpc *grpc.Server,
	countryService services.CountryService,
	travelService services.TravelService,

) *Server {
	return &Server{
		grpc:           grpc,
		CountryService: countryService,
		TravelService:  travelService,
	}
}
