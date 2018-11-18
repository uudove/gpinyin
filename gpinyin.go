package gpinyin

import (
	"github.com/uudove/gpinyin/pconv"
)

const (
	// FormatWithToneMark Convert to pinyin with tone mark.
	// For example:
	//  - "杭州西湖" -> háng zhōu xī hú
	FormatWithToneMark = 0

	// FormatWithoutTone Convert to pinyin without tone.
	// For example:
	//  - "杭州西湖" -> hang zhou xi hu
	FormatWithoutTone = 1

	// FormatWithToneNumber Convert to pinyin without tone.
	// For example:
	//  - "杭州西湖" -> hang2 zhou1 xi1 hu2
	FormatWithToneNumber = 2
)

// ConvertToPinyinString convert string to pinyin with specific formats.
//  - text: the input text to be converted
//  - separator: the separator between pinyins, can use "" if don't want to put a separator
//  - format: any of FormatWithToneMark, FormatWithoutTone or FormatWithToneNumber
func ConvertToPinyinString(text string, separator string, format int) (string, error) {
	return pconv.ConvertToPinyinString(text, separator, format)

}

// ConvertToPinyinArray convert string to pinyin array with specific formats.
//  - text: the input text to be converted
//  - format: any of FormatWithToneMark, FormatWithoutTone or FormatWithToneNumber
func ConvertToPinyinArray(text string, format int) ([]string, error) {
	return pconv.ConvertToPinyinArray(text, format)
}

// ConvertToSimplifiedChinese convert Traditional Chinese to Simplified Chinese
//  - text: the input text to be converted
func ConvertToSimplifiedChinese(text string) (string, error) {
	return pconv.ConvertToSimplifiedChinese(text)
}

// ConvertToTraditionalChinese convert Simplified Chinese to Traditional Chinese
//  - text: the input text to be converted
func ConvertToTraditionalChinese(text string) (string, error) {
	return pconv.ConvertToTraditionalChinese(text)
}

// LoadDict can load dict before use any convertor.
// If called, the first time call any convertor, will be faster
func LoadDict() error {
	return pconv.LoadAllDict()
}
