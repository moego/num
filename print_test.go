package number

import (
	"testing"
)

func TestMatrixPrinter_Print(t *testing.T) {
	p := &MatrixPrinter{}
	m := Eye(4)
	p.Print(2, "*", m, "=", DotC(m, 2))

}

func TestMatrixPrinter_Println(t *testing.T) {
	p := &MatrixPrinter{}
	m := Eyef(4)
	m2 := RandMatrixf(4, 2, 0, 10)
	p.Println(m, "*", m2, "=", Dotf(m, m2))
}
