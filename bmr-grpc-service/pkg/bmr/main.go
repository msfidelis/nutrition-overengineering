package bmr

import "strings"

func Calc(gender string, weight float64, height float64, age int64, activity string) (float64, float64) {
	var basal float64
	var necessity float64
	if strings.ToUpper(gender) == "M" {
		basal = 66 + (13.75 * weight) + (5.0 * (height * 100)) - (6.8 * float64(age))
	} else {
		basal = 665 + (9.56 * weight) + (1.8 * (height * 100)) - (4.7 * float64(age))
	}

	// Calc Necessity

	switch strings.ToLower(activity) {
	case "sedentary":
		necessity = basal * 1.2
	case "lightly_active":
		necessity = basal * 1.375
	case "moderately_active":
		necessity = basal * 1.55
	case "very_active":
		necessity = basal * 1.725
	case "extra_active":
		necessity = basal * 1.9
	}

	return basal, necessity
}
