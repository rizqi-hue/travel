package serve

import (
	"context"
	"log"
	"net"

	def "github.com/travel-api-main/cmd/config"
	"github.com/travel-api-main/domain/handler"
	"github.com/travel-api-main/domain/repository"
	"github.com/travel-api-main/domain/services"
	"github.com/travel-api-main/pkg/config"
	"github.com/travel-api-main/pkg/pb"
	"google.golang.org/grpc"
)

func NewGrpcServer(ctx context.Context) error {
	if err := config.Load(def.DefaultConfig, Config); err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", config.GetString("rpc_travel"))

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	countryRepo := repository.NewCountryRepository()
	countryService := services.NewCountryService().
		SetCountryRepository(countryRepo).
		Validate()

	travelRepo := repository.NewTravelRepository()
	travelCSVRepo := repository.NewTravelCSVRepository()
	travelService := services.NewTravelService().
		SetTravelRepository(travelRepo).
		SetTravelCSVRepository(travelCSVRepo).
		Validate()

	s := grpc.NewServer()
	handler := handler.InitServer(s, countryService, travelService)
	pb.RegisterCountryServiceServer(s, handler)
	pb.RegisterTravelServiceServer(s, handler)

	log.Println("Serving gRPC on 0.0.0.0:" + config.GetString("rpc_travel"))
	s.Serve(lis)

	return nil
}
