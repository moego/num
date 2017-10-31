package number

import (
	"math/rand"
	"time"
)

func RandInt(min, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if min > max {
		return r.Intn(min-max) + min
	}
	return r.Intn(max-min) + min
}

func RandFloat(min, max float64) float64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if min > max {
		return  (min-max) * r.Float64() +  max
	}
	return  (max -min) *r.Float64() + min
}
