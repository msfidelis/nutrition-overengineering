package processor

import "app/models"

func Execute(m models.Message) models.Report {

	var report models.Report

	report.Id = m.Id
	report.Age = m.HealthInfo.Age
	report.Weight = m.HealthInfo.Weight
	report.Height = m.HealthInfo.Height
	report.Gender = m.HealthInfo.Gender
	report.ActivityIntensity = m.HealthInfo.ActivityIntensity
	report.Imc = m.Imc.Result
	report.ImcClass = m.Imc.Class
	report.Bmr = m.Basal.BMR.Value
	report.Water = m.Recomendations.Water.Value
	report.Protein = m.Recomendations.Protein.Value
	report.CaloriesToLoss = m.Recomendations.Calories.Loss.Value
	report.CaloriesToGain = m.Recomendations.Calories.Gain.Value
	report.CaloriesToMaintain = m.Recomendations.Calories.Maintain.Value

	return report
}
