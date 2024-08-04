package client

import (
	"github.com/travel-api-main/pkg/config"
	"github.com/travel-api-main/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type TravelClient struct {
	CountryService pb.CountryServiceClient
	TravelService pb.TravelServiceClient
}

func InitAuthClient() TravelClient {
	// ---- start dialing auth grpc ----
	conn, err := grpc.Dial(config.GetString("rpc_travel"), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic(err)
	}

	countryGrpc := pb.NewCountryServiceClient(conn)
	travelGrpc := pb.NewTravelServiceClient(conn)

	return TravelClient{
		CountryService: countryGrpc,
		TravelService: travelGrpc,
	}
}
