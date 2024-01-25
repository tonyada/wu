package beijing

import (
	"os"
)

func init() {
	os.Setenv("TZ", "America/Los_Angeles")
}
