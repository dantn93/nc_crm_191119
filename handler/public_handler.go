package handler

import (
	"net/http"

	"github.com/golang191119/nc_crm/db"
	"github.com/golang191119/nc_crm/model"
	request "github.com/golang191119/nc_crm/model/request"
	utils "github.com/golang191119/nc_crm/utils"
	"github.com/labstack/echo/v4"
)

func HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

// ==================== CONTRACT =========================//
func GetAllContracts(c echo.Context) error {
	contracts, err := db.GetAllContracts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, contracts)
}
func AddContract(c echo.Context) error {
	var r model.CustomerContract
	if err := c.Bind(&r); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Check input
	if r.CustomerID == 0 {
		return c.JSON(http.StatusBadRequest, "customer_id is required!")
	}
	if r.CreatedBy == 0 {
		return c.JSON(http.StatusBadRequest, "created_by is required!")
	}
	if r.ContractType == "" {
		return c.JSON(http.StatusBadRequest, "contract_type is required!")
	}
	if r.ExpiredDate == "" {
		return c.JSON(http.StatusBadRequest, "expired_date is required!")
	}
	if r.ReturnRatio == 0 {
		return c.JSON(http.StatusBadRequest, "return_ratio is required!")
	}
	if r.CodRatio == 0 {
		return c.JSON(http.StatusBadRequest, "cod_ratio is required!")
	}
	if r.StopFee == 0 {
		return c.JSON(http.StatusBadRequest, "stop_fee is required!")
	}
	if r.PaperFee == 0 {
		return c.JSON(http.StatusBadRequest, "paper_fee is required!")
	}
	if r.LiftFee == 0 {
		return c.JSON(http.StatusBadRequest, "lift_fee is required!")
	}
	if r.CheckFee == 0 {
		return c.JSON(http.StatusBadRequest, "check_fee is required!")
	}
	if r.ValueRatio == 0 {
		return c.JSON(http.StatusBadRequest, "value_ratio is required!")
	}

	// Create contract
	if err := db.CreateContract(&r); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, "Contract is created successfully")
}

func UpdateContract(c echo.Context) error {
	var r model.CustomerContract
	if err := c.Bind(&r); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Check input
	if r.ID == 0 {
		return c.JSON(http.StatusBadRequest, "id (contractID) is required!")
	}
	if r.UpdatedBy == 0 {
		return c.JSON(http.StatusBadRequest, "updated_by is required!")
	}

	// Update contract
	contractID := r.ID
	r.ID = 0
	if err := db.UpdateContract(contractID, &r); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, "Contract is updated successfully")
}

// ===================== RATE CARD ======================//
func AddRateCard(c echo.Context) error {
	var r model.TruckRateCard
	if err := c.Bind(&r); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Check request
	if r.ContractID == 0 {
		return c.JSON(http.StatusBadRequest, "contract_id is required!")
	}
	if r.TruckType == 0 {
		return c.JSON(http.StatusBadRequest, "truct_type is required!")
	}
	if r.CreatedBy == 0 {
		return c.JSON(http.StatusBadRequest, "created_by is required!")
	}

	// Create
	if err := db.CreateRateCard(&r); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, "Rate card is created successfully")
}

func UpdateRateCard(c echo.Context) error {
	var r model.TruckRateCard
	if err := c.Bind(&r); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	// Check request
	if r.ID == 0 {
		return c.JSON(http.StatusBadRequest, "id is required!")
	}

	// Update
	id := r.ID
	r.ID = 0
	if err := db.UpdateRateCard(id, &r); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, "Rate card is updated successfully")
}

func DeleteRateCard(c echo.Context) error {
	var r model.TruckRateCard
	if err := c.Bind(&r); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Check request
	if r.ID == 0 {
		return c.JSON(http.StatusBadRequest, "id is required!")
	}

	// Delete
	if err := db.DeleteRateCard(r.ID); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, "Rate card is deleted successfully")
}

// ===================== RATE CARD LEVEL =============//
func AddRateCardLevel(c echo.Context) error {
	var r request.LevelPrice
	if err := c.Bind(&r); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if r.RateCardID == 0 {
		return c.JSON(http.StatusBadRequest, "rate_card_id is required!")
	}
	if r.Prices == nil {
		return c.JSON(http.StatusBadRequest, "price is required. It is a descrease integer-array!")
	}

	// Insert
	if err := db.AddRateCardLevel(&r); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, "Truck rate card level is added successfully")
}

func DeleteRateCardLevel(c echo.Context) error {
	var r model.TruckRateCardLevel

	if err := c.Bind(&r); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if r.RateCardID == 0 {
		return c.JSON(http.StatusBadRequest, "rate_card_id is required!")
	}

	// Delete
	if err := db.DeleteRateCardLevel(r.RateCardID); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Rate card level are deleted successfully")

}

func UpdateRateCardLevel(c echo.Context) error {
	var r request.LevelPrice
	if err := c.Bind(&r); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	// Check requets
	if r.RateCardID == 0 {
		return c.JSON(http.StatusBadRequest, "rate_card_id is required!")
	}
	if r.Prices == nil {
		return c.JSON(http.StatusBadRequest, "prices is required!")
	}

	// Update
	err := db.UpdateRateCardLevel(&r)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "Rate card level are updated successfully")
}

// ================ TRUCK LEVEL =============== //
func GetTruckLevel(c echo.Context) error {
	truckLevels, err := db.GetTruckLevel()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, truckLevels)
}

func AddTruckLevel(c echo.Context) error {
	var r request.TruckLevel
	if err := c.Bind(&r); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Check request
	if r.Tag == "" {
		return c.JSON(http.StatusBadRequest, "tag is required!")
	}
	if r.Levels == nil {
		return c.JSON(http.StatusBadRequest, "truck level array is required!")
	}

	if err := db.AddTruckLevel(&r); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Truck levels are added successfully")
}

func UpdateTruckLevel(c echo.Context) error {
	var r []*model.TruckLevel
	if err := c.Bind(&r); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Check request
	if r == nil {
		return c.JSON(http.StatusBadRequest, "Truck level array is required!")
	}
	for _, v := range r {
		if v.ID == 0 {
			return c.JSON(http.StatusBadRequest, "Any of truck level ids are required!")
		}
	}
	// Update
	if err := db.UpdateTruckLevel(r); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Truck levels are updated successfully!")
}

// ================ TRUCK TYPE ================ //
func GetTruckType(c echo.Context) error {
	truckTypes, err := db.GetTruckType()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, truckTypes)
}

// ================== PRICE ==================== //
func GetPrice(c echo.Context) error {
	var r request.Consumption
	if err := c.Bind(&r); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Check request
	if r.ContractID == 0 {
		return c.JSON(http.StatusBadRequest, "customer_id is required!")
	}
	if r.Distance < 0 {
		return c.JSON(http.StatusBadRequest, "distance is required!")
	}
	if r.PaperCount < 0 {
		return c.JSON(http.StatusBadRequest, "paper_count is required!")
	}

	// Calculate price
	fee, err := utils.CalFee(r.ContractID, &r)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	basePrice, err := utils.CalBasePrice(r.ContractID, r.TruckType, r.Distance)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"base_price": basePrice,
		"fee":        fee,
		"total":      basePrice + fee,
	})
}
