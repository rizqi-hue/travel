package repository

import (
	"github.com/travel-api-main/domain/models"
	"github.com/travel-api-main/pkg/database"
	"github.com/travel-api-main/pkg/pb"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type CountryRepository interface {
	Gets(req *pb.GetsCountryRequest) (countrys []models.Country, totalRow int64, totalPage int64, err error)
	Insert(country models.Country) (*models.Country, error)
	Update(country models.Country, id int64) (*models.Country, error)
	Delete(id int64) (bool, error)
	GetById(req *pb.GetCountryRequest) (*models.Country, error)
}

type countryRepository struct {
	H *gorm.DB
}

func NewCountryRepository() *countryRepository {
	return &countryRepository{
		H: database.DB,
	}
}

func FilterCountry(country *models.Country) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		DB := db

		// if country.UnikId != "" {
		// 	DB.Where("unik_id = ?", country.UnikId)
		// }

		return DB
	}
}

func (r *countryRepository) Gets(req *pb.GetsCountryRequest) (countrys []models.Country, totalRow int64, totalPage int64, err error) {

	page := req.GETS_COUNTRY.Page
	size := req.GETS_COUNTRY.Size

	if result := r.H.Scopes(
		Paginate(page, size),
		FilterCountry(&models.Country{})).
		Find(&countrys); result.Error != nil {
		return nil, 0, 0, status.Error(400, "error gets country")
	}

	if result := r.H.Scopes(FilterCountry(&models.Country{})).Model(models.Country{}).Count(&totalRow); result.Error != nil {
		return nil, 0, 0, status.Error(400, "error count country")
	}

	if size == 0 {
		size = 1
	}

	totalPage = (totalRow / size) + 1

	return countrys, totalRow, totalPage, nil
}

func (r *countryRepository) Insert(country models.Country) (*models.Country, error) {

	err := r.H.Save(&country).Error
	if err != nil {
		return nil, err
	}

	return &country, nil
}

func (r *countryRepository) Update(country models.Country, id int64) (*models.Country, error) {

	if result := r.H.Where("id = ?", id).Updates(&country); result.Error != nil {
		return nil, status.Error(400, "Record not found")
	}

	return &country, nil
}

func (r *countryRepository) Delete(id int64) (bool, error) {

	err := r.H.Where("id = ?", id).Delete(&models.Country{}).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *countryRepository) GetById(req *pb.GetCountryRequest) (country *models.Country, err error) {

	if result := r.H.
		Scopes(
			FilterCountry(&models.Country{
				Id: req.GET_COUNTRY.Id,
			})).First(&country); result.Error != nil {
		return nil, status.Error(400, "Record not found")
	}

	return country, nil
}
