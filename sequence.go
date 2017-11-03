package num

//Sequence 数列
type Sequence interface {
	//An 等差数列通项公式
	An() int
	//An 等差数列前n项和公式
	Sn() int
}

//Arithmeticsequence 等差数列
type Arithmeticsequence struct {
	A1, D int
}

//An 等差数列通项公式
func (A Arithmeticsequence) An(n int) int {
	if n < 1 {
		panic("n>=1,thanks")
	}
	return A.A1 + A.D*(n-1)
}
