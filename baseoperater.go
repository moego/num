package number

func DotC(m Matrix, b int) (mr Matrix) {
	x, y := m.Size()
	for i := 0; i < x; i++ {
		mr = append(mr, make([]int, y))
		for j := 0; j < y; j++ {
			mr[i][j] = b * m[i][j]
		}
	}
	return
}
func DotfC(m Matrixf, b float64) (mr Matrixf) {
	x, y := m.Size()
	for i := 0; i < x; i++ {
		mr = append(mr, make([]float64, y))
		for j := 0; j < y; j++ {
			mr[i][j] = b * m[i][j]
		}
	}
	return
}

//Doc 当矩阵A的列数等于矩阵B的行数时，A与B可以相乘。
func Dot(a, b Matrix) (mr Matrix) {

	x, p1 := a.Size()
	p2, y := b.Size()

	if p2 != p1 {
		panic("matrix size error")
	}

	for i := 0; i < x; i++ {
		mr = append(mr, make([]int, y))
		for j := 0; j < y; j++ {
			mr[i][j] = 0
			for k := 0; k < p2; k++ {
				if k < p2 {
					s := a[i][k] * b[k][j]
					mr[i][j] += s
				}

			}
		}
	}
	return
}
func Dotf(a, b Matrixf) (mr Matrixf) {

	x, p1 := a.Size()
	p2, y := b.Size()

	if p2 != p1 {
		panic("matrix size error")
	}

	for i := 0; i < x; i++ {
		mr = append(mr, make([]float64, y))
		for j := 0; j < y; j++ {
			mr[i][j] = 0
			for k := 0; k < p2; k++ {
				if k < p2 {
					s := a[i][k] * b[k][j]
					mr[i][j] += s
				}

			}
		}
	}
	return
}

func Sum(ms ...Matrix) (m Matrix) {
	if len(ms) == 0 {
		return
	} else if len(ms) == 1 {
		return ms[0]
	} else {
		l := len(ms)
		x, y := ms[0].Size()
		for ll := 0; ll < l; ll++ {
			x1, y1 := ms[ll].Size()
			if x1 != x || y1 != y {
				panic("matrix size error")
			}
		}

		for i := 0; i < x; i++ {
			m = append(m, make([]int, y))
			for j := 0; j < y; j++ {
				m[i][j] = 0
				for ll := 0; ll < l; ll++ {
					m[i][j] += ms[ll][i][j]
				}
			}
		}
	}
	return
}

func Sumf(ms ...Matrixf) (m Matrixf) {
	if len(ms) == 0 {
		return
	} else if len(ms) == 1 {
		return ms[0]
	} else {
		l := len(ms)
		x, y := ms[0].Size()
		for ll := 0; ll < l; ll++ {
			x1, y1 := ms[ll].Size()
			if x1 != x || y1 != y {
				panic("matrix size error")
			}
		}

		for i := 0; i < x; i++ {
			m = append(m, make([]float64, y))
			for j := 0; j < y; j++ {
				m[i][j] = 0
				for ll := 0; ll < l; ll++ {
					m[i][j] += ms[ll][i][j]
				}
			}
		}
	}
	return
}
