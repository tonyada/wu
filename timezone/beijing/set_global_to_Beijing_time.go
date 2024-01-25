package beijing

import (
	"os"
)

func init() {
	os.Setenv("TZ", "Asia/Shanghai")
}
