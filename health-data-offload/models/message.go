package models

type Message struct {
	Id     string `json:"id" binding:"required"`
	Status int    `json:"status" binding:"required"`
	Imc    struct {
		Result float64 `json:"result"`
		Class  string  `json:"class"`
	} `json:"imc"`
	Basal struct {
		BMR struct {
			Value float64 `json:"value"`
			Unit  string  `json:"unit"`
		} `json:"bmr"`
		Necessity struct {
			Value float64 `json:"value"`
			Unit  string  `json:"unit"`
		} `json:"necessity"`
	} `json:"basal"`
	HealthInfo struct {
		Age               int     `json:"age"`
		Weight            float64 `json:"weight"`
		Height            float64 `json:"height"`
		Gender            string  `json:"gender"`
		ActivityIntensity string  `json:"activity_intensity"`
	} `json:"health_info"`
	Recomendations struct {
		Protein struct {
			Value int64  `json:"value"`
			Unit  string `json:"unit"`
		} `json:"protein"`
		Water struct {
			Value float64 `json:"value"`
			Unit  string  `json:"unit"`
		} `json:"water"`
		Calories struct {
			Maintain struct {
				Value float64 `json:"value"`
				Unit  string  `json:"unit"`
			} `json:"maintain_weight"`
			Loss struct {
				Value float64 `json:"value"`
				Unit  string  `json:"unit"`
			} `json:"loss_weight"`
			Gain struct {
				Value float64 `json:"value"`
				Unit  string  `json:"unit"`
			} `json:"gain_weight"`
		} `json:"calories"`
	} `json:"recomendations"`
}
