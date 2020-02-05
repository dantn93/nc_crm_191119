package db

import (
	"errors"
	"time"

	"github.com/golang191119/nc_crm/model"
	request "github.com/golang191119/nc_crm/model/request"
)

func AddRateCardLevel(r *request.LevelPrice) error {
	tx := GetDB().Begin()
	defer tx.Close()
	if tx.Error != nil {
		return tx.Error
	}

	for i, v := range r.Prices {
		rateCardLevelInstance := model.TruckRateCardLevel{
			ID:         int(time.Now().Unix()) + i,
			RateCardID: r.RateCardID,
			LevelID:    v.LevelID,
			Price:      v.Price,
		}
		if dbc := tx.Create(&rateCardLevelInstance); dbc.Error != nil {
			tx.Rollback()
			return dbc.Error
		}
	}

	return tx.Commit().Error
}

func DeleteRateCardLevel(rateCardID int) error {
	db := GetDB()
	defer db.Close()
	dbc := db.Where("rate_card_id = ?", rateCardID).Delete(&model.TruckRateCardLevel{})
	if dbc.Error != nil {
		return dbc.Error
	}
	if dbc.RowsAffected == 0 {
		return errors.New("Dont have anythings deleted")
	}
	return nil
}

func UpdateRateCardLevel(r *request.LevelPrice) error {
	tx := GetDB().Begin()
	defer tx.Close()
	if tx.Error != nil {
		return tx.Error
	}
	for _, v := range r.Prices {
		dataUpdate := model.TruckRateCardLevel{
			Price: v.Price,
		}
		dbc := tx.Model(model.TruckRateCardLevel{}).Where("rate_card_id = ? AND level_id = ?", r.RateCardID, v.LevelID).Update(&dataUpdate)
		if dbc.Error != nil {
			tx.Rollback()
			return dbc.Error
		}
	}
	return tx.Commit().Error
}
