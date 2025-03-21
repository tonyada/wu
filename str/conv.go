package str

import (
	"fmt"
	"strconv"
	"wu"

	"github.com/pkg/errors"
)

func parseInt(s string) (n int64, ok bool) {
	var i int
	var sign bool
	if len(s) > 0 && s[0] == '-' {
		sign = true
		i++
	}
	if i == len(s) {
		return 0, false
	}
	for ; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			n = n*10 + int64(s[i]-'0')
		} else {
			return 0, false
		}
	}
	if sign {
		return n * -1, true
	}
	return n, true
}

func parseUint(s string) (n uint64, ok bool) {
	var i int
	if i == len(s) {
		return 0, false
	}
	for ; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			n = n*10 + uint64(s[i]-'0')
		} else {
			return 0, false
		}
	}
	return n, true
}

// string to int
func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		wu.LogErr(err)
	}
	return i
}

// string to int64
func StringToInt64(s string) int64 {
	i, ok := parseInt(s)
	if !ok {
		wu.LogDanger("not ok")
	}
	return i
}

// string to uint64
func StringToUint64(s string) uint64 {
	i, ok := parseUint(s)
	if !ok {
		wu.LogDanger("not ok")
	}
	return i
}

// string to uint64
func StrToUint64(s string) uint64 {
	return StringToUint64(s)
}

// sting to float
func StringToFloat64(s string) float64 {
	i, err := strconv.ParseFloat(s, 64)
	if err != nil {
		wu.LogErr(err)
	}
	return i
}

// float to decimal 2, i.e: 1.12
func FloatToDot2(f float64) float64 {
	return FloatDecimal(f, "%.2f")
}

// float to str & Keep decimal 2, i.e: 1.12
func FloatToDot2Str(f float64) string {
	return FloatToStr(f, 2)
}

// float to decimal 3, i.e: 1.123
func FloatToDot3(f float64) float64 {
	return FloatDecimal(f, "%.3f")
}

// float to str & Keep decimal 3, i.e: 1.123
func FloatToDot3Str(f float64) string {
	return FloatToStr(f, 3)
}

// float to decimal 4, i.e: 1.1234
func FloatToDot4(f float64) float64 {
	return FloatDecimal(f, "%.4f")
}

// float to str & Keep decimal 4, i.e: 1.1234
func FloatToDot4Str(f float64) string {
	return FloatToStr(f, 4)
}

// float to float format ie: "%.2f"
func FloatDecimal(f float64, format string) float64 {
	f, _ = strconv.ParseFloat(fmt.Sprintf(format, f), 64)
	return f
}

// float to format string n is decimal
func FloatToStr(f float64, n int) string {
	return strconv.FormatFloat(f, 'f', n, 64)
}

// string to float64
func StrToFloat(s string) float64 {
	return StringToFloat64(s)
}

// string to float
func StrToFloat64(s string) float64 {
	return StringToFloat64(s)
}

// int to string
func IntToString(i int) string {
	return strconv.Itoa(i)
}

// string to int
func StrToInt(s string) int {
	return StringToInt(s)
}

// string to int64
func StrToInt64(s string) int64 {
	return StringToInt64(s)
}

// int to string
func IntToStr(i int) string {
	return strconv.Itoa(i)
}

// Convert string to specify type.
type StrTo string

func (f StrTo) Exist() bool {
	return string(f) != string(rune(0x1E))
}

func (f StrTo) Uint8() (uint8, error) {
	v, err := strconv.ParseUint(f.String(), 10, 8)
	return uint8(v), err
}

func (f StrTo) Int() (int, error) {
	v, err := strconv.ParseInt(f.String(), 10, 0)
	return int(v), err
}

func (f StrTo) Int64() (int64, error) {
	v, err := strconv.ParseInt(f.String(), 10, 64)
	return int64(v), err
}

func (f StrTo) MustUint8() uint8 {
	v, _ := f.Uint8()
	return v
}

func (f StrTo) MustInt() int {
	v, _ := f.Int()
	return v
}

func (f StrTo) MustInt64() int64 {
	v, _ := f.Int64()
	return v
}

func (f StrTo) String() string {
	if f.Exist() {
		return string(f)
	}
	return ""
}

// Convert any type to string.
func ToStr(value any, args ...int) (s string) {
	switch v := value.(type) {
	case bool:
		s = strconv.FormatBool(v)
	case float32:
		s = strconv.FormatFloat(float64(v), 'f', argInt(args).Get(0, -1), argInt(args).Get(1, 32))
	case float64:
		s = strconv.FormatFloat(v, 'f', argInt(args).Get(0, -1), argInt(args).Get(1, 64))
	case int:
		s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
	case int8:
		s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
	case int16:
		s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
	case int32:
		s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
	case int64:
		s = strconv.FormatInt(v, argInt(args).Get(0, 10))
	case uint:
		s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
	case uint8:
		s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
	case uint16:
		s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
	case uint32:
		s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
	case uint64:
		s = strconv.FormatUint(v, argInt(args).Get(0, 10))
	case string:
		s = v
	case []byte:
		s = string(v)
	default:
		s = fmt.Sprintf("%v", v)
	}
	return s
}

type argInt []int

func (a argInt) Get(i int, args ...int) (r int) {
	if i >= 0 && i < len(a) {
		r = a[i]
	} else if len(args) > 0 {
		r = args[0]
	}
	return
}
func PowInt(x int, y int) int {
	if y <= 0 {
		return 1
	} else {
		if y%2 == 0 {
			sqrt := PowInt(x, y/2)
			return sqrt * sqrt
		} else {
			return PowInt(x, y-1) * x
		}
	}
}

// HexStr2int converts hex format string to decimal number.
func HexStr2int(hexStr string) (int, error) {
	num := 0
	length := len(hexStr)
	for i := 0; i < length; i++ {
		char := hexStr[length-i-1]
		var factor int
		switch {
		case char >= '0' && char <= '9':
			factor = int(char) - '0'
		case char >= 'a' && char <= 'f':
			factor = int(char) - 'a' + 10
		default:
			return -1, errors.Errorf("invalid hex: %s", string(char))
		}

		num += factor * PowInt(16, i)
	}
	return num, nil
}

// Int2HexStr converts decimal number to hex format string.
func Int2HexStr(num int) (hex string) {
	if num == 0 {
		return "0"
	}

	for num > 0 {
		r := num % 16
		var c string
		if r >= 0 && r <= 9 {
			c = string(rune(r + '0'))
		} else {
			c = string(rune(r + 'a' - 10))
		}
		hex = c + hex
		num = num / 16
	}
	return hex
}
