package utils

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/golang191119/nc_crm/db"
	"github.com/golang191119/nc_crm/model"
	"github.com/golang191119/nc_crm/model/request"
)

func TimeFormat(t time.Time) string {
	month := strconv.Itoa(int(t.Month()))
	if len(month) == 1 {
		month = "0" + month
	}

	day := strconv.Itoa(t.Day())
	if len(day) == 1 {
		day = "0" + day
	}

	hour := strconv.Itoa(t.Hour())
	if len(hour) == 1 {
		hour = "0" + hour
	}

	minute := strconv.Itoa(t.Minute())
	if len(minute) == 1 {
		minute = "0" + minute
	}

	second := strconv.Itoa(t.Second())
	if len(second) == 1 {
		second = "0" + second
	}

	str := fmt.Sprintf("%v-%v-%v %v:%v:%v", t.Year(), month, day, hour, minute, second)
	fmt.Println("TIME STAMP: ", str)
	return str
}

func CalBasePrice(contractID int, truckType int, distance int) (float64, error) {
	var total float64
	var ratePrice []model.RatePrice
	database := db.GetDB()

	selectItems := "truck_level.rate, truck_level.code, truck_rate_card_level.price"
	dbc := database.Table("truck_level").Joins("JOIN truck_rate_card_level ON truck_level.id = truck_rate_card_level.level_id").Joins("JOIN truck_rate_card ON truck_rate_card.id = truck_rate_card_level.rate_card_id").Joins("JOIN customer_contract ON customer_contract.id = truck_rate_card.contract_id").Where("customer_contract.id = ? AND truck_rate_card.truck_type = ?", contractID, truckType).Order("truck_level.rate asc").Select(selectItems)
	defer database.Close()

	if dbc.Error != nil {
		return 0, dbc.Error
	}
	rows, err := dbc.Rows()
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	for rows.Next() {
		var rate int
		var code string
		var price float64
		err = rows.Scan(&rate, &code, &price)
		if err != nil {
			return 0, err
		}

		ratePrice = append(ratePrice, model.RatePrice{
			Rate:  rate,
			Code:  code,
			Price: price,
		})
	}

	if len(ratePrice) == 0 {
		return 0, errors.New("Rate card level is not found!")
	}

	if len(ratePrice) == 1 {
		return float64(distance) * ratePrice[0].Price, nil
	}

	// Calculate base price by distance steps
	for i := 0; i < len(ratePrice)-1; i++ {
		if distance <= 0 {
			break
		}
		if distance <= ratePrice[i+1].Rate {
			total += float64(distance) * ratePrice[i].Price
			distance -= ratePrice[i+1].Rate
		} else {
			distance -= ratePrice[i+1].Rate
			total += float64(ratePrice[i+1].Rate) * ratePrice[i].Price
		}
	}
	if distance > 0 {
		total += ratePrice[len(ratePrice)-1].Price * float64(distance)
	}

	return math.Ceil(total), nil
}

func CalFee(contractID int, r *request.Consumption) (float64, error) {
	var contract model.CustomerContract
	database := db.GetDB()
	defer database.Close()
	dbc := database.Model(model.CustomerContract{}).Where("id = ?", contractID).Find(&contract)
	if dbc.Error != nil {
		return 0, dbc.Error
	}
	var total float64
	// RETURN FEE
	if r.ReturnAmount > 0 {
		total += float64(r.ReturnAmount) * float64(contract.ReturnRatio)
	}
	// COD
	if r.Cod > 0 {
		total += float64(r.Cod) * float64(contract.CodRatio)
	}
	// STOP COUNT
	if r.StopCount > 0 {
		total += float64(r.StopCount) * float64(contract.StopFee)
	}
	// PAPER
	if r.PaperCount > 0 {
		total += float64(r.PaperCount) * float64(contract.PaperFee)
	}
	// LIFTING
	if r.LiftingCount > 0 {
		total += float64(r.LiftingCount) * float64(contract.LiftFee)
	}
	// CHECKING
	if r.CheckingCount > 0 {
		total += float64(r.CheckingCount) * float64(contract.CheckFee)
	}
	// VALUE (Khai gia)
	if r.Value > 0 {
		total += r.Value * float64(contract.ValueRatio)
	}

	return math.Ceil(total), nil
}
