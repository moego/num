package num

//Pi 莱布尼茨级数
func Pi(n int) (pi float64) {
	b := true
	for i := 1; i <= n; i++ {
		if b {
			pi += (1.0 / float64(2*i-1))
		} else {
			pi -= (1.0 / float64(2*i-1))
		}
		b = !b
	}
	return pi * 4.0
}
