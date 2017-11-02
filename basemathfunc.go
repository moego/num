package num

import "math"

//Sgn 阶越函数
func Sgn(x float64) float64 {
	if x >= 0 {
		return 1
	} else {
		return -1
	}
}

//Signoid函数 Signoid函数
func Signoid(x float64) float64 {
	return 1.0 / (1 + math.Exp(-1.0*x))
}
