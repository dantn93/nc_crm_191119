package db

import (
	"github.com/golang191119/nc_crm/model"
)

func GetTruckType() (*[]model.TruckType, error) {
	var truckTypes []model.TruckType
	db := GetDB()
	defer db.Close()
	if dbc := db.Find(&truckTypes); dbc.Error != nil {
		return nil, dbc.Error
	}
	return &truckTypes, nil
}
