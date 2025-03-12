package models

type Report struct {
	Id                 string  `json:"id" dynamodbav:"id" binding:"required" `
	Age                int     `json:"age" dynamodbav:"age" binding:"required"`
	Weight             float64 `json:"weight" dynamodbav:"weight" binding:"required"`
	Height             float64 `json:"height" dynamodbav:"height" binding:"required"`
	Gender             string  `json:"gender" dynamodbav:"gender" binding:"required,oneof=M F"`
	ActivityIntensity  string  `json:"activity_intensity" dynamodbav:"activity_intensity" binding:"required"`
	Imc                float64 `json:"imc" dynamodbav:"imc" binding:"required"`
	ImcClass           string  `json:"imc_class" dynamodbav:"imc_class" binding:"required"`
	Bmr                float64 `json:"bmr" dynamodbav:"bmr" binding:"required"`
	Necessity          float64 `json:"necessity" dynamodbav:"necessity" binding:"required"`
	Protein            int64   `json:"protein" dynamodbav:"protein" binding:"required"`
	Water              float64 `json:"water" dynamodbav:"water" binding:"required"`
	CaloriesToLoss     float64 `json:"calories_to_loss" dynamodbav:"calories_to_loss" binding:"required"`
	CaloriesToGain     float64 `json:"calories_to_gain" dynamodbav:"calories_to_gain" binding:"required"`
	CaloriesToMaintain float64 `json:"calories_to_maintain" dynamodbav:"calories_to_maintain" binding:"required"`
}
