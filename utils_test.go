package wu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplaceSpecialCharsTo(t *testing.T) {
	// replace email special chars with underscore
	email := "user@domain.com"
	replacedEmail := ReplaceSpecialCharsTo(email, "_")
	assert.Equal(t, "user_domain_com", replacedEmail, "they should be equal")
}
