package request

type LevelPrice struct {
	RateCardID int `json:"rate_card_id"`
	Prices     []struct {
		LevelID int     `json:"level_id"`
		Price   float64 `json:"price"`
	} `json:"prices"`
}
