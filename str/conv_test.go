package str

import (
	"fmt"
	"testing"

	a "github.com/smartystreets/goconvey/convey"
)

var f1 float64 = 3.1415926
var f2 float64 = 1.1

func TestDefine(t *testing.T) {
	a.Convey("2变量定义", t, func() {
		got := fmt.Sprintf("f1 = %f f2 = %f", f1, f2)
		t.Logf("f1 = %f f2 = %f", f1, f2)
		a.So(got, a.ShouldResemble, "f1 = 3.141593 f2 = 1.100000")
	})
}

func TestFloatToDot2(t *testing.T) {
	a.Convey("TestFloatToDot2", t, func() {
		got := fmt.Sprintf("f1 = %v f2 = %v", FloatToDot2(f1), FloatToDot2(f2))
		t.Log(FloatToDot2(f1), FloatToDot2(f2))
		a.So(got, a.ShouldResemble, "f1 = 3.14 f2 = 1.1")
	})

}

func TestDot3(t *testing.T) {
	t.Log(FloatToDot3(f1), FloatToDot3(f2))
}
func TestFloatToStr(t *testing.T) {
	t.Log(FloatToStr(f2, 12))
}
func TestDecimal(t *testing.T) {
	t.Log(FloatToDot3(f1), FloatDecimal(f2, "%3f"))
}
func TestFloatToDot2Str(t *testing.T) {
	t.Log(FloatToDot2Str(f2))
}

func TestFloatToDot3Str(t *testing.T) {
	t.Log(FloatToDot3Str(f2))
}
