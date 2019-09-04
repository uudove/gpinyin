package gpinyin

import (
	"testing"
	"time"
)

func TestLoadDict(t *testing.T) {
	startTime := time.Now()
	err := LoadDict()
	if err != nil {
		t.Fatal("TestLoadDict with error: ", err.Error())
	}

	t.Log("Load dict spend(ms):", (time.Now().Nanosecond()-startTime.Nanosecond())/1000000)
}

func TestConvertToSimplifiedChinese(t *testing.T) {
	s := "今天天氣好晴朗，處處好風光"
	simplied, err := ConvertToSimplifiedChinese(s)
	if err != nil {
		t.Fatal("TestConvertToSimplifiedChinese with error: ", err.Error())
	}
	if simplied != "今天天气好晴朗，处处好风光" {
		t.Fatal("TestConvertToSimplifiedChinese not match: ", simplied)
	}
}

func TestConvertToTraditionalChinese(t *testing.T) {
	s := "今天天气好晴朗，处处好风光"
	traditional, err := ConvertToTraditionalChinese(s)
	if err != nil {
		t.Fatal("TestConvertToTraditionalChinese with error: ", err.Error())
	}
	if traditional != "今天天氣好晴朗，處處好風光" {
		t.Fatal("TestConvertToTraditionalChinese not match: ", traditional)
	}

}

func TestConvertToPinyinStringFormatWithToneMark(t *testing.T) {
	s := "杭州西湖"
	// test FormatWithToneMark
	pinyin, err := ConvertToPinyinString(s, " ", FormatWithToneMark)
	if err != nil {
		t.Fatal("ConvertToPinyinString FormatWithToneMark with error: ", err.Error())
	}
	if pinyin != "háng zhōu xī hú" {
		t.Fatal("ConvertToPinyinString WithToneMark not match: ", pinyin)
	}
}

func TestConvertToPinyinStringFormatWithoutTone(t *testing.T) {
	s := "杭州西湖"
	// test FormatWithoutTone
	pinyin, err := ConvertToPinyinString(s, ",", FormatWithoutTone)
	if err != nil {
		t.Fatal("ConvertToPinyinString FormatWithoutTone with error: ", err.Error())
	}
	if pinyin != "hang,zhou,xi,hu" {
		t.Fatal("ConvertToPinyinString FormatWithoutTone not match: ", pinyin)
	}
}

func TestConvertToPinyinStringFormatWithToneNumber(t *testing.T) {
	s := "杭州西湖"
	// test FormatWithToneNumber
	pinyin, err := ConvertToPinyinString(s, "", FormatWithToneNumber)
	if err != nil {
		t.Fatal("ConvertToPinyinString FormatWithToneNumber with error: ", err.Error())
	}
	if pinyin != "hang2zhou1xi1hu2" {
		t.Fatal("ConvertToPinyinString FormatWithToneNumber not match: ", pinyin)
	}
}
