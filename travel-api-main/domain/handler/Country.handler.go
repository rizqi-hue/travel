package handler

import (
	"context"

	"github.com/travel-api-main/pkg/pb"
)

func (s Server) GetsCountry(ctx context.Context, req *pb.GetsCountryRequest) (*pb.GetsCountryResponse, error) {

	result, err := s.CountryService.GetsCountry(ctx, req)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s Server) GetCountry(ctx context.Context, req *pb.GetCountryRequest) (*pb.GetCountryResponse, error) {
	result, err := s.CountryService.GetCountry(ctx, req)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s Server) CreateCountry(ctx context.Context, req *pb.CreateCountryRequest) (*pb.CreateCountryResponse, error) {
	result, err := s.CountryService.CreateCountry(ctx, req)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s Server) UpdateCountry(ctx context.Context, req *pb.UpdateCountryRequest) (*pb.UpdateCountryResponse, error) {
	result, err := s.CountryService.UpdateCountry(ctx, req)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s Server) DeleteCountry(ctx context.Context, req *pb.DeleteCountryRequest) (*pb.DeleteCountryResponse, error) {
	result, err := s.CountryService.DeleteCountry(ctx, req)
	if err != nil {
		return nil, err
	}

	return result, nil
}
