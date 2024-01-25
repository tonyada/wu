package md5

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
)

// gen 32bit MD5
func New(text string) (string, error) {
	ctx := md5.New()
	_, err := ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil)), err
}

// gen 32bit MD5 with ? repeat times
func NewWithRepeatTimes(text string, repeatTimes int) (string, error) {
	var err error
	myhash := md5.New()
	_, err = io.WriteString(myhash, text)
	hashedStr := fmt.Sprintf("%x", myhash.Sum([]byte{}))

	if repeatTimes == 0 {
		repeatTimes = 2 // min is 2
	}
	for i := 1; i < repeatTimes; i++ {
		text = hashedStr
		newHash := md5.New()
		_, err = io.WriteString(newHash, text)
		hashedStr = fmt.Sprintf("%x", newHash.Sum([]byte{}))
	}
	return hashedStr, err
}
