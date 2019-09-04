# gpinyin

A pinyin converter project written in golang.

中文文档点 [这里](README_zh.md)

## Features

1. Convert a string to pinyin
2. Convert Simplified Chinese to Traditional Chinese
3. Convert Traditional Chinese to Simplified Chinese

## Quick Start

#### Download and install

```
go get github.com/uudove/gpinyin
```

If you're using go modules, just add this below in your `go.mod` file

```
require github.com/uudove/gpinyin v1.0.2
```

#### Import gpinyin

```
import "github.com/uudove/gpinyin"
```

#### Now happy to use

1. Convert a string to pinyin
   
```
s := "杭州西湖"
pinyin, err := gpinyin.ConvertToPinyinString(s, " ", gpinyin.FormatWithToneMark) // pinyin = "háng zhōu xī hú"
pinyin, err := gpinyin.ConvertToPinyinString(s, ",", gpinyin.FormatWithoutTone) // pinyin = "hang,zhou,xi,hu"
pinyin, err := gpinyin.ConvertToPinyinString(s, "", gpinyin.FormatWithToneMark) // pinyin = "hang2zhou1xi1hu2"
```

2. Convert Simplified Chinese to Traditional Chinese

```
s := "我爱你"
traditional, err := gpinyin.ConvertToTraditionalChinese(s) // traditional = "我愛你"
```

3. Convert Traditional Chinese to Simplified Chinese

```
s := "我愛你"
simplified, err := gpinyin.ConvertToSimplifiedChinese(s) // simplified = "我爱你"
```


## License

This project is licensed under the Apache Licence, Version 2.0 (https://www.apache.org/licenses/LICENSE-2.0.html).
