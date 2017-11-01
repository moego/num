package number

import (
	"testing"
	"fmt"
)

func TestNewMatrix(t *testing.T) {
	m:=NewMatrix([]int{1},[]int{2},[]int{3,23,4})
	fmt.Println(m )
}
func TestRandMatrix(t *testing.T) {
	m:=RandMatrix(4,3,0,100)
	fmt.Println(m)
	fmt.Println(m.Size())
}

func TestRandMatrixf(t *testing.T) {
	m:=RandMatrixf(4,3,0.0,100.0)
	fmt.Println(m)
	fmt.Println(m.Size())
}

func TestNewMatrixf(t *testing.T) {
	m:=NewMatrixf([]float64{1.23},[]float64{2},[]float64{3,23,4})
	fmt.Println(m)
	fmt.Println(m.Size())

}

