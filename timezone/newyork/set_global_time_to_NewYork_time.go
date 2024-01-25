package beijing

import (
	"os"
)

func init() {
	os.Setenv("TZ", "America/New_York")
}
