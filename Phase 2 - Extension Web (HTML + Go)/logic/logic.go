package logic

func Kilometers2Miles(nb float64) float64 {
	nb = nb / 1.609344
	return nb
}

func Kilograms2pounds(nb float64) float64 {
	nb = nb / 0.45359237
	return nb
}

func Celsius2Fahreneit(nb float64) float64 {
	nb = (nb*9)/5 + 32
	return nb
}
