package db

import (
	"time"

	"github.com/golang191119/nc_crm/model"
	request "github.com/golang191119/nc_crm/model/request"
)

func AddTruckLevel(r *request.TruckLevel) error {
	tx := GetDB().Begin()
	defer tx.Close()
	if tx.Error != nil {
		return tx.Error
	}

	for i, v := range r.Levels {
		truckLevelInstance := model.TruckLevel{
			ID:   int(time.Now().Unix()) + i,
			Rate: v.Rate,
			Code: v.Code,
			Name: v.Name,
			Tag:  r.Tag,
		}
		if dbc := tx.Create(&truckLevelInstance); dbc.Error != nil {
			tx.Rollback()
			return dbc.Error
		}
	}
	return tx.Commit().Error
}

func UpdateTruckLevel(r []*model.TruckLevel) error {
	tx := GetDB().Begin()
	defer tx.Close()
	if tx.Error != nil {
		return tx.Error
	}

	for _, v := range r {
		dataUpdate := model.TruckLevel{
			Rate: v.Rate,
			Code: v.Code,
			Name: v.Name,
		}
		if dbc := tx.Model(model.TruckLevel{}).Where("id = ?", v.ID).Update(&dataUpdate); dbc.Error != nil {
			tx.Rollback()
			return dbc.Error
		}
	}
	return tx.Commit().Error
}
