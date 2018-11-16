package pconv

import (
	"bufio"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var chineseMap map[string]string
var pinyinMap map[string]string
var mutilPinyinMap map[string]string

var mutilPinyinKeys []string

var isChineseMapLoaded bool
var isPinyinMapLoaded bool
var isMutilChineseMapLoaded bool

func loadChineseDict() error {
	if !isChineseMapLoaded {
		if chineseMap == nil {
			chineseMap = make(map[string]string)
		}

		// init Chinese map dict
		chineseLoadErr := loadDict(chineseMap, "data/chinese.dict")
		if chineseLoadErr != nil {
			log.Println(chineseLoadErr.Error())
			return chineseLoadErr
		}
		isChineseMapLoaded = true
	}
	return nil
}

func loadPinyinDict() error {
	if !isPinyinMapLoaded {
		if pinyinMap == nil {
			pinyinMap = make(map[string]string)
		}

		// init pinyin map dict
		pinyinLoadErr := loadDict(pinyinMap, "data/pinyin.dict")
		if pinyinLoadErr != nil {
			log.Println(pinyinLoadErr.Error())
			return pinyinLoadErr
		}
		isPinyinMapLoaded = true
	}

	if !isMutilChineseMapLoaded {
		if mutilPinyinMap == nil {
			mutilPinyinMap = make(map[string]string)
		}

		// init multi pinyin map dict
		mutilPinyinLoadErr := loadDict(mutilPinyinMap, "data/mutil_pinyin.dict")
		if mutilPinyinLoadErr != nil {
			log.Println(mutilPinyinLoadErr.Error())
			return mutilPinyinLoadErr
		}

		mutilPinyinKeys = make([]string, len(mutilPinyinMap))

		index := 0
		for k, _ := range mutilPinyinMap {
			mutilPinyinKeys[index] = k
			index++
		}

		isMutilChineseMapLoaded = true
	}

	return nil
}

func loadDict(dict map[string]string, path string) error {
	currentPath, err := os.Getwd()
	if err != nil {
		return err
	}

	dictpath := filepath.Join(currentPath, path)
	fi, err := os.Open(dictpath)
	if err != nil {
		return err
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		items := strings.Split(string(a), "=")
		if len(items) == 2 {
			dict[items[0]] = items[1]
		}
	}
	return nil
}
