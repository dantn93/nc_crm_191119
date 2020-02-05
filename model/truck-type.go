package model

type TruckType struct {
	Name   string `json:"name" gorm:"column:name"`
	Weight int    `json:"weight" grom:"column:weight"`
}

func (c TruckType) TableName() string {
	return "truck_type"
}
