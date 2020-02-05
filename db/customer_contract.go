package db

import (
	"errors"
	"time"

	"github.com/golang191119/nc_crm/model"
)

func CreateContract(r *model.CustomerContract) error {
	r.CreatedAt = time.Now().In(Loc()).Format("2006-01-02 15:04:05")
	r.UpdatedAt = time.Now().In(Loc()).Format("2006-01-02 15:04:05")
	r.ID = int(time.Now().Unix())
	db := GetDB()
	defer db.Close()
	dbc := db.Create(&r)
	if dbc.Error != nil {
		return dbc.Error
	}
	if dbc.RowsAffected == 0 {
		return errors.New("Dont have anythings added")
	}
	return nil
}

func UpdateContract(id int, r *model.CustomerContract) error {
	db := GetDB()
	db.Close()
	dbc := db.Model(model.CustomerContract{}).Where("id = ?", id).Update(r)
	if dbc.Error != nil {
		return dbc.Error
	}
	if dbc.RowsAffected == 0 {
		return errors.New("Dont have anythings changed")
	}
	return nil
}

func GetAllContracts() ([]*model.CustomerContract, error) {
	contracts := []*model.CustomerContract{}
	db := GetDB()
	defer db.Close()
	dbc := db.Model(model.CustomerContract{}).Find(&contracts)
	if dbc.Error != nil {
		return nil, dbc.Error
	}
	return contracts, nil
}
