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
func TestChineseToPinyin(t *testing.T) {
	// replace email special chars with underscore
	py := ChineseToPinyin("贵州茅台")
	assert.Equal(t, []string{"gui", "zhou", "mao", "tai"}, py, "they should be equal")
}

func TestChineseToPinyinInitials(t *testing.T) {
	// replace email special chars with underscore
	py := ChineseToPinyinInitials("贵州茅台")
	assert.Equal(t, "gzmt", py, "they should be equal")
}
func TestChineseToPinyinInitialsWithEnglishLetters(t *testing.T) {
	// replace email special chars with underscore
	py := ChineseToPinyinInitials("TCL科技")
	assert.Equal(t, "tclkj", py, "they should be equal")
}
