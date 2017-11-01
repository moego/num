package num

import (
	"testing"
)

func TestMatrixPrinter_Print(t *testing.T) {

	m := Eye(4)
	Print(2, "*", m, "=", DotC(m, 2))

}

func TestMatrixPrinter_Println(t *testing.T) {

	m := Eyef(4)
	m2 := RandMatrixf(4, 2, 0, 10)
	Println(m, "*", m2, "=", Dotf(m, m2))
}
