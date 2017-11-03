// This example
package num

import (
	"fmt"
	"testing"
)

func ExampleMatrix() {
	A := NewMatrix([]int{1, 2}, []int{2, 3}, []int{3, 4})
	B := Eye(2)
	Println(A, ".", B, "=", Dot(A, B))
	//Output:
	//1	2	 	1	0	 	1	2
	//2	3	.	0	1	=	2	3
	//3	4	 	 	 	 	3	4
	//
}

func TestSum(t *testing.T) {
	m1 := RandMatrix(5, 2, 10, 50)
	m2 := RandMatrix(5, 2, 0, 30)
	m3 := RandMatrix(5, 2, 0, 30)
	Println(m1, "+", m2, "+", m3, "=", Sum(m1, m2, m3))
}

func TestSumf(t *testing.T) {
	m1 := RandMatrixf(5, 2, 10, 50)
	m2 := RandMatrixf(5, 2, 0, 30)
	m3 := RandMatrixf(5, 2, 0, 30)
	fmt.Println(m1, m2, m3, Sumf(m1, m2, m3))
	fmt.Println(m1.Sumf(m2, m3))
}

func TestDotC(t *testing.T) {
	m1 := RandMatrix(5, 2, 10, 50)
	Println(m1, "*", 2, "=", DotC(m1, 2))
}

func TestDot(t *testing.T) {
	m1 := RandMatrix(1, 3, 0, 10)
	m2 := RandMatrix(3, 3, 0, 10)
	Println(m1, "*", m2, "=", Dot(m1, m2))
}

func TestDot2(t *testing.T) {
	m1 := NewMatrix([]int{1, 2, 3}, []int{4, 5, 6})
	m2 := NewMatrix([]int{1, 4}, []int{2, 5}, []int{3, 6})

	fmt.Println(m1, m2, Dot(m1, m2))
}

func TestDotEys(t *testing.T) {
	m1 := NewMatrix([]int{1, 2, 3}, []int{4, 5, 6})
	fmt.Println(Eye(3).Dot(m1))
}
