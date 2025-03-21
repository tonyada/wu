package num

import (
	"math/rand"
	"strings"
	"time"
)

// random
func Random(max int) int { return rand.Intn(max) }
func RandomByTime(max int) int {
	source := rand.NewSource(time.Now().UnixNano())
	localRand := rand.New(source)
	return localRand.Intn(max)
}

func RandomRange(min, max int) int { return rand.Intn(max-min) + min }
func RandomRangeByTime(min, max int) int {
	source := rand.NewSource(time.Now().UnixNano())
	localRand := rand.New(source)
	return localRand.Intn(max-min) + min
}

func RandomExpFloat64() float64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.ExpFloat64()
}

func RandomFloat32() float32 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Float32()
}

func RandomFloat64() float64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Float64()
}

func RandomInt() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int()
}

func RandomInt31() int32 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int31()
}

func RandomInt31n(n int32) int32 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int31n(n)
}

func RandomInt63() int64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int63()
}

func RandomInt63n(n int64) int64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int63n(n)
}

func RandomIntn(n int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(n)
}

func RandomNormFloat64() float64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.NormFloat64()
}

func RandomPerm(n int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Perm(n)
}

func RandomUint32() uint32 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Uint32()
}

// 一个自定义的秘钥或随机字符串生成器

const (
	StrSource = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	HexSource = "0123456789ABCDEF"
	NumSource = "0123456789"
)

// 用内置字符串生成随机字符串或秘钥
func RandomStr(length int) string {
	sb := []string{}
	if length > 0 {
		for range length {
			sb = append(sb, string(StrSource[rand.Intn(len(StrSource))]))
		}
	}
	return strings.Join(sb, "")
}

// 生成十六进制随机串
func RandomHex(length int) string {
	sb := []string{}
	if length > 0 {
		for range length {
			sb = append(sb, string(HexSource[rand.Intn(len(HexSource))]))
		}
	}
	return strings.Join(sb, "")
}

// 生成纯数字的随机串
func RandomNum(length int) string {
	sb := []string{}
	if length > 0 {
		for range length {
			sb = append(sb, string(NumSource[rand.Intn(len(NumSource))]))
		}
	}
	return strings.Join(sb, "")
}
