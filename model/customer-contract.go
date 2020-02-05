package model

type CustomerContract struct {
	ID           int     `json:"id" gorm:"column:id"`
	CustomerID   int     `json:"customer_id" gorm:"column:customer_id"`
	IsActive     bool    `json:"is_active" gorm:"column:is_active"`
	CreatedAt    string  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt    string  `json:"updated_at" gorm:"column:updated_at"`
	CreatedBy    int     `json:"created_by" gorm:"column:created_by"`
	UpdatedBy    int     `json:"updated_by" gorm:"column:updated_by"`
	ContractType string  `json:"contract_type" gorm:"column:contract_type"`
	ExpiredDate  string  `json:"expired_date" gorm:"column:expired_date"`
	ReturnRatio  float32 `json:"return_ratio" gorm:"column:return_ratio"`
	CodRatio     float32 `json:"cod_ratio" gorm:"column:cod_ratio"`
	StopFee      float32 `json:"stop_fee" gorm:"column:stop_fee"`
	PaperFee     float32 `json:"paper_fee" gorm:"column:paper_fee"`
	LiftFee      float32 `json:"lift_fee" gorm:"column:lift_fee"`
	CheckFee     float32 `json:"check_fee" gorm:"column:check_fee"`
	ValueRatio   float32 `json:"value_ratio,omitempty" gorm:"column:value_ratio"`
}

func (c CustomerContract) TableName() string {
	return "customer_contract"
}
