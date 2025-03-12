package processor

import (
	"app/models"
	"app/pkg/logger"
	"os"
	"strconv"
	"time"

	"math/rand"
)

func Execute(m models.Message) models.Report {

	rand.Seed(time.Now().UnixNano())

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

	logInternal := logger.Instance()

	jitter := os.Getenv("WORKERS_JITTER_MS")
	if jitter == "" {
		jitter = "0"
	}

	jitterInt, err := strconv.Atoi(jitter)
	if err != nil {
		logInternal.Fatal().
			Err(err).
			Str("WORKERS_JITTER_MS", os.Getenv("WORKERS_JITTER_MS")).
			Msg("Failed to paser WORKERS_JITTER_MS variable")
	}

	jitterDuration := time.Duration(jitterInt) * time.Millisecond

	var randomJitter time.Duration
	if jitterInt > 0 {
		randomJitter = time.Duration(rand.Intn(int(jitterDuration.Milliseconds()))) * time.Millisecond
	} else {
		randomJitter = 0
	}

	logInternal.Info().Any("JITTER_MS", randomJitter.Milliseconds()).Msg("Processing message")
	time.Sleep(randomJitter)

	// convert strint to time.Duration

	return report
}
