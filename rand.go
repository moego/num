package num

import (
	"math/rand"
	"time"
)

//RandInt 生成随机整数
func RandInt(min, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if min > max {
		return r.Intn(min-max) + min
	}
	return r.Intn(max-min) + min
}

//RandFloat 生成随机浮点数
func RandFloat(min, max float64) float64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if min > max {
		return (min-max)*r.Float64() + max
	}
	return (max-min)*r.Float64() + min
}
