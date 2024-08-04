package services

import (
	"context"
	"fmt"

	"github.com/travel-api-main/api/app/travel/client"
	"github.com/travel-api-main/pkg/pb"
)

type TravelService interface {
	GetsTravel(ctx context.Context, req *pb.GetsTravelRequest) (*pb.GetsTravelResponse, error)
	GetTravel(ctx context.Context, req *pb.GetTravelRequest) (*pb.GetTravelResponse, error)
	CreateTravel(ctx context.Context, req *pb.CreateTravelRequest) (*pb.CreateTravelResponse, error)
	UpdateTravel(ctx context.Context, req *pb.UpdateTravelRequest) (*pb.UpdateTravelResponse, error)
	DeleteTravel(ctx context.Context, req *pb.DeleteTravelRequest) (*pb.DeleteTravelResponse, error)
}

type travelService struct {
	TravelClient client.TravelClient
}

func NewTravelService(client client.TravelClient) TravelService {
	return &travelService{TravelClient: client}
}

func (s *travelService) GetsTravel(ctx context.Context, req *pb.GetsTravelRequest) (*pb.GetsTravelResponse, error) {

	out, err := s.TravelClient.TravelService.GetsTravel(ctx, req)

	if err != nil {
		fmt.Printf("err gets travel %s", err.Error())
		return out, err
	}

	return out, nil
}

func (s *travelService) GetTravel(ctx context.Context, req *pb.GetTravelRequest) (*pb.GetTravelResponse, error) {

	out, err := s.TravelClient.TravelService.GetTravel(ctx, req)

	if err != nil {
		fmt.Printf("err get travel %s", err.Error())
		return out, err
	}

	return out, nil
}

func (s *travelService) CreateTravel(ctx context.Context, req *pb.CreateTravelRequest) (*pb.CreateTravelResponse, error) {

	out, err := s.TravelClient.TravelService.CreateTravel(ctx, req)

	if err != nil {
		fmt.Printf("err create travel %s", err.Error())
		return out, err
	}

	return out, nil
}

func (s *travelService) UpdateTravel(ctx context.Context, req *pb.UpdateTravelRequest) (*pb.UpdateTravelResponse, error) {

	out, err := s.TravelClient.TravelService.UpdateTravel(ctx, req)

	if err != nil {
		fmt.Printf("err update travel %s", err.Error())
		return out, err
	}

	return out, nil
}

func (s *travelService) DeleteTravel(ctx context.Context, req *pb.DeleteTravelRequest) (*pb.DeleteTravelResponse, error) {
	out, err := s.TravelClient.TravelService.DeleteTravel(ctx, req)

	if err != nil {
		fmt.Printf("err delete %s", err.Error())
		return out, err
	}

	return out, nil
}
