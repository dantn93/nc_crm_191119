package db

import (
	"fmt"

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

func GetTruckLevel() (*[]model.TruckLevel, error) {
	var truckLevels []model.TruckLevel
	db := GetDB()
	defer db.Close()
	if dbc := db.Find(&truckLevels); dbc.Error != nil {
		fmt.Println("\n\n\n\n\n\n\n\n")
		fmt.Println(dbc.Error)
		return nil, dbc.Error
	}
	return &truckLevels, nil
}
