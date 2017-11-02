package num

import (
	"testing"
)

func TestNewMatrix(t *testing.T) {
	m := NewMatrix([]int{1}, []int{2}, []int{3, 23, 4})
	Println(m)
}
func TestRandMatrix(t *testing.T) {
	m := RandMatrix(4, 3, 0, 100)
	Println(m)
	Println(m.Size())
}

func TestRandMatrixf(t *testing.T) {
	m := RandMatrixf(4, 3, 0.0, 100.0)
	Println(m)
	Println(m.Size())
}

func TestNewMatrixf(t *testing.T) {
	m := NewMatrixf([]float64{1.23, 2}, []float64{2.1}, []float64{3, 23.23, 4})
	Println(m)
	Println(m.Size())
}

func TestEye(t *testing.T) {
	Println(Eye(4))
}
func TestEyef(t *testing.T) {
	Println(Eyef(4))
}
