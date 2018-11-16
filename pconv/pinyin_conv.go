package pconv

func ConvertToPinyinString(text string, separator string, format int) (string, error) {
	err := loadPinyinDict()
	if err != nil {
		return "", err
	}

	return convertToPinyinString(text, separator, format), nil
}

func ConvertToPinyinArray(text string, format int) ([]string, error) {
	err := loadPinyinDict()
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func convertToPinyinString(text string, separator string, format int) string {

	return ""
}
