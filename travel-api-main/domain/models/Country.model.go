package models

import (
	"time"

	"gorm.io/gorm"
)

type Country struct {
	Id                     int64          `json:"id" gorm:"primaryKey"`
	Name                   string         `json:"name" gorm:"column:name;type:varchar(250)"`
	Alpha2                 string         `json:"alpha_2" gorm:"column:alpha_2;type:varchar(20);default:null"`
	Alpha3                 string         `json:"alpha_3" gorm:"column:alpha_3;type:varchar(20);default:null"`
	CountryCode            string         `json:"country_code" gorm:"column:country_code;type:varchar(20);default:null"`
	Iso31662               string         `json:"iso_3166_2" gorm:"column:iso_3166_2;type:varchar(250);default:null"`
	Region                 string         `json:"region" gorm:"column:region;type:varchar(250);default:null"`
	SubRegion              string         `json:"sub_region" gorm:"column:sub_region;type:varchar(250);default:null"`
	IntermediateRegion     string         `json:"intermediate_region" gorm:"column:intermediate_region;type:varchar(250);default:null"`
	RegionCode             string         `json:"region_code" gorm:"column:region_code;type:varchar(250);default:null"`
	SubRegionCode          string         `json:"sub_region_code" gorm:"column:sub_region_code;type:varchar(250);default:null"`
	IntermediateRegionCode string         `json:"intermediate_region_code" gorm:"column:intermediate_region_code;type:varchar(250);default:null"`
	CreatedAt              time.Time      `json:"created_at"`
	UpdatedAt              time.Time      `json:"updated_at"`
	Deleted                gorm.DeletedAt `json:"deleted" gorm:"index"`
}

func (Country) TableName() string {
	return "country"
}
