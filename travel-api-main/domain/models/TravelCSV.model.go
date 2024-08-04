package models

type TravelCSV struct {
	Id                     int64  `json:"id" gorm:"primaryKey"`
	TravelledCountriesCode string `json:"travelled_countries_code" gorm:"column:travelled_countries_code;type:varchar(20)"`
	Total                  int64  `json:"total" gorm:"column:total;type:int;default:null"`
	Status                 bool   `json:"status" gorm:"column:status;type:boolean;default:null"`
}

func (TravelCSV) TableName() string {
	return "travel_from_csv"
}
