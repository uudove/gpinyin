package pconv

import (
	"bytes"
	"strconv"
	"strings"
)

const (
	// FormatWithToneMark Convert to pinyin with tone mark.
	// For example:
	//  - "杭州西湖" -> háng zhōu xī hú
	FormatWithToneMarkInternal = iota

	// FormatWithoutTone Convert to pinyin without tone.
	// For example:
	//  - "杭州西湖" -> hang zhou xi hu
	FormatWithoutToneInternal

	// FormatWithToneNumber Convert to pinyin without tone.
	// For example:
	//  - "杭州西湖" -> hang2 zhou1 xi1 hu2
	FormatWithToneNumberInternal
)

const (
	pinyinSeprator   = ","
	chineseLing      = "〇"
	allUnmarkedVowel = "aeiouv"
	allMarkedVowel   = "āáǎàēéěèīíǐìōóǒòūúǔùǖǘǚǜ"
)

var (
	chineseLingChar       = []rune(chineseLing)[0]
	allUnmarkedVowelChars = []rune(allUnmarkedVowel)
	allMarkedVowelChars   = []rune(allMarkedVowel)
)

func ConvertToPinyinString(text string, separator string, format int) (string, error) {
	text, err := ConvertToSimplifiedChinese(text)
	if err != nil {
		return "", err
	}

	err = loadPinyinDict()
	if err != nil {
		return "", err
	}

	return convertToPinyinString(text, separator, format), nil
}

func ConvertToPinyinArray(text string, format int) ([]string, error) {
	text, err := ConvertToSimplifiedChinese(text)
	if err != nil {
		return nil, err
	}

	err = loadPinyinDict()
	if err != nil {
		return nil, err
	}

	return convertToPinyinArray(text, format), nil
}

func convertToPinyinString(text string, separator string, format int) string {
	buffer := bytes.Buffer{}
	textChars := []rune(text)

	for i := 0; i < len(textChars); {
		substr := string(textChars[i:])
		commonPrefixList := arrayTrie.commonPrefixSearch(substr)
		if len(commonPrefixList) == 0 {
			c := textChars[i]
			if isChinese(c) {
				pinyinArray := convertCharToPinyinArray(c, format)
				if pinyinArray != nil && len(pinyinArray) > 0 {
					buffer.WriteString(pinyinArray[0])
				} else {
					buffer.WriteRune(c)
				}
			} else {
				buffer.WriteRune(c)
			}
			i++
		} else {
			words := multiPinyinKeys[commonPrefixList[len(commonPrefixList)-1]]
			pinyinArray := formatPinyin(multiPinyinMap[words], format)
			for j, l := 0, len(pinyinArray); j < l; j++ {
				buffer.WriteString(pinyinArray[j])
				if j < l-1 {
					buffer.WriteString(separator)
				}
			}
			i += len([]rune(words))
		}

		if i < len(textChars) {
			buffer.WriteString(separator)
		}

	}
	return buffer.String()
}

func convertToPinyinArray(text string, format int) []string {
	pinyin := convertToPinyinString(text, pinyinSeprator, format)
	return strings.Split(pinyin, pinyinSeprator)
}

func convertCharToPinyinArray(c rune, format int) []string {
	pinyin := pinyinMap[string(c)]
	if pinyin != "" {
		return formatPinyin(pinyin, format)
	}
	return []string{}
}

func formatPinyin(pinyinString string, format int) []string {
	if format == FormatWithToneMarkInternal { // gpinyin.FormatWithToneMark
		return strings.Split(pinyinString, pinyinSeprator)

	} else if format == FormatWithoutToneInternal { // gpinyin.FormatWithoutTone
		return convertWithoutTone(pinyinString)

	} else if format == FormatWithToneNumberInternal { // gpinyin.FormatWithToneNumber
		return convertWithToneNumber(pinyinString)
	}

	return []string{}
}

func convertWithoutTone(pinyinString string) []string {
	for i := len(allMarkedVowelChars) - 1; i >= 0; i-- {
		originalChar := allMarkedVowelChars[i]
		replaceChar := allUnmarkedVowelChars[(i-i%4)/4]
		pinyinString = strings.Replace(pinyinString, string(originalChar), string(replaceChar), -1)
	}
	// replace ü to v
	pinyinString = strings.Replace(pinyinString, "ü", "v", -1)
	return strings.Split(pinyinString, pinyinSeprator)
}

func convertWithToneNumber(pinyinString string) []string {
	pinyinArray := strings.Split(pinyinString, pinyinSeprator)
	for i := len(pinyinArray) - 1; i >= 0; i-- {
		hasMarkedChar := false
		originalPinyin := strings.Replace(pinyinArray[i], "ü", "v", -1) // ü -> v
		originalPinyinChars := []rune(originalPinyin)

		for j := len(originalPinyinChars) - 1; j >= 0; j-- {
			originalChar := originalPinyinChars[j]

			// 搜索带声调的拼音字母，如果存在则替换为对应不带声调的英文字母
			if originalChar < 'a' || originalChar > 'z' {
				indexInAllMarked := indexOfAllMarked(originalChar)
				// indexInAllMarked := strings.IndexRune(allMarkedVowel, originalChar)
				toneNumber := indexInAllMarked%4 + 1 // 声调数
				replaceChar := allUnmarkedVowelChars[(indexInAllMarked-indexInAllMarked%4)/4]
				pinyinArray[i] = strings.Replace(originalPinyin, string(originalChar), string(replaceChar), -1) + strconv.Itoa(toneNumber)
				hasMarkedChar = true
				break
			}
		}
		if !hasMarkedChar {
			// 找不到带声调的拼音字母说明是轻声，用数字5表示
			pinyinArray[i] = originalPinyin + "5"
		}
	}

	return pinyinArray
}

func indexOfAllMarked(item rune) int {
	for i, v := range allMarkedVowelChars {
		if v == item {
			return i
		}
	}
	return -1
}
