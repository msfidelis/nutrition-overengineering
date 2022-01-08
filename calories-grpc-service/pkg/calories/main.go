package calories

func Maintain(necessity float64) float64 {
	return necessity
}

func Loss(necessity float64) float64 {
	return necessity * 0.90
}

func Gain(necessity float64) float64 {
	return necessity * 1.30
}
