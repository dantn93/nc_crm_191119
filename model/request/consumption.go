package request

type Consumption struct {
	ContractID    int     `json:"contract_id"`
	TruckType     int     `json:"truck_type"`
	Distance      int     `json:"distance"`      //km
	ReturnAmount  float32 `json:"return_amount"` //vnd
	Cod           float64 `json:"code_amount"`   //vnd
	StopCount     int     `json:"stop_count"`
	PaperCount    int     `json:"paper_count"`
	LiftingCount  int     `json:"lifting_count"`
	CheckingCount int     `json:"checking_count"`
	Value         float64 `json:"value"` //vnd
}
