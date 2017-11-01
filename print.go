package number

import (
	"fmt"
)

var DefaultPrinter = new(MatrixPrinter)

func Println(args ...interface{}) (result string) {
	return DefaultPrinter.Println(args...)
}
func Print(args ...interface{}) (result string) {
	return DefaultPrinter.Print(args...)
}

type MatrixPrinter struct {
	Point  [][]interface{}
	Inputs []interface{}
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

	var l_now int
	for i := 0; i < len(mp.Inputs); i++ {
		switch mp.Inputs[i].(type) {
		case Matrix:
			m := mp.Inputs[i].(Matrix)
			h, l := mp.Inputs[i].(Matrix).Size()
			startx := l_now
			starty := (len(mp.Point) - h) / 2
			for p := 0; p < h; p++ {
				for q := 0; q < l; q++ {
					point := m[p][q]
					mp.Point[starty+p][startx+q] = point
				}
			}
			l_now += l
		case Matrixf:

			m := mp.Inputs[i].(Matrixf)
			h, l := mp.Inputs[i].(Matrixf).Size()

			startx := l_now
			starty := (len(mp.Point) - h) / 2
			for p := 0; p < h; p++ {
				for q := 0; q < l; q++ {
					point := m[p][q]
					//fmt.Println(heigth,length,starty,startx,p,q,starty+q,startx+p)
					mp.Point[starty+p][startx+q] = point
				}
			}
			l_now += l
		default:

			hh := 0
			if len(mp.Point)%2 == 0 {
				hh = len(mp.Point)/2 - 1
			} else {
				hh = len(mp.Point) / 2
			}
			mp.Point[hh][l_now] = mp.Inputs[i]
			l_now++
		}
	}
}
func (mp MatrixPrinter) String() (result string) {
	for i := 0; i < len(mp.Point); i++ {
		for j := 0; j < len(mp.Point[i]); j++ {
			if mp.Point[i][j] != nil {
				result += fmt.Sprint(mp.Point[i][j]) + "\t"
			} else {
				result += " \t"
			}

		}
		result += "\n"
	}
	result = result[:len(result)-2]
	return
}
func (mp MatrixPrinter) Print(args ...interface{}) (result string) {
	mp.Inputs = args
	mp.Point = make([][]interface{}, 0)
	mp.draw()
	result = mp.String()

	fmt.Print(result)
	return
}
func (mp MatrixPrinter) Println(args ...interface{}) (result string) {
	mp.Inputs = args
	mp.Point = make([][]interface{}, 0)
	mp.draw()
	result = mp.String()

	fmt.Println(result)
	return
}
