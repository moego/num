package num

//Factorial 阶乘
func Factorial(n int) (f int) {
	f = 1
	for i := 1; i <= n; i++ {
		f *= i
	}
	return
}
