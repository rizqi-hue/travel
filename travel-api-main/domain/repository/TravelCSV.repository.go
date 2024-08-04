package repository

import (
	"github.com/travel-api-main/domain/models"
	"github.com/travel-api-main/pkg/database"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type TravelCSVRepository interface {
	Gets() (travelCSVs []models.TravelCSV, totalRow int64, totalPage int64, err error)
	Insert(travelCSV models.TravelCSV) (*models.TravelCSV, error)
	Update(travelCSV models.TravelCSV, id int64) (*models.TravelCSV, error)
	Delete(id int64) (bool, error)
}

type travelCSVRepository struct {
	H *gorm.DB
}

func NewTravelCSVRepository() *travelCSVRepository {
	return &travelCSVRepository{
		H: database.DB,
	}
}

func FilterTravelCSV(travelCSV *models.TravelCSV) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		DB := db

		DB.Where("status = ?", false)

		return DB
	}
}

func (r *travelCSVRepository) Gets() (travelCSVs []models.TravelCSV, totalRow int64, totalPage int64, err error) {

	var page int64 = 1
	var size int64 = 100

	if result := r.H.Scopes(
		Paginate(page, size),
		FilterTravelCSV(&models.TravelCSV{})).
		Find(&travelCSVs); result.Error != nil {
		return nil, 0, 0, status.Error(400, "error gets travelCSV")
	}

	if result := r.H.Scopes(FilterTravelCSV(&models.TravelCSV{})).Model(models.TravelCSV{}).Count(&totalRow); result.Error != nil {
		return nil, 0, 0, status.Error(400, "error count travelCSV")
	}

	if size == 0 {
		size = 1
	}

	totalPage = (totalRow / size) + 1

	return travelCSVs, totalRow, totalPage, nil
}

func (r *travelCSVRepository) Insert(travelCSV models.TravelCSV) (*models.TravelCSV, error) {

	err := r.H.Save(&travelCSV).Error
	if err != nil {
		return nil, err
	}

	return &travelCSV, nil
}

func (r *travelCSVRepository) Update(travelCSV models.TravelCSV, id int64) (*models.TravelCSV, error) {

	if result := r.H.Where("id = ?", id).Updates(&travelCSV); result.Error != nil {
		return nil, status.Error(400, "Record not found")
	}

	return &travelCSV, nil
}

func (r *travelCSVRepository) Delete(id int64) (bool, error) {

	err := r.H.Where("id = ?", id).Delete(&models.TravelCSV{}).Error
	if err != nil {
		return false, err
	}

	return true, nil
}
