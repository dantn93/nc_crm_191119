package model

type TruckRateCardLevel struct {
	ID         int     `json:"id" grom:"column:id"`
	RateCardID int     `json:"rate_card_id" gorm:"column:rate_card_id"`
	LevelID    int     `json:"level_id" gorm:"column:level_id"`
	Price      float64 `json:"price" gorm:"column:price"`
}

func (c TruckRateCardLevel) TableName() string {
	return "truck_rate_card_level"
}
