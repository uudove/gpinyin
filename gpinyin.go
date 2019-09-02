package gpinyin

import (
	"errors"

	"./pconv"
)

const (
	// FormatWithToneMark Convert to pinyin with tone mark.
	// For example:
	//  - "杭州西湖" -> háng zhōu xī hú
	FormatWithToneMark = pconv.FormatWithToneMarkInternal

	// FormatWithoutTone Convert to pinyin without tone.
	// For example:
	//  - "杭州西湖" -> hang zhou xi hu
	FormatWithoutTone = pconv.FormatWithoutToneInternal

	// FormatWithToneNumber Convert to pinyin without tone.
	// For example:
	//  - "杭州西湖" -> hang2 zhou1 xi1 hu2
	FormatWithToneNumber = pconv.FormatWithToneNumberInternal
)

// ConvertToPinyinString - Convert a string to pinyin with specific formats.
//  - text: the input text to be converted
//  - separator: the separator between pinyins, can use "" if don't want to put a separator
//  - format: one of FormatWithToneMark, FormatWithoutTone or FormatWithToneNumber
func ConvertToPinyinString(text string, separator string, format int) (string, error) {
	err := validFormat(format)
	if err != nil {
		return "", err
	}
	return pconv.ConvertToPinyinString(text, separator, format)
}

// ConvertToPinyinArray - Convert a string to pinyin array with specific formats.
//  - text: the input text to be converted
//  - format: one of FormatWithToneMark, FormatWithoutTone or FormatWithToneNumber
func ConvertToPinyinArray(text string, format int) ([]string, error) {
	err := validFormat(format)
	if err != nil {
		return nil, err
	}
	return pconv.ConvertToPinyinArray(text, format)
}

// ConvertToSimplifiedChinese - Convert a Traditional Chinese string to Simplified Chinese string
//  - text: the input text to be converted
func ConvertToSimplifiedChinese(text string) (string, error) {
	return pconv.ConvertToSimplifiedChinese(text)
}

// ConvertToTraditionalChinese - Convert a Simplified Chinese string to Traditional Chinese string
//  - text: the input text to be converted
func ConvertToTraditionalChinese(text string) (string, error) {
	return pconv.ConvertToTraditionalChinese(text)
}

// LoadDict can load dict before use any convertor.
// If loaded, the first time call any convertor, will be faster
func LoadDict() error {
	return pconv.LoadAllDict()
}

// validFormat - verify PinyinFormat
func validFormat(format int) error {
	if format != FormatWithToneMark && format != FormatWithToneNumber && format != FormatWithoutTone {
		return errors.New("Invalid format, must be one of FormatWithToneMark, FormatWithoutTone or FormatWithToneNumber")
	}
	return nil
}
