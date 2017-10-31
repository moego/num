package number

import (
	"testing"
	"fmt"
)

func TestSum(t *testing.T) {
	m1:=RandMatrix(5,2,10,50)
	m2:=RandMatrix(5,2,0,30)
	fmt.Println(m1,m2,Sum(m1,m2))
}
