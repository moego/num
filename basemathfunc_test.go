package num

import (
	"fmt"
	"testing"
)

func TestSgn(t *testing.T) {
	for i := -10.0; i < 10; i += 1 {
		fmt.Print(Sgn(i))
	}
}
func TestSignoid(t *testing.T) {
	for i := -10.0; i < 10; i += 1 {
		fmt.Printf("%f \t %3f\t\n", i, Signoid(i))
	}
}
