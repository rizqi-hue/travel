package services

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/travel-api-main/domain/models"
	"github.com/travel-api-main/domain/repository"
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
	TravelRepository    repository.TravelRepository
	TravelCSVRepository repository.TravelCSVRepository
}

func NewTravelService() *travelService {
	return &travelService{}
}

func (s *travelService) SetTravelRepository(repo repository.TravelRepository) *travelService {
	s.TravelRepository = repo
	return s
}

func (s *travelService) SetTravelCSVRepository(repo repository.TravelCSVRepository) *travelService {
	s.TravelCSVRepository = repo
	return s
}

func (s *travelService) Validate() *travelService {
	if s.TravelRepository == nil {
		panic("handler need travel repository")
	}

	return s
}

func (s *travelService) GetsTravel(ctx context.Context, req *pb.GetsTravelRequest) (*pb.GetsTravelResponse, error) {
	datas, totalRow, totalPage,  canDoNextPage, canDoPreviousPage, err := s.TravelRepository.Gets(req)

	if err != nil {
		return nil, err
	}

	result := make([]*pb.GetsTravelResponseData_DataTravel, 0)
	for _, data := range datas {
		result = append(result, &pb.GetsTravelResponseData_DataTravel{
			Field: data.TravelledCountriesCode,
			Label: data.Country.Name,
			Total: data.Total,
		})
	}

	pagination := repository.Pagination(req.GETS_TRAVEL.Page, req.GETS_TRAVEL.Size)

	return &pb.GetsTravelResponse{
		GETS_TRAVEL: &pb.GetsTravelResponseData{
			Search:       req.GETS_TRAVEL.Query.Search,
			Next:         canDoNextPage,
			Back:         canDoPreviousPage,
			Limit:        pagination.Limit,
			Offset:       pagination.Offset,
			TotalPage:    totalPage,
			CurrentPage:  req.GETS_TRAVEL.Page,
			Sort:         req.GETS_TRAVEL.Sort.Field,
			Order:        req.GETS_TRAVEL.Sort.Order,
			RecordsTotal: totalRow,
			Data:         result,
		},
	}, nil
}

func (s *travelService) GetTravel(ctx context.Context, req *pb.GetTravelRequest) (*pb.GetTravelResponse, error) {
	data, err := s.TravelRepository.GetById(req)

	if err != nil {
		return nil, err
	}

	return &pb.GetTravelResponse{
		GET_TRAVEL: &pb.GetTravelResponseData{
			Status: http.StatusOK,
			Data: &pb.Travel{
				Id:                     data.Id,
				TravelledCountriesCode: data.TravelledCountriesCode,
				Total:                  data.Total,
				CreatedAt:              data.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt:              data.UpdatedAt.Format("2006-01-02 15:04:05"),
			},
		},
	}, nil
}

func (s *travelService) CreateTravel(ctx context.Context, req *pb.CreateTravelRequest) (*pb.CreateTravelResponse, error) {
	for {
		csv, totalRow, _, err := s.TravelCSVRepository.Gets()
		if err != nil {
			return nil, err
		}

		if csv == nil {
			return &pb.CreateTravelResponse{
				CREATE_TRAVEL: &pb.CreateTravelResponseData{
					Status: http.StatusOK,
				},
			}, nil
		}

		if totalRow == 0 {
			return &pb.CreateTravelResponse{
				CREATE_TRAVEL: &pb.CreateTravelResponseData{
					Status: http.StatusOK,
				},
			}, nil
		}

		for _, travelCSV := range csv {
			arr := strings.Split(travelCSV.TravelledCountriesCode, ",")
			// Print the result
			for _, v := range arr {
				getTravel, _ := s.TravelRepository.GetByCountry(v)
				if getTravel != nil {
					// update
					comp := models.Travel{
						Id:                     getTravel.Id,
						TravelledCountriesCode: v,
						Total:                  travelCSV.Total + getTravel.Total,
						UpdatedAt:              time.Now(),
					}

					_, err := s.TravelRepository.Update(comp, getTravel.Id)

					if err != nil {
						return nil, err
					}

					// update csv
					compCSV := models.TravelCSV{
						Id:     travelCSV.Id,
						Status: true,
					}

					_, errUpdate := s.TravelCSVRepository.Update(compCSV, travelCSV.Id)
					if errUpdate != nil {
						return nil, errUpdate
					}

				} else {
					// insert
					comp := models.Travel{
						TravelledCountriesCode: v,
						Total:                  travelCSV.Total,
						CreatedAt:              time.Now(),
					}

					_, err := s.TravelRepository.Insert(comp)

					if err != nil {
						return nil, err
					}

					// update csv
					compCSV := models.TravelCSV{
						Id:     travelCSV.Id,
						Status: true,
					}

					_, errUpdate := s.TravelCSVRepository.Update(compCSV, travelCSV.Id)
					if errUpdate != nil {
						return nil, errUpdate
					}
				}
			}
		}

		time.Sleep(3 * time.Second) // Wait for 3 seconds before checking again
	}
}

func (s *travelService) UpdateTravel(ctx context.Context, req *pb.UpdateTravelRequest) (*pb.UpdateTravelResponse, error) {
	UPDATE_TRAVEL := req.UPDATE_TRAVEL

	comp := models.Travel{
		Id:                     UPDATE_TRAVEL.Id,
		TravelledCountriesCode: UPDATE_TRAVEL.TravelledCountriesCode,
		Total:                  UPDATE_TRAVEL.Total,
		UpdatedAt:              time.Now(),
	}

	data, err := s.TravelRepository.Update(comp, UPDATE_TRAVEL.Id)

	if err != nil {
		return nil, err
	}

	return &pb.UpdateTravelResponse{
		UPDATE_TRAVEL: &pb.UpdateTravelResponseData{
			Status: http.StatusOK,
			Data: &pb.Travel{
				Id:                     data.Id,
				TravelledCountriesCode: data.TravelledCountriesCode,
				Total:                  data.Total,
				CreatedAt:              data.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt:              data.UpdatedAt.Format("2006-01-02 15:04:05"),
			},
		},
	}, nil
}

func (s *travelService) DeleteTravel(ctx context.Context, req *pb.DeleteTravelRequest) (*pb.DeleteTravelResponse, error) {
	DELETE_TRAVEL := req.DELETE_TRAVEL

	_, err := s.TravelRepository.Delete(DELETE_TRAVEL.Id)

	if err != nil {
		return nil, err
	}

	return &pb.DeleteTravelResponse{
		DELETE_TRAVEL: &pb.DeleteTravelResponseData{
			Status: http.StatusOK,
		},
	}, nil
}
