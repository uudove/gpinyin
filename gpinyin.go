package gpinyin

import (
	"github.com/uudove/gpinyin/pconv"
)

const (
	// Convert to pinyin with tone mark.
	// For example:
	//  - "杭州西湖" -> háng zhōu xī hú
	FORMAT_WITH_TONE_MARK = iota

	// Convert to pinyin without tone.
	// For example:
	//  - "杭州西湖" -> hang zhou xi hu
	FORMAT_WITHOUT_TONE

	// Convert to pinyin without tone.
	// For example:
	//  - "杭州西湖" -> hang2 zhou1 xi1 hu2
	FORMAT_WITH_TONE_NUMBER
)

// ConvertToPinyinString convert string to pinyin with specific formats.
//  - text the input text to be converted
//  - separator the separator between pinyins, can use "" if don't want to put a separator
//  - format any of FORMAT_WITH_TONE_MARK, FORMAT_WITHOUT_TONE or FORMAT_WITH_TONE_NUMBER
func ConvertToPinyinString(text string, separator string, format int) (string, error) {
	return "", nil

}

// ConvertToPinyinString convert string to pinyin array with specific formats.
//  - text the input text to be converted
//  - format any of FORMAT_WITH_TONE_MARK, FORMAT_WITHOUT_TONE or FORMAT_WITH_TONE_NUMBER
func ConvertToPinyinArray(text string, format int) ([]string, error) {
	return nil, nil
}

func ConvertToSimplifiedChinese(text string) (string, error) {
	return pconv.ConvertToSimplifiedChinese(text)
}

func ConvertToTraditionalChinese(text string) (string, error) {
	return pconv.ConvertToTraditionalChinese(text)
}
