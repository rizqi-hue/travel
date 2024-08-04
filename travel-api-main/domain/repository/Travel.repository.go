package repository

import (
	"github.com/travel-api-main/domain/models"
	"github.com/travel-api-main/pkg/database"
	"github.com/travel-api-main/pkg/pb"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type TravelRepository interface {
	Gets(req *pb.GetsTravelRequest) (travels []models.Travel, totalRow int64, totalPage int64, canDoNextPage bool, canDoPreviousPage bool, err error)
	Insert(travel models.Travel) (*models.Travel, error)
	Update(travel models.Travel, id int64) (*models.Travel, error)
	Delete(id int64) (bool, error)
	GetById(req *pb.GetTravelRequest) (*models.Travel, error)
	GetByCountry(country string) (*models.Travel, error)
}

type travelRepository struct {
	H *gorm.DB
}

func NewTravelRepository() *travelRepository {
	return &travelRepository{
		H: database.DB,
	}
}

func filterTravel(travel *models.Travel) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		DB := db

		if travel.TravelledCountriesCode != "" {
			DB.Where("travelled_countries_code = ?", travel.TravelledCountriesCode)
		}

		return DB
	}
}

func sortTravel(req *pb.GetsTravelRequest) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		DB := db

		if req.GETS_TRAVEL.Sort.Field != "" && req.GETS_TRAVEL.Sort.Order != "" {
			DB.Order(req.GETS_TRAVEL.Sort.Field + " " + req.GETS_TRAVEL.Sort.Order)
		}

		return DB
	}
}

func (r *travelRepository) Gets(req *pb.GetsTravelRequest) (travels []models.Travel, totalRow int64, totalPage int64, canDoNextPage bool, canDoPreviousPage bool, err error) {

	page := req.GETS_TRAVEL.Page
	size := req.GETS_TRAVEL.Size
	query := req.GETS_TRAVEL.Query

	if size == 0 {
		size = 5
	}

	db := r.H.Scopes(
		Paginate(page, size),
		filterTravel(&models.Travel{
			TravelledCountriesCode: query.TravelledCountriesCode,
		}),
		sortTravel(req),
	).
		Preload("Country").
		Joins("JOIN country ON travel.travelled_countries_code = country.alpha_3")

	count := r.H.Scopes(
		filterTravel(&models.Travel{
			TravelledCountriesCode: query.TravelledCountriesCode,
		}),
	).
		Preload("Country").
		Joins("JOIN country ON travel.travelled_countries_code = country.alpha_3")

	if query.Search != "" {
		searchPattern := "%" + query.Search + "%"

		db = db.Where("country.name ILIKE ?", searchPattern)
		count = count.Where("country.name ILIKE ?", searchPattern)
	}

	if result := db.Find(&travels); result.Error != nil {
		return nil, 0, 0, false, false, status.Error(400, "error gets travel")
	}

	if result := count.Model(&models.Travel{}).Count(&totalRow); result.Error != nil {
		return nil, 0, 0, false, false, status.Error(400, "error count travel")
	}

	totalPage = (totalRow / size) + 1

	canDoNextPage = page < totalPage
	canDoPreviousPage = page > 1

	return travels, totalRow, totalPage, canDoNextPage, canDoPreviousPage, nil
}

// func (r *travelRepository) Gets(req *pb.GetsTravelRequest) (travels []models.Travel, totalRow int64, totalPage int64, canDoNextPage bool, canDoPreviousPage bool, err error) {

// 	page := req.GETS_TRAVEL.Page
// 	size := req.GETS_TRAVEL.Size
// 	query := req.GETS_TRAVEL.Query

// 	if size == 0 {
// 		size = 5
// 	}

// 	mapQuery := make(map[string]interface{}, 0)

// 	if query.Search != "" {
// 		mapQuery["country.name"] = query.Search
// 	}

// 	if result := r.H.Scopes(
// 		Paginate(page, size),
// 		filterTravel(&models.Travel{
// 			TravelledCountriesCode: req.GETS_TRAVEL.Query.TravelledCountriesCode,
// 		}),
// 		sortTravel(req),
// 	).
// 		Preload("Country").
// 		Joins("JOIN country ON travel.travelled_countries_code = country.alpha_3").
// 		Where(mapQuery).
// 		Find(&travels); result.Error != nil {
// 		return nil, 0, 0, false, false, status.Error(400, "error gets travel")
// 	}

// 	if result := r.H.Scopes(filterTravel(&models.Travel{})).Model(models.Travel{}).Count(&totalRow); result.Error != nil {
// 		return nil, 0, 0, false, false, status.Error(400, "error count travel")
// 	}

// 	totalPage = (totalRow / size) + 1

// 	// Determine if next and previous pages are available
// 	canDoNextPage = page < totalPage
// 	canDoPreviousPage = page > 1

// 	return travels, totalRow, totalPage, canDoNextPage, canDoPreviousPage, nil
// }

func (r *travelRepository) Insert(travel models.Travel) (*models.Travel, error) {

	err := r.H.Save(&travel).Error
	if err != nil {
		return nil, err
	}

	return &travel, nil
}

func (r *travelRepository) Update(travel models.Travel, id int64) (*models.Travel, error) {

	if result := r.H.Where("id = ?", id).Updates(&travel); result.Error != nil {
		return nil, status.Error(400, "Record not found")
	}

	return &travel, nil
}

func (r *travelRepository) Delete(id int64) (bool, error) {

	err := r.H.Where("id = ?", id).Delete(&models.Travel{}).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *travelRepository) GetById(req *pb.GetTravelRequest) (travel *models.Travel, err error) {

	if result := r.H.
		Scopes(
			filterTravel(&models.Travel{
				Id: req.GET_TRAVEL.Id,
			})).First(&travel); result.Error != nil {
		return nil, status.Error(400, "Record not found")
	}

	return travel, nil
}

func (r *travelRepository) GetByCountry(country string) (travel *models.Travel, err error) {

	if result := r.H.
		Scopes(
			filterTravel(&models.Travel{
				TravelledCountriesCode: country,
			})).First(&travel); result.Error != nil {
		return nil, status.Error(400, "Record not found")
	}

	return travel, nil
}
