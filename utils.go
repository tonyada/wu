package wu

import "regexp"

// - Replace string's special chars to [replaceToStr]
// ReplaceSpecialCharsTo("user@domain.com", "_") = user_domain_com
func ReplaceSpecialCharsTo(str string, replaceToStr string) string {
	// define regular expression pattern, match non letter, number,
	// underscore, dash symbol and dot symbol of the characters
	re := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	// replace all special chars to [replaceTo]
	return re.ReplaceAllString(str, replaceToStr)
}
