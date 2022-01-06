package imc

func Calc(weight float64, height float64) (float64, string) {
	var imc float64
	var class string

	imc = (weight / (height * height)) * 1

	if imc <= 18.5 {
		class = "underweight"
	} else if imc > 18.5 && imc < 24.9 {
		class = "healthy"
	} else if imc >= 25 && imc <= 29.9 {
		class = "overweight"
	} else if imc > 30 {
		class = "obese"
	}

	return imc, class
}
