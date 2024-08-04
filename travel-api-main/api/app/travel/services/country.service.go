package services

import (
	"context"
	"fmt"

	"github.com/travel-api-main/api/app/travel/client"
	"github.com/travel-api-main/pkg/pb"
)

type CountryService interface {
	GetsCountry(ctx context.Context, req *pb.GetsCountryRequest) (*pb.GetsCountryResponse, error)
	GetCountry(ctx context.Context, req *pb.GetCountryRequest) (*pb.GetCountryResponse, error)
	CreateCountry(ctx context.Context, req *pb.CreateCountryRequest) (*pb.CreateCountryResponse, error)
	UpdateCountry(ctx context.Context, req *pb.UpdateCountryRequest) (*pb.UpdateCountryResponse, error)
	DeleteCountry(ctx context.Context, req *pb.DeleteCountryRequest) (*pb.DeleteCountryResponse, error)
}

type countryService struct {
	TravelClient client.TravelClient
}

func NewCountryService(client client.TravelClient) CountryService {
	return &countryService{TravelClient: client}
}

func (s *countryService) GetsCountry(ctx context.Context, req *pb.GetsCountryRequest) (*pb.GetsCountryResponse, error) {

	out, err := s.TravelClient.CountryService.GetsCountry(ctx, req)

	if err != nil {
		fmt.Printf("err gets country %s", err.Error())
		return out, err
	}

	return out, nil
}

func (s *countryService) GetCountry(ctx context.Context, req *pb.GetCountryRequest) (*pb.GetCountryResponse, error) {

	out, err := s.TravelClient.CountryService.GetCountry(ctx, req)

	if err != nil {
		fmt.Printf("err get country %s", err.Error())
		return out, err
	}

	return out, nil
}

func (s *countryService) CreateCountry(ctx context.Context, req *pb.CreateCountryRequest) (*pb.CreateCountryResponse, error) {

	out, err := s.TravelClient.CountryService.CreateCountry(ctx, req)

	if err != nil {
		fmt.Printf("err create country %s", err.Error())
		return out, err
	}

	return out, nil
}

func (s *countryService) UpdateCountry(ctx context.Context, req *pb.UpdateCountryRequest) (*pb.UpdateCountryResponse, error) {

	out, err := s.TravelClient.CountryService.UpdateCountry(ctx, req)

	if err != nil {
		fmt.Printf("err update country %s", err.Error())
		return out, err
	}

	return out, nil
}

func (s *countryService) DeleteCountry(ctx context.Context, req *pb.DeleteCountryRequest) (*pb.DeleteCountryResponse, error) {
	out, err := s.TravelClient.CountryService.DeleteCountry(ctx, req)

	if err != nil {
		fmt.Printf("err delete %s", err.Error())
		return out, err
	}

	return out, nil
}
