package str

import (
	"bytes"
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
	. "wu"
)

// Find Find获取一个切片并在其中查找元素。如果找到它，它将返回它的密钥，否则它将返回-1和一个错误的bool。
func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

// IntersectArray 求两个切片的交集
func IntersectArray(a []string, b []string) []string {
	var inter []string
	mp := make(map[string]bool)

	for _, s := range a {
		if _, ok := mp[s]; !ok {
			mp[s] = true
		}
	}
	for _, s := range b {
		if _, ok := mp[s]; ok {
			inter = append(inter, s)
		}
	}

	return inter
}

// DiffArray 求两个切片的差集
func DiffArray(a []int, b []int) []int {
	var diffArray []int
	temp := map[int]struct{}{}

	for _, val := range b {
		if _, ok := temp[val]; !ok {
			temp[val] = struct{}{}
		}
	}

	for _, val := range a {
		if _, ok := temp[val]; !ok {
			diffArray = append(diffArray, val)
		}
	}

	return diffArray
}

// RemoveArrayRepeatedElement 切片去重实现
func RemoveArrayRepeatedElement(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat && arr[i] != "" {
			newArr = append(newArr, arr[i])
		}
	}
	return
}

// remove empty string from array
func RemoveArrayEmptyElement(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		if arr[i] != "" {
			newArr = append(newArr, arr[i])
		}
	}
	return
}

// 切片去重实现
func arrayUnique(arr []string) []string {
	result := make([]string, 0, len(arr))
	temp := map[string]struct{}{}
	for i := 0; i < len(arr); i++ {
		if _, ok := temp[arr[i]]; ok != true {
			temp[arr[i]] = struct{}{}
			result = append(result, arr[i])
		}
	}
	return result
}

func ReverseArray(s interface{}) {
	ReverseSlice(s)
}

// panic if s is not a slice
func ReverseSlice(s interface{}) {
	size := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, size-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}

func HasPrefix(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

func HasSuffix(s, suffix string) bool {
	return strings.HasSuffix(s, suffix)
}

// Split("a,b,c", ",") -> ["a" "b" "c"]
func Split(s, sep string) []string {
	return strings.Split(s, sep)
}

// SplitAfter("a,b,c", ",") -> ["a," "b," "c"]
func SplitAfter(s, sep string) []string {
	return strings.SplitAfter(s, sep)
}

// Fields("  foo bar  baz   ") -> ["foo" "bar" "baz"]
func Fields(s string) []string {
	return strings.Fields(s)
}

func TrimString(s, cutset string) string { return strings.Trim(s, cutset) }

func TrimLeft(s, cutset string) string { return strings.TrimLeft(s, cutset) }
func DelLeft(s, cutset string) string  { return TrimLeft(s, cutset) }

func TrimRight(s, cutset string) string { return strings.TrimRight(s, cutset) }
func DelRight(s, cutset string) string  { return TrimRight(s, cutset) }

func DelSpace(s string) string  { return TrimSpace(s) }
func TrimSpace(s string) string { return strings.TrimSpace(s) }

// Contains("seafood", "foo") -> true
func Contains(s, substr string) bool {
	return strings.Contains(s, substr)
}

// ContainsAny("failure", "u & i") -> true
func ContainsAny(s, chars string) bool {
	return strings.ContainsAny(s, chars)
}

// del specific content from string
func DelFromString(s, start, end string) (string, error) {
	startIndex := strings.Index(s, start)
	if startIndex == -1 {
		return "", Errf("cannot found start index: %v", start)
	}
	// cut start pos from string and get left content
	startLeft := string([]byte(s)[:startIndex])

	endIndex := strings.Index(s, end)
	// add end string length
	endIndex += len(end)

	if endIndex == -1 {
		return "", Errf("cannot found end index: %v", end)
	}
	// endIndex number is greater than original string size
	if endIndex > len(s) {
		return "", Errf("endIndex %v > original string size: %v", endIndex, len(s))
	}
	// cut end pos from string and get right content
	endRight := string([]byte(s)[endIndex:])
	s = fmt.Sprintf("%v%v", startLeft, endRight)
	return TrimSpace(s), nil
}

// del specific content from string last
func DelFromStringLast(s, start, end string) (string, error) {
	startIndex := strings.LastIndex(s, start)
	if startIndex == -1 {
		return "", Errf("cannot found start index: %v", start)
	}
	// cut start pos from string and get left content
	startLeft := string([]byte(s)[:startIndex])

	endIndex := strings.LastIndex(s, end)
	// add end string length
	endIndex += len(end)

	if endIndex == -1 {
		return "", Errf("cannot found end index: %v", end)
	}
	// endIndex number is greater than original string size
	if endIndex > len(s) {
		return "", Errf("endIndex %v > original string size: %v", endIndex, len(s))
	}
	// cut end pos from string and get right content
	endRight := string([]byte(s)[endIndex:])
	s = fmt.Sprintf("%v%v", startLeft, endRight)
	return TrimSpace(s), nil
}

// cut specific content from string
func CutFromString(s, start, end string) (string, error) {
	startIndex := strings.Index(s, start)
	if startIndex == -1 {
		return "", Errf("cannot found start index: %v", start)
	}
	// add start string length
	startIndex += len(start)
	s = string([]byte(s)[startIndex:])
	endIndex := strings.Index(s, end)
	if endIndex == -1 {
		return "", Errf("cannot found end index: %v", end)
	}
	s = string([]byte(s)[:endIndex])
	return TrimSpace(s), nil
}
func CutFromStringLast(s, start, end string) (string, error) {
	startIndex := strings.LastIndex(s, start)
	if startIndex == -1 {
		return "", Errf("cannot found start index: %v", start)
	}
	// add start string length
	startIndex += len(start)
	s = string([]byte(s)[startIndex:])
	endIndex := strings.LastIndex(s, end)
	if endIndex == -1 {
		return "", Errf("cannot found end index: %v", end)
	}
	s = string([]byte(s)[:endIndex])
	return TrimSpace(s), nil
}

// del all [delStr] from s
func DelStr(s, delStr string) string {
	return strings.Replace(s, delStr, "", -1)
}

// del [delStr] from s with times If n < 0, there is no limit on the number of replacements.
func DelStrTimes(s, delStr string, times int) string {
	return strings.Replace(s, delStr, "", times)
}

// replace all [oldStr] to [newStr] from s
func ReplaceStr(s, oldStr, newStr string) string {
	return strings.Replace(s, oldStr, newStr, -1)
}

// replace [oldStr] to [newStr] with times If n < 0, there is no limit on the number of replacements.
func ReplaceStrTimes(s, oldStr, newStr string, times int) string {
	return strings.Replace(s, oldStr, newStr, times)
}

func Lowercase(s string) string {
	return strings.ToLower(s)
}

func Uppercase(s string) string {
	return strings.ToUpper(s)
}

// sub chinese
// func Substrcn(s string, l int) string {
// 	if len(s) <= l {
// 		return s
// 	}
// 	ss, sl, rl, rs := "", 0, 0, []rune(s)
// 	for _, r := range rs {
// 		rint := int(r)
// 		if rint < 128 {
// 			rl = 1
// 		} else {
// 			rl = 2
// 		}

// 		if sl+rl > l {
// 			break
// 		}
// 		sl += rl
// 		ss += string(r)
// 	}
// 	return ss
// }

// IsLetter returns true if the 'l' is an English letter
// func IsLetter(l uint8) bool {
// 	n := (l | 0x20) - 'a'
// 	if n >= 0 && n < 26 {
// 		return true
// 	}
// 	return false
// }

// if chinese?
func IsChineseChar(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) {
			return true
		}
	}
	return false
}

// len super 计算多语言和英语混合计数，为了itc的keyword
// func Lencn(s string) int {
// 	length := 0
// 	for _, r := range s {
// 		for k, _ := range unicode.Scripts {
// 			if unicode.Is(unicode.Scripts[k], r) {
// 				// Printf("language: %v s: %v len: %v\n", k, r, len(string(r)))
// 				length += 1 // unicode language len is 3 then make it to 1
// 				break
// 			}
// 		}
// 		// Han cn
// 		// Katakana 日语片假名 Hiragana 平假名
// 		// Hangul korea
// 		// if unicode.Is(unicode.Scripts["Han"], r) ||
// 		// 	unicode.Is(unicode.Scripts["Katakana"], r) ||
// 		// 	unicode.Is(unicode.Scripts["Hiragana"], r) ||
// 		// 	unicode.Is(unicode.Scripts["Hangul"], r) {
// 		// 	length += len(string(r)) / 3
// 		// } else {
// 		// 	length += len(string(r))
// 		// }
// 	}
// 	return length
// }

// deleteString removes the specified string from the slice.
// Returns the modified slice and nil if the string is found and deleted,
// otherwise returns the original slice and an error.
func DelFromSlice(slice []string, delStr string) ([]string, error) {
	for i, str := range slice {
		if str == delStr {
			return append(slice[:i], slice[i+1:]...), nil
		}
	}
	return slice, ErrNew(delStr + " not found")
}

// Append string appends string to slice with no duplicates.
func AppendToSliceWithNoDuplicates(slice []string, str string) ([]string, error) {
	for _, s := range slice {
		if s == str {
			return slice, ErrNew(str + " is duplicated")
		}
	}
	return append(slice, str), nil
}

// Append number appends string to slice with no duplicates.
func AppendToSliceWithNoDuplicatesInt64(slice []int64, i int64) ([]int64, error) {
	for _, s := range slice {
		if s == i {
			return slice, ErrNew(Sprintf("%v", i) + " is duplicated")
		}
	}
	return append(slice, i), nil
}

// CompareSliceStr compares two 'string' type slices.
// It returns true if elements and order are both the same.
func CompareSliceStr(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}

// CompareSliceStr compares two 'string' type slices.
// It returns true if elements are the same, and ignores the order
// func CompareSliceStrU(s1, s2 []string) bool {
// 	if len(s1) != len(s2) {
// 		return false
// 	}

// 	for i := range s1 {
// 		for j := len(s2) - 1; j >= 0; j-- {
// 			if s1[i] == s2[j] {
// 				s2 = append(s2[:j], s2[j+1:]...)
// 				break
// 			}
// 		}
// 	}
// 	if len(s2) > 0 {
// 		return false
// 	}
// 	return true
// }

// IsSliceContainsStr returns true if the string exists in given slice, ignore case.
func IsSliceContainsStr(sl []string, str string) bool {
	str = strings.ToLower(str)
	for _, s := range sl {
		if strings.ToLower(s) == str {
			return true
		}
	}
	return false
}

// IsSliceContainsInt64 returns true if the int64 exists in given slice.
func IsSliceContainsInt64(sl []int64, i int64) bool {
	for _, s := range sl {
		if s == i {
			return true
		}
	}
	return false
}

// Expand replaces {k} in template with match[k] or subs[atoi(k)] if k is not in match.
func Expand(template string, match map[string]string, subs ...string) string {
	var p []byte
	var i int
	for {
		i = strings.Index(template, "{")
		if i < 0 {
			break
		}
		p = append(p, template[:i]...)
		template = template[i+1:]
		i = strings.Index(template, "}")
		if s, ok := match[template[:i]]; ok {
			p = append(p, s...)
		} else {
			j, _ := strconv.Atoi(template[:i])
			if j >= len(subs) {
				p = append(p, []byte("Missing")...)
			} else {
				p = append(p, subs[j]...)
			}
		}
		template = template[i+1:]
	}
	p = append(p, template...)
	return string(p)
}

// Reverse s string, support unicode
func Reverse(s string) string {
	n := len(s)
	runes := make([]rune, n)
	for _, rune := range s {
		n--
		runes[n] = rune
	}
	return string(runes[n:])
}

// RandomCreateBytes generate random []byte by specify chars.
func RandomCreateBytes(n int, alphabets ...byte) []byte {
	const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, n)
	var randby bool
	if num, err := rand.Read(bytes); num != n || err != nil {
		rand.Seed(time.Now().UnixNano())
		randby = true
	}
	for i, b := range bytes {
		if len(alphabets) == 0 {
			if randby {
				bytes[i] = alphanum[rand.Intn(len(alphanum))]
			} else {
				bytes[i] = alphanum[b%byte(len(alphanum))]
			}
		} else {
			if randby {
				bytes[i] = alphabets[rand.Intn(len(alphabets))]
			} else {
				bytes[i] = alphabets[b%byte(len(alphabets))]
			}
		}
	}
	return bytes
}

// ToSnakeCase can convert all upper case characters in a string to
// underscore format.
//
// Some samples.
//
//	"FirstName"  => "first_name"
//	"HTTPServer" => "http_server"
//	"NoHTTPS"    => "no_https"
//	"GO_PATH"    => "go_path"
//	"GO PATH"    => "go_path"      // space is converted to underscore.
//	"GO-PATH"    => "go_path"      // hyphen is converted to underscore.
//
// From https://github.com/huandu/xstrings
func ToSnakeCase(str string) string {
	if len(str) == 0 {
		return ""
	}

	buf := &bytes.Buffer{}
	var prev, r0, r1 rune
	var size int

	r0 = '_'

	for len(str) > 0 {
		prev = r0
		r0, size = utf8.DecodeRuneInString(str)
		str = str[size:]

		switch {
		case r0 == utf8.RuneError:
			buf.WriteByte(byte(str[0]))

		case unicode.IsUpper(r0):
			if prev != '_' {
				buf.WriteRune('_')
			}

			buf.WriteRune(unicode.ToLower(r0))

			if len(str) == 0 {
				break
			}

			r0, size = utf8.DecodeRuneInString(str)
			str = str[size:]

			if !unicode.IsUpper(r0) {
				buf.WriteRune(r0)
				break
			}

			// find next non-upper-case character and insert `_` properly.
			// it's designed to convert `HTTPServer` to `http_server`.
			// if there are more than 2 adjacent upper case characters in a word,
			// treat them as an abbreviation plus a normal word.
			for len(str) > 0 {
				r1 = r0
				r0, size = utf8.DecodeRuneInString(str)
				str = str[size:]

				if r0 == utf8.RuneError {
					buf.WriteRune(unicode.ToLower(r1))
					buf.WriteByte(byte(str[0]))
					break
				}

				if !unicode.IsUpper(r0) {
					if r0 == '_' || r0 == ' ' || r0 == '-' {
						r0 = '_'

						buf.WriteRune(unicode.ToLower(r1))
					} else {
						buf.WriteRune('_')
						buf.WriteRune(unicode.ToLower(r1))
						buf.WriteRune(r0)
					}

					break
				}

				buf.WriteRune(unicode.ToLower(r1))
			}

			if len(str) == 0 || r0 == '_' {
				buf.WriteRune(unicode.ToLower(r0))
				break
			}

		default:
			if r0 == ' ' || r0 == '-' {
				r0 = '_'
			}

			buf.WriteRune(r0)
		}
	}

	return buf.String()
}
