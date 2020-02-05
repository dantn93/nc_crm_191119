package request

import "github.com/golang191119/nc_crm/model"

type TruckLevel struct {
	Tag    string              `json:"tag"`
	Levels []*model.TruckLevel `json:"levels"`
}
