package main

func Max(x, y float64) float64 {
	if x < y {
		return y
	}
	return x
}

func Min(x, y float64) float64 {
	if x > y {
		return y
	}
	return x
}

func clamp(vaule float64, min float64, max float64) float64 {
	return Max(Min(vaule, max), min)
}
