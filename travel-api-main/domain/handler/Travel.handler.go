package handler

import (
	"context"

	"github.com/travel-api-main/pkg/pb"
)

func (s Server) GetsTravel(ctx context.Context, req *pb.GetsTravelRequest) (*pb.GetsTravelResponse, error) {

	result, err := s.TravelService.GetsTravel(ctx, req)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s Server) GetTravel(ctx context.Context, req *pb.GetTravelRequest) (*pb.GetTravelResponse, error) {
	result, err := s.TravelService.GetTravel(ctx, req)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s Server) CreateTravel(ctx context.Context, req *pb.CreateTravelRequest) (*pb.CreateTravelResponse, error) {
	result, err := s.TravelService.CreateTravel(ctx, req)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s Server) UpdateTravel(ctx context.Context, req *pb.UpdateTravelRequest) (*pb.UpdateTravelResponse, error) {
	result, err := s.TravelService.UpdateTravel(ctx, req)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s Server) DeleteTravel(ctx context.Context, req *pb.DeleteTravelRequest) (*pb.DeleteTravelResponse, error) {
	result, err := s.TravelService.DeleteTravel(ctx, req)
	if err != nil {
		return nil, err
	}

	return result, nil
}
