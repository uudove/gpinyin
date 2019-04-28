package gpinyin

import "testing"
import "time"

func TestLoadDict(t *testing.T) {
	startTime := time.Now()
	err := LoadDict()
	if err != nil {
		t.Fatal("TestLoadDict with error: ", err.Error())
	}

	t.Log("Load dict spend(ms):", (time.Now().Nanosecond()-startTime.Nanosecond())/1000000)
}

func TestConvertToSimplifiedChinese(t *testing.T) {
	s := "中華人民共和國"
	simplied, err := ConvertToSimplifiedChinese(s)
	if err != nil {
		t.Fatal("TestConvertToSimplifiedChinese with error: ", err.Error())
	}
	if simplied != "中华人民共和国" {
		t.Fatal("TestConvertToSimplifiedChinese not match: ", simplied)
	}
}

func TestConvertToTraditionalChinese(t *testing.T) {
	s := "中华人民共和国"
	traditional, err := ConvertToTraditionalChinese(s)
	if err != nil {
		t.Fatal("TestConvertToTraditionalChinese with error: ", err.Error())
	}
	if traditional != "中華人民共和國" {
		t.Fatal("TestConvertToTraditionalChinese not match: ", traditional)
	}

}

func TestConvertToPinyinString(t *testing.T) {
	s := "杭州西湖"
	pinyin, err := ConvertToPinyinString(s, " ", FormatWithToneMark)
	if err != nil {
		t.Fatal("ConvertToPinyinString with error: ", err.Error())
	}
	if pinyin != "háng zhōu xī hú" {
		t.Fatal("ConvertToPinyinString WithToneMark not match: ", pinyin)
	}

}
