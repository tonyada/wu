package wu

import (
	"regexp"
	"strings"

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

// - Extract English letters from string
func extractEnglishLetters(input string) string {
	// 定义正则表达式匹配规则
	regex := regexp.MustCompile(`[a-zA-Z]+`)

	// 使用正则表达式查找匹配项
	matches := regex.FindAllString(input, -1)

	// 将所有匹配项拼接成一个字符串
	result := ""
	for _, match := range matches {
		result += match
	}

	return strings.ToLower(result)
}

// - Chinese to pinyin initails
// 贵州茅台 to gzmt
func ChineseToPinyinInitials(hans string) string {
	// 创建拼音转换器
	py := pinyin.NewArgs()

	// 将中文字符串转换为拼音
	pinyinResult := pinyin.Pinyin(hans, py)

	// 提取每个拼音的首字母并连接它们
	var initials strings.Builder
	for _, p := range pinyinResult {
		initials.WriteString(strings.ToLower(string(p[0][0])))
	}
	return extractEnglishLetters(hans) + initials.String()
}
