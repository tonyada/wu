package wu

import (
	"regexp"

	"github.com/mozillazg/go-pinyin"
)

// - Replace string's special chars to [replaceToStr]
// ReplaceSpecialCharsTo("user@domain.com", "_") = user_domain_com
func ReplaceSpecialCharsTo(str string, replaceToStr string) string {
	// define regular expression pattern, match non letter, number,
	// underscore, dash symbol and dot symbol of the characters
	re := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	// replace all special chars to [replaceTo]
	return re.ReplaceAllString(str, replaceToStr)
}

// - Chinese to pinyin
// 贵州茅台 to gui zhou mao tai
func ChineseToPinyin(hans string) []string {
	return pinyin.LazyPinyin(hans, pinyin.NewArgs())
}

// - Chinese to pinyin initails
// 贵州茅台 to gzmt
func ChineseToPinyinInitials(hans string) string {
	// 将中文转换为拼音
	cn_pinyins := ChineseToPinyin(hans)
	// 提取拼音首字母
	var pinyin_initials string
	for _, v := range cn_pinyins {
		if len(v) > 0 {
			pinyin_initials += string(v[0])
		}
	}
	return pinyin_initials
}
