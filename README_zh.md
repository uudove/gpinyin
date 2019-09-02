# gpinyin

go语言版本的pinyin转换库


## 主要功能

1. 将中文（简体）转换为拼音
2. 将简体中文转换为繁体中文
3. 将繁体中文转换为简体中文

## 快速开始

#### 下载及安装

```
go get github.com/uudove/gpinyin
```

如果你在使用go modules, 只需要在你的 `go.mod` 文件中添加如下代码

```
require github.com/uudove/gpinyin v1.0.0
```

#### 导入 gpinyin

```
import "github.com/uudove/gpinyin"
```

#### 现在可以愉快的玩耍了

1. 将中文（简体）转换为拼音
   
```
s := "杭州西湖"
pinyin, err := gpinyin.ConvertToPinyinString(s, " ", gpinyin.FormatWithToneMark) // pinyin = "háng zhōu xī hú"
pinyin, err := gpinyin.ConvertToPinyinString(s, ",", gpinyin.FormatWithoutTone) // pinyin = "hang,zhou,xi,hu"
pinyin, err := gpinyin.ConvertToPinyinString(s, "", gpinyin.FormatWithToneMark) // pinyin = "hang2zhou1xi1hu2"
```

2. 将简体中文转换为繁体中文

```
s := "我爱你"
traditional, err := gpinyin.ConvertToTraditionalChinese(s) // traditional = "我愛你"
```

3. 将繁体中文转换为简体中文

```
s := "我愛你"
simplified, err := gpinyin.ConvertToSimplifiedChinese(s) // simplified = "我爱你"
```

## License

This project is licensed under the Apache Licence, Version 2.0 (https://www.apache.org/licenses/LICENSE-2.0.html).
