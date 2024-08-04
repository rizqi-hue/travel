package models

import (
	"time"

	"gorm.io/gorm"
)

type Travel struct {
	Id                     int64          `json:"id" gorm:"primaryKey"`
	TravelledCountriesCode string         `json:"travelled_countries_code" gorm:"column:travelled_countries_code;type:varchar(20)"`
	Country                Country        `gorm:"foreignKey:alpha_3;references:travelled_countries_code"`
	Total                  int64          `json:"total" gorm:"column:total;type:int;default:null"`
	CreatedAt              time.Time      `json:"created_at"`
	UpdatedAt              time.Time      `json:"updated_at"`
	Deleted                gorm.DeletedAt `json:"deleted" gorm:"index"`
}

type Tabler interface {
	TableName() string
}

func (Travel) TableName() string {
	return "travel"
}
