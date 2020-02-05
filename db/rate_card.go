package db

import (
	"errors"
	"time"

	"github.com/golang191119/nc_crm/model"
)

func CreateRateCard(r *model.TruckRateCard) error {
	r.UpdatedAt = time.Now().In(Loc()).Format("2006-01-02 15:04:05")
	r.CreatedAt = time.Now().In(Loc()).Format("2006-01-02 15:04:05")
	r.ID = int(time.Now().Unix())
	db := GetDB()
	defer db.Close()
	dbc := db.Create(&r)
	if dbc.Error != nil {
		return dbc.Error
	}
	return nil
}

func UpdateRateCard(id int, r *model.TruckRateCard) error {
	db := GetDB()
	defer db.Close()
	dbc := db.Model(model.TruckRateCard{}).Where("id = ?", id).Update(r)
	if dbc.Error != nil {
		return dbc.Error
	}
	if dbc.RowsAffected == 0 {
		return errors.New("Dont have anythings changed")
	}
	return nil
}

func DeleteRateCard(id int) error {
	// Note the use of tx as the database handle once you are within a transaction
	tx := GetDB().Begin()
	defer tx.Close()
	if tx.Error != nil {
		return tx.Error
	}

	if dbc := tx.Where("rate_card_id = ? ", id).Delete(&model.TruckRateCardLevel{}); dbc.Error != nil {
		tx.Rollback()
		return dbc.Error
	}

	if dbc := tx.Where("id = ?", id).Delete(&model.TruckRateCard{}); dbc.Error != nil {
		tx.Rollback()
		return dbc.Error
	}
	return tx.Commit().Error
}
