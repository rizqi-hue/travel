package services

import (
	"context"
	"net/http"
	"time"

	"github.com/travel-api-main/domain/models"
	"github.com/travel-api-main/domain/repository"
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
	CountryRepository repository.CountryRepository
}

func NewCountryService() *countryService {
	return &countryService{}
}

func (s *countryService) SetCountryRepository(countryRepo repository.CountryRepository) *countryService {
	s.CountryRepository = countryRepo
	return s
}

func (s *countryService) Validate() *countryService {
	if s.CountryRepository == nil {
		panic("handler need country repository")
	}

	return s
}

func (s *countryService) GetsCountry(ctx context.Context, req *pb.GetsCountryRequest) (*pb.GetsCountryResponse, error) {
	datas, totalRow, totalPage, err := s.CountryRepository.Gets(req)

	if err != nil {
		return nil, err
	}

	result := make([]*pb.Country, 0)
	for _, data := range datas {
		result = append(result, &pb.Country{
			Id:                     data.Id,
			Name:                   data.Name,
			Alpha2:                 data.Alpha2,
			Alpha3:                 data.Alpha3,
			CountryCode:            data.CountryCode,
			Iso31662:               data.Iso31662,
			Region:                 data.Region,
			SubRegion:              data.SubRegion,
			IntermediateRegion:     data.IntermediateRegion,
			RegionCode:             data.Region,
			SubRegionCode:          data.SubRegionCode,
			IntermediateRegionCode: data.IntermediateRegionCode,
			CreatedAt:              data.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:              data.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return &pb.GetsCountryResponse{
		GETS_COUNTRY: &pb.GetsCountryResponseData{
			Status:      http.StatusOK,
			TotalData:   totalRow,
			TotalPage:   totalPage,
			CurrentPage: req.GETS_COUNTRY.Page,
			Data:        result,
		},
	}, nil
}

func (s *countryService) GetCountry(ctx context.Context, req *pb.GetCountryRequest) (*pb.GetCountryResponse, error) {
	data, err := s.CountryRepository.GetById(req)

	if err != nil {
		return nil, err
	}

	return &pb.GetCountryResponse{
		GET_COUNTRY: &pb.GetCountryResponseData{
			Status: http.StatusOK,
			Data: &pb.Country{
				Id:                     data.Id,
				Name:                   data.Name,
				Alpha2:                 data.Alpha2,
				Alpha3:                 data.Alpha3,
				CountryCode:            data.CountryCode,
				Iso31662:               data.Iso31662,
				Region:                 data.Region,
				SubRegion:              data.SubRegion,
				IntermediateRegion:     data.IntermediateRegion,
				RegionCode:             data.Region,
				SubRegionCode:          data.SubRegionCode,
				IntermediateRegionCode: data.IntermediateRegionCode,
				CreatedAt:              data.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt:              data.UpdatedAt.Format("2006-01-02 15:04:05"),
			},
		},
	}, nil
}

func (s *countryService) CreateCountry(ctx context.Context, req *pb.CreateCountryRequest) (*pb.CreateCountryResponse, error) {
	CREATE_COUNTRY := req.CREATE_COUNTRY

	comp := models.Country{
		Name:                   CREATE_COUNTRY.Name,
		Alpha2:                 CREATE_COUNTRY.Alpha2,
		Alpha3:                 CREATE_COUNTRY.Alpha3,
		CountryCode:            CREATE_COUNTRY.CountryCode,
		Iso31662:               CREATE_COUNTRY.Iso31662,
		Region:                 CREATE_COUNTRY.Region,
		SubRegion:              CREATE_COUNTRY.SubRegion,
		IntermediateRegion:     CREATE_COUNTRY.IntermediateRegion,
		RegionCode:             CREATE_COUNTRY.Region,
		SubRegionCode:          CREATE_COUNTRY.SubRegionCode,
		IntermediateRegionCode: CREATE_COUNTRY.IntermediateRegionCode,
		CreatedAt:              time.Now(),
	}

	data, err := s.CountryRepository.Insert(comp)

	if err != nil {
		return nil, err
	}

	return &pb.CreateCountryResponse{
		CREATE_COUNTRY: &pb.CreateCountryResponseData{
			Status: http.StatusOK,
			Data: &pb.Country{
				Id:                     data.Id,
				Name:                   data.Name,
				Alpha2:                 data.Alpha2,
				Alpha3:                 data.Alpha3,
				CountryCode:            data.CountryCode,
				Iso31662:               data.Iso31662,
				Region:                 data.Region,
				SubRegion:              data.SubRegion,
				IntermediateRegion:     data.IntermediateRegion,
				RegionCode:             data.Region,
				SubRegionCode:          data.SubRegionCode,
				IntermediateRegionCode: data.IntermediateRegionCode,
				CreatedAt:              data.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt:              data.UpdatedAt.Format("2006-01-02 15:04:05"),
			},
		},
	}, nil
}

func (s *countryService) UpdateCountry(ctx context.Context, req *pb.UpdateCountryRequest) (*pb.UpdateCountryResponse, error) {
	UPDATE_COUNTRY := req.UPDATE_COUNTRY

	comp := models.Country{
		Id:                     UPDATE_COUNTRY.Id,
		Name:                   UPDATE_COUNTRY.Name,
		Alpha2:                 UPDATE_COUNTRY.Alpha2,
		Alpha3:                 UPDATE_COUNTRY.Alpha3,
		CountryCode:            UPDATE_COUNTRY.CountryCode,
		Iso31662:               UPDATE_COUNTRY.Iso31662,
		Region:                 UPDATE_COUNTRY.Region,
		SubRegion:              UPDATE_COUNTRY.SubRegion,
		IntermediateRegion:     UPDATE_COUNTRY.IntermediateRegion,
		RegionCode:             UPDATE_COUNTRY.Region,
		SubRegionCode:          UPDATE_COUNTRY.SubRegionCode,
		IntermediateRegionCode: UPDATE_COUNTRY.IntermediateRegionCode,
		CreatedAt:              time.Now(),
	}

	data, err := s.CountryRepository.Update(comp, UPDATE_COUNTRY.Id)

	if err != nil {
		return nil, err
	}

	return &pb.UpdateCountryResponse{
		UPDATE_COUNTRY: &pb.UpdateCountryResponseData{
			Status: http.StatusOK,
			Data: &pb.Country{
				Id:                     data.Id,
				Name:                   data.Name,
				Alpha2:                 data.Alpha2,
				Alpha3:                 data.Alpha3,
				CountryCode:            data.CountryCode,
				Iso31662:               data.Iso31662,
				Region:                 data.Region,
				SubRegion:              data.SubRegion,
				IntermediateRegion:     data.IntermediateRegion,
				RegionCode:             data.Region,
				SubRegionCode:          data.SubRegionCode,
				IntermediateRegionCode: data.IntermediateRegionCode,
				CreatedAt:              data.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt:              data.UpdatedAt.Format("2006-01-02 15:04:05"),
			},
		},
	}, nil
}

func (s *countryService) DeleteCountry(ctx context.Context, req *pb.DeleteCountryRequest) (*pb.DeleteCountryResponse, error) {
	DELETE_COUNTRY := req.DELETE_COUNTRY

	_, err := s.CountryRepository.Delete(DELETE_COUNTRY.Id)

	if err != nil {
		return nil, err
	}

	return &pb.DeleteCountryResponse{
		DELETE_COUNTRY: &pb.DeleteCountryResponseData{
			Status: http.StatusOK,
		},
	}, nil
}
