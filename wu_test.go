package wu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// - Test log
func TestLog(t *testing.T) {
	LogDebugOn()
	Log("log1")
	Log("log2")
	LogNote("log note")
	LogErr("log error")
	LogInfo("log Info")
	LogInfoPrefix("log Info")
	LogWarn("warning")
	LogWarnPrefix("warning")
	LogDanger("danger")
	LogDangerPrefix("danger")
	LogSuccess("success")
	LogSuccessPrefix("success")
	Log("log")
	LogWithPrefix("my prefix 测试", "log with prefix")
}

func TestPrint(t *testing.T) {
	Println("hello")
	str := "string"
	num := 12
	Printfln("hello %v %v", str, num)
}

func TestError(t *testing.T) {
	ErrorOn()
	msg := ""
	// set to no err
	var err error = nil
	if Err(err) {
		msg = "is err"
	} else {
		msg = "no err"
	}
	// assert equality
	assert.Equal(t, msg, "no err", "they should be equal")
	println(msg)
	// trigger an error
	err2 := ErrNew("I am new err")
	if Err(err2) {
		msg = "is err"
	} else {
		msg = "no err"
	}
	assert.Equal(t, msg, "is err", "they should be equal")
	println(msg)

	if NotErr(err) {
		msg = "not err"
	} else {
		msg = "err"
	}
	assert.Equal(t, msg, "not err", "they should be equal")
	println(msg)

	if OK(err) {
		msg = "ok"
	} else {
		msg = "err"
	}
	assert.Equal(t, msg, "ok", "they should be equal")
	println(msg)
}
