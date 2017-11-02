package num

import (
	"fmt"
	"strconv"
)

//DefaultPrinter 默认打印
var DefaultPrinter = &MatrixPrinter{Tab: "\t"}

//Println 格式化输出矩阵计算公式,并换行
//
// Println(m1,"*", m2,"=", Dot(m1, m2))
//
//7	5	2
//
//3	1	8	*	4	2	5	=	89	57	11
//
//8	5	0
func Println(args ...interface{}) (result string) {
	return DefaultPrinter.Println(args...)
}

//Print 格式化输出矩阵计算公式,并换行
func Print(args ...interface{}) (result string) {
	return DefaultPrinter.Print(args...)
}

//MatrixPrinter 命令行打印工具
//Tab 设置打印的制表符
type MatrixPrinter struct {
	Tab    string
	Point  [][]interface{}
	Inputs []interface{}
}

//MaxCharlen 某列最长字符串的行序号，及长度
func (mp MatrixPrinter) MaxCharlen(length int) (max int) {
	for i := 0; i < len(mp.Point); i++ {
		str := []rune(fmt.Sprint(mp.Point[i][length]))
		if len(str) > max {
			max = len(str)

		}
	}
	return
}
func (mp *MatrixPrinter) resize(height, length int) *MatrixPrinter {
	for len(mp.Point) < height {
		mp.Point = append(mp.Point, make([]interface{}, length))
	}
	for i := 0; i < height; i++ {
		if len(mp.Point[i]) < length {
			mp.Point[i] = append(mp.Point[i], make([]interface{}, length-len(mp.Point[i]))...)
		}
	}
	return mp
}

func (mp *MatrixPrinter) draw() {
	var heigth, length int
	heigth = 1
	for i := 0; i < len(mp.Inputs); i++ {
		switch mp.Inputs[i].(type) {
		case Matrix:
			x, y := mp.Inputs[i].(Matrix).Size()
			length += y
			if heigth < x {
				heigth = x
			}
		case Matrixf:
			x, y := mp.Inputs[i].(Matrixf).Size()
			length += y
			if heigth < x {
				heigth = x
			}
		default:
			length++
		}
	}
	mp.resize(heigth, length)

	var lNow int
	for i := 0; i < len(mp.Inputs); i++ {
		switch mp.Inputs[i].(type) {
		case Matrix:
			m := mp.Inputs[i].(Matrix)
			h, l := mp.Inputs[i].(Matrix).Size()
			startx := lNow
			starty := (len(mp.Point) - h) / 2
			for p := 0; p < h; p++ {
				for q := 0; q < l; q++ {
					point := m[p][q]
					mp.Point[starty+p][startx+q] = point
				}
			}
			lNow += l
		case Matrixf:

			m := mp.Inputs[i].(Matrixf)
			h, l := mp.Inputs[i].(Matrixf).Size()

			startx := lNow
			starty := (len(mp.Point) - h) / 2
			for p := 0; p < h; p++ {
				for q := 0; q < l; q++ {
					point := m[p][q]
					//fmt.Println(heigth,length,starty,startx,p,q,starty+q,startx+p)
					mp.Point[starty+p][startx+q] = point
				}
			}
			lNow += l
		default:

			hh := 0
			if len(mp.Point)%2 == 0 {
				hh = len(mp.Point)/2 - 1
			} else {
				hh = len(mp.Point) / 2
			}
			mp.Point[hh][lNow] = mp.Inputs[i]
			lNow++
		}
	}
}

//String 计算MatrixPrinter的字符串
func (mp MatrixPrinter) String() (result string) {
	maxlen := []int{}
	for j := 0; j < len(mp.Point[0]); j++ {
		maxlen = append(maxlen, mp.MaxCharlen(j))
	}

	for i := 0; i < len(mp.Point); i++ {
		for j := 0; j < len(mp.Point[i]); j++ {
			if mp.Point[i][j] != nil {
				result += fmt.Sprintf("%-"+strconv.Itoa(maxlen[j])+"s", fmt.Sprint(mp.Point[i][j])) + mp.Tab
			} else {
				result += " \t"
			}

		}
		result += "\n"
	}
	result = result[:len(result)-2]
	return
}

//Print 格式化输出矩阵计算公式
func (mp MatrixPrinter) Print(args ...interface{}) (result string) {
	mp.Inputs = args
	mp.Point = make([][]interface{}, 0)
	mp.draw()
	result = mp.String()

	fmt.Print(result)
	return
}

//Println 格式化输出矩阵计算公式,并换行
func (mp MatrixPrinter) Println(args ...interface{}) (result string) {
	mp.Inputs = args
	mp.Point = make([][]interface{}, 0)
	mp.draw()
	result = mp.String()

	fmt.Println(result)
	return
}
