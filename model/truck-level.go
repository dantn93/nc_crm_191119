package model

type TruckLevel struct {
	ID   int    `json:"id" gorm:"column:id"`
	Rate int    `json:"rate" gorm:"column:rate"`
	Code string `json:"code" gorm:"column:code"`
	Name string `json:"name" gorm:"column:name"`
	Tag  string `json:"tag" gorm:"column:tag"`
}

func (c TruckLevel) TableName() string {
	return "truck_level"
}
