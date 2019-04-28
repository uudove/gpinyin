package pconv

import (
	"bufio"
	"io"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

var traditionalToSimplifiedMap map[string]string
var simplifiedToTraditionalMap map[string]string
var pinyinMap map[string]string
var multiPinyinMap map[string]string

var multiPinyinKeys []string

var arrayTrie *DoubleArrayTrie

var isChineseMapLoaded bool
var isPinyinMapLoaded bool
var isMultiChineseMapLoaded bool

func LoadAllDict() error {
	err := loadChineseDict()
	if err != nil {
		return err
	}

	err = loadPinyinDict()
	return err
}

func loadChineseDict() error {
	if !isChineseMapLoaded {
		if traditionalToSimplifiedMap == nil {
			traditionalToSimplifiedMap = make(map[string]string)
		}
		if simplifiedToTraditionalMap == nil {
			simplifiedToTraditionalMap = make(map[string]string)
		}
		// init Chinese map dict
		chineseLoadErr := loadDict(traditionalToSimplifiedMap, "data/chinese.dict")
		if chineseLoadErr != nil {
			log.Println(chineseLoadErr.Error())
			return chineseLoadErr
		}

		for k, v := range traditionalToSimplifiedMap {
			simplifiedToTraditionalMap[v] = k
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

	if !isMultiChineseMapLoaded {
		if multiPinyinMap == nil {
			multiPinyinMap = make(map[string]string)
		}

		// init multi pinyin map dict
		multiPinyinLoadErr := loadDict(multiPinyinMap, "data/multi_pinyin.dict")
		if multiPinyinLoadErr != nil {
			log.Println(multiPinyinLoadErr.Error())
			return multiPinyinLoadErr
		}

		// get all multi pinyin keys, and sort them
		multiPinyinKeys = make([]string, len(multiPinyinMap))
		index := 0
		for k := range multiPinyinMap {
			multiPinyinKeys[index] = k
			index++
		}
		sort.Strings(multiPinyinKeys)

		// build multi dict
		arrayTrie = new(DoubleArrayTrie)
		arrayTrie.build(multiPinyinKeys)

		isMultiChineseMapLoaded = true
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
