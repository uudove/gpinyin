package pconv

import (
	"regexp"
)

const (
	chineseRegex = "[\\u4e00-\\u9fa5]"
)

func ConvertToSimplifiedChinese(text string) (string, error) {
	err := loadChineseDict()
	if err != nil {
		return "", err
	}

	s := ""
	chars := []rune(text)
	for _, v := range chars {
		s += convertToSimplifiedChinese(v)
	}

	return s, nil
}

func ConvertToTraditionalChinese(text string) (string, error) {
	err := loadChineseDict()
	if err != nil {
		return "", err
	}

	s := ""
	chars := []rune(text)
	for _, v := range chars {
		s += convertToTraditionalChinese(v)
	}

	return s, nil
}

func convertToSimplifiedChinese(char rune) string {
	s := string(char)
	value := chineseMap[s]
	if value == "" {
		return s
	}
	return value
}

func convertToTraditionalChinese(char rune) string {
	s := string(char)
	for k, v := range chineseMap {
		if v == s {
			return k
		}
	}
	return s
}

func isChinese(c rune) bool {
	reg := regexp.MustCompile(chineseRegex)
	return c == chineseLingChar || reg.MatchString(string(c))
}
