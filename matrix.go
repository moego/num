package number

import (
	"strconv"
)

//Matrix  矩阵
type Matrix [][]int

//Matrixf  浮点矩阵
type Matrixf [][]float64

//autosize  对人工输入的矩阵初始数据，进行矩阵补齐
func (matrix *Matrix) autosize() *Matrix {
	height, length := matrix.Size()
	for len((*matrix)) < height {
		(*matrix) = append((*matrix), make([]int, length))
	}
	for i := 0; i < height; i++ {
		if len((*matrix)[i]) < length {
			(*matrix)[i] = append((*matrix)[i], make([]int, length-len((*matrix)[i]))...)
		}
	}
	return matrix
}

//autosize  对人工输入的矩阵初始数据，进行矩阵补齐
func (matrixf *Matrixf) autosize() *Matrixf {
	height, length := matrixf.Size()
	for len((*matrixf)) < height {
		(*matrixf) = append((*matrixf), make([]float64, length))
	}
	for i := 0; i < height; i++ {
		if len((*matrixf)[i]) < length {
			(*matrixf)[i] = append((*matrixf)[i], make([]float64, length-len((*matrixf)[i]))...)
		}
	}
	return matrixf
}

//RandMatrix  生成随机矩阵
//矩阵高height,宽length
//随机数范围,duemin< mynumber <duemax
func RandMatrix(height, length int, duemin, duemax int) (m Matrix) {
	for x := 0; x < height; x++ {
		m = append(m, make([]int, length))
		for y := 0; y < length; y++ {
			m[x][y] = RandInt(duemin, duemax)
		}
	}
	return
}

//RandMatrixf 生成随机矩阵
//矩阵高height,宽length
//随机数范围,duemin< mynumber <duemax
func RandMatrixf(lenx, leny int, duemin, duemax float64) (m Matrixf) {
	for x := 0; x < lenx; x++ {
		m = append(m, make([]float64, leny))
		for y := 0; y < leny; y++ {
			m[x][y] = RandFloat(duemin, duemax)
		}
	}

	return
}

//Size 矩阵高宽
//矩阵高height,宽length
func (matrix Matrix) Size() (height, length int) {
	height = len(matrix)
	for _, line := range matrix {
		if len(line) > length {
			length = len(line)
		}
	}
	return
}

//Size 矩阵高宽
//矩阵高height,宽length
func (matrixf Matrixf) Size() (height, length int) {
	height = len(matrixf)
	for _, line := range matrixf {
		if len(line) > length {
			length = len(line)
		}
	}
	return
}

//String 打印矩阵
func (matrixf Matrixf) String() string {

	s := "\n"
	for x := 0; x < len(matrixf); x++ {
		for y := 0; y < len(matrixf[x])-1; y++ {
			s += strconv.FormatFloat(matrixf[x][y], 'f', -1, 64) + "\t"
		}
		s += strconv.FormatFloat(matrixf[x][len(matrixf[x])-1], 'f', -1, 64) + "\n"
	}
	return s

}

//String 打印矩阵
func (matrix Matrix) String() string {

	s := "\n"
	for x := 0; x < len(matrix); x++ {
		for y := 0; y < len(matrix[x])-1; y++ {
			s += strconv.Itoa(matrix[x][y]) + "\t"
		}
		s += strconv.Itoa(matrix[x][len(matrix[x])-1]) + "\n"
	}
	return s

}

//NewMatrix new矩阵
func NewMatrix(datas ...[]int) Matrix {
	result := make(Matrix, len(datas))
	for index, data := range datas {
		result[index] = data
	}
	result.autosize()
	return result
}

//NewMatrixf new矩阵
func NewMatrixf(datas ...[]float64) Matrixf {
	result := make(Matrixf, len(datas))
	for index, data := range datas {
		result[index] = data
	}
	result.autosize()
	return result
}

//Eye new单位矩阵
func Eye(n int) (m Matrix) {
	for x := 0; x < n; x++ {
		m = append(m, make([]int, n))
		m[x][x] = 1
	}
	return
}

//Eyef new单位矩阵
func Eyef(n int) (m Matrixf) {
	for x := 0; x < n; x++ {
		m = append(m, make([]float64, n))
		m[x][x] = 1
	}
	return
}

//Sum  矩阵加法
func (matrix Matrix) Sum(b ...Matrix) Matrix {
	c := append(b, matrix)
	return Sum(c...)
}

//Sumf  矩阵加法
func (matrixf Matrixf) Sumf(b ...Matrixf) Matrixf {
	c := append(b, matrixf)
	return Sumf(c...)
}

//Dot  矩阵乘法
func (matrix Matrix) Dot(b Matrix) Matrix {
	return Dot(matrix, b)
}

//Dotf  矩阵乘法
func (matrixf Matrixf) Dotf(b Matrixf) Matrixf {
	return Dotf(matrixf, b)
}
