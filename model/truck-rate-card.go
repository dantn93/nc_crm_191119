package model

type TruckRateCard struct {
	ID         int    `json:"id" gorm:"column:id"`
	ContractID int    `json:"contract_id" gorm:"column:contract_id"`
	TruckType  int    `json:"truck_type" gorm:"column:truck_type"`
	UpdatedAt  string `json:"updated_at" gorm:"column:updated_at"`
	CreatedAt  string `json:"created_at" gorm:"column:created_at"`
	CreatedBy  int    `json:"created_by" gorm:"column:created_by"`
}

func (c TruckRateCard) TableName() string {
	return "truck_rate_card"
}
