package core

import "strings"

//调用格式： 〈整数型〉 取文本长度 （文本型 文本数据） - 系统核心支持库->文本操作
//英文名称：len
//取文本型数据的长度，不包含结束0。本命令为初级命令。
//参数<1>的名称为“文本数据”，类型为“文本型（text）”。参数值指定欲检查其长度的文本数据。
//
//操作系统需求： Windows、Linux

func E取文本长度(value string) int {
	return len(string(value))
}

//调用格式： 〈文本型〉 取文本左边 （文本型 欲取其部分的文本，整数型 欲取出字符的数目） - 系统核心支持库->文本操作
//英文名称：left
//返回一个文本，其中包含指定文本中从左边算起指定数量的字符。本命令为初级命令。
//参数<1>的名称为“欲取其部分的文本”，类型为“文本型（text）”。
//参数<2>的名称为“欲取出字符的数目”，类型为“整数型（int）”。
//
//操作系统需求： Windows、Linux
func E取文本左边(欲取其部分的文本 string, 欲取出字符的数目 int) string {
	if E取文本长度(欲取其部分的文本) < 欲取出字符的数目 {
		欲取出字符的数目 = E取文本长度(欲取其部分的文本)
	}
	return string([]byte(欲取其部分的文本)[:欲取出字符的数目])
}

//调用格式： 〈文本型〉 取文本右边 （文本型 欲取其部分的文本，整数型 欲取出字符的数目） - 系统核心支持库->文本操作
//英文名称：right
//返回一个文本，其中包含指定文本中从右边算起指定数量的字符。本命令为初级命令。
//参数<1>的名称为“欲取其部分的文本”，类型为“文本型（text）”。
//参数<2>的名称为“欲取出字符的数目”，类型为“整数型（int）”。
//
//操作系统需求： Windows、Linux

func E取文本右边(欲取其部分的文本 string, 欲取出字符的数目 int) string {
	l := E取文本长度(欲取其部分的文本)
	lpos := l - 欲取出字符的数目
	if lpos < 0 {
		lpos = 0
	}
	return string([]byte(欲取其部分的文本)[lpos:l])
}

//取文本中间


//调用格式： 〈文本型〉 字符 （字节型 欲取其字符的字符代码） - 系统核心支持库->文本操作
//英文名称：chr
//返回一个文本，其中包含有与指定字符代码相关的字符。本命令为初级命令。
//参数<1>的名称为“欲取其字符的字符代码”，类型为“字节型（byte）”。
//
//操作系统需求： Windows、Linux
func E字符(字节型 int8) string {
	return string(byte(字节型))
}

//调用格式： 〈整数型〉 取代码 （文本型 欲取字符代码的文本，［整数型 欲取其代码的字符位置］） - 系统核心支持库->文本操作
//英文名称：asc
//返回文本中指定位置处字符的代码。如果指定位置超出文本长度，返回0。本命令为初级命令。
//参数<1>的名称为“欲取字符代码的文本”，类型为“文本型（text）”。
//参数<2>的名称为“欲取其代码的字符位置”，类型为“整数型（int）”，可以被省略。1为首位置，2为第2个位置，如此类推。如果本参数被省略，默认为首位置。
//
//操作系统需求： Windows、Linux

func E取代码(欲取字符代码的文本 string) int {
	for _, char := range []rune(欲取字符代码的文本) {
		return int(char)
	}
	return 0
}

//调用格式： 〈整数型〉 寻找文本 （文本型 被搜寻的文本，文本型 欲寻找的文本，［整数型 起始搜寻位置］，逻辑型 是否不区分大小写） - 系统核心支持库->文本操作
//英文名称：InStr
//返回一个整数值，指定一文本在另一文本中最先出现的位置，位置值从 1 开始。如果未找到，返回-1。本命令为初级命令。
//参数<1>的名称为“被搜寻的文本”，类型为“文本型（text）”。
//参数<2>的名称为“欲寻找的文本”，类型为“文本型（text）”。
//参数<3>的名称为“起始搜寻位置”，类型为“整数型（int）”，可以被省略。位置值从 1 开始。如果本参数被省略，默认为 1 。
//参数<4>的名称为“是否不区分大小写”，类型为“逻辑型（bool）”，初始值为“假”。为真不区分大小写，为假区分。
//
//操作系统需求： Windows、Linux

func E寻找文本(被搜寻的文本 string, 欲寻找的文本 string) int {
	return strings.Index(被搜寻的文本, 欲寻找的文本)
}

func E倒找文本(被搜寻的文本 string, 欲寻找的文本 string) int {
	return strings.LastIndex(被搜寻的文本, 欲寻找的文本)
}

func E到大写(value string) string {
	return strings.ToUpper(value)
}
func E到小写(value string) string {
	return strings.ToLower(value)
}
func E到全角(value string) string {
	return SBCtoDBC(value)
}
func E到半角(value string) string {
	return DBCtoSBCNew(value)
}
//调用格式： 〈文本型〉 删首空 （文本型 欲删除空格的文本） - 系统核心支持库->文本操作
//英文名称：LTrim
//返回一个文本，其中包含被删除了首部全角或半角空格的指定文本。本命令为初级命令。
//参数<1>的名称为“欲删除空格的文本”，类型为“文本型（text）”。
//
//操作系统需求： Windows、Linux

func E删首空(欲删除空格的文本 string) string {
	return strings.TrimLeft(欲删除空格的文本, " ")
}

//
//调用格式： 〈文本型〉 删尾空 （文本型 欲删除空格的文本） - 系统核心支持库->文本操作
//英文名称：RTrim
//返回一个文本，其中包含被删除了尾部全角或半角空格的指定文本。本命令为初级命令。
//参数<1>的名称为“欲删除空格的文本”，类型为“文本型（text）”。
//
//操作系统需求： Windows、Linux

func E删尾空(欲删除空格的文本 string) string {
	return strings.TrimRight(欲删除空格的文本, " ")
}

func E删首尾空(内容 string) string {
	return strings.TrimSpace(内容)
}

//删全部空

//文本替换
//调用格式： 〈文本型〉 文本替换 （文本型 欲被替换的文本，整数型 起始替换位置，整数型 替换长度，［文本型 用作替换的文本］） - 系统核心支持库->文本操作
//英文名称：ReplaceText
//将指定文本的某一部分用其它的文本替换。本命令为初级命令。
//参数<1>的名称为“欲被替换的文本”，类型为“文本型（text）”。
//参数<2>的名称为“起始替换位置”，类型为“整数型（int）”。替换的起始位置，1为首位置，2为第2个位置，如此类推。
//参数<3>的名称为“替换长度”，类型为“整数型（int）”。
//参数<4>的名称为“用作替换的文本”，类型为“文本型（text）”，可以被省略。如果本参数被省略，则删除文本中的指定部分。
//
//操作系统需求： Windows、Linux


//调用格式： 〈文本型〉 子文本替换 （文本型 欲被替换的文本，文本型 欲被替换的子文本，［文本型 用作替换的子文本］，［整数型 进行替换的起始位置］，［整数型 替换进行的次数］，逻辑型 是否区分大小写） - 系统核心支持库->文本操作
//英文名称：RpSubText
//返回一个文本，该文本中指定的子文本已被替换成另一子文本，并且替换发生的次数也是被指定的。本命令为初级命令。
//参数<1>的名称为“欲被替换的文本”，类型为“文本型（text）”。
//参数<2>的名称为“欲被替换的子文本”，类型为“文本型（text）”。
//参数<3>的名称为“用作替换的子文本”，类型为“文本型（text）”，可以被省略。如果本参数被省略，默认为空文本。
//参数<4>的名称为“进行替换的起始位置”，类型为“整数型（int）”，可以被省略。参数值指定被替换子文本的起始搜索位置。如果省略，默认从 1 开始。
//参数<5>的名称为“替换进行的次数”，类型为“整数型（int）”，可以被省略。参数值指定对子文本进行替换的次数。如果省略，默认进行所有可能的替换。
//参数<6>的名称为“是否区分大小写”，类型为“逻辑型（bool）”，初始值为“真”。为真区分大小写，为假不区分。
//
//操作系统需求： Windows、Linux
func E子文本替换(欲被替换的文本 string, 欲被替换的子文本 string, 用作替换的子文本 string) string {

	return strings.Replace(欲被替换的文本, 欲被替换的子文本, 用作替换的子文本, -1)
}


//调用格式： 〈文本型〉 取空白文本 （整数型 重复次数） - 系统核心支持库->文本操作
//英文名称：space
//返回具有指定数目半角空格的文本。本命令为初级命令。
//参数<1>的名称为“重复次数”，类型为“整数型（int）”。
//
//操作系统需求： Windows、Linux
func E取空白文本(重复次数 int) string {
	var str string
	for i := 0; i < 重复次数; i++ {
		str = str + " "
	}
	return str
}

//调用格式： 〈文本型〉 取重复文本 （整数型 重复次数，文本型 待重复文本） - 系统核心支持库->文本操作
//英文名称：string
//返回一个文本，其中包含指定次数的文本重复结果。本命令为初级命令。
//参数<1>的名称为“重复次数”，类型为“整数型（int）”。
//参数<2>的名称为“待重复文本”，类型为“文本型（text）”。该文本将用于建立返回的文本。如果为空，将返回一个空文本。
//
//操作系统需求： Windows、Linux

func E取重复文本(重复次数 int, 待重复文本 string) string {
	var str string
	for i := 0; i < 重复次数; i++ {
		str = str + 待重复文本
	}
	return str
}

//调用格式： 〈文本型数组〉 分割文本 （文本型 待分割文本，［文本型 用作分割的文本］，［整数型 要返回的子文本数目］） - 系统核心支持库->文本操作
//英文名称：split
//将指定文本进行分割，返回分割后的一维文本数组。本命令为初级命令。
//参数<1>的名称为“待分割文本”，类型为“文本型（text）”。如果参数值是一个长度为零的文本，则返回一个空数组，即没有任何成员的数组。
//参数<2>的名称为“用作分割的文本”，类型为“文本型（text）”，可以被省略。参数值用于标识子文本边界。如果被省略，则默认使用半角逗号字符作为分隔符。如果是一个长度为零的文本，则返回的数组仅包含一个成员，即完整的“待分割文本”。
//参数<3>的名称为“要返回的子文本数目”，类型为“整数型（int）”，可以被省略。如果被省略，则默认返回所有的子文本。
//
//操作系统需求： Windows、Linux

func E分割文本(待分割文本 string, 用作分割的文本 string) []string {
	return strings.Split(待分割文本, 用作分割的文本)
}