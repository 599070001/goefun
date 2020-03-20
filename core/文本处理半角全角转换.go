package E

import (
	"bytes"
	"sort"
)

type dict struct {
	x, y rune
}

var toSBC_ = []dict{
	{0x00A2, 0xFFE0},
	{0x00A3, 0xFFE1},
	{0x00A5, 0xFFE5},
	{0x00A6, 0xFFE4},
	{0x00AC, 0xFFE2},
	{0x00AF, 0xFFE3},
	{0xFF61, 0x3002},
	{0xFF62, 0x300C},
	{0xFF63, 0x300D},
	{0xFF64, 0x3001},
	{0xFF65, 0x30FB},
	{0xFF66, 0x30F2},
	{0xFF67, 0x30A1},
	{0xFF68, 0x30A3},
	{0xFF69, 0x30A5},
	{0xFF6A, 0x30A7},
	{0xFF6B, 0x30A9},
	{0xFF6C, 0x30E3},
	{0xFF6D, 0x30E5},
	{0xFF6E, 0x30E7},
	{0xFF6F, 0x30C3},
	{0xFF70, 0x30FC},
	{0xFF71, 0x30A2},
	{0xFF72, 0x30A4},
	{0xFF73, 0x30A6},
	{0xFF74, 0x30A8},
	{0xFF75, 0x30AA},
	{0xFF76, 0x30AB},
	{0xFF77, 0x30AD},
	{0xFF78, 0x30AF},
	{0xFF79, 0x30B1},
	{0xFF7A, 0x30B3},
	{0xFF7B, 0x30B5},
	{0xFF7C, 0x30B7},
	{0xFF7D, 0x30B9},
	{0xFF7E, 0x30BB},
	{0xFF7F, 0x30BD},
	{0xFF80, 0x30BF},
	{0xFF81, 0x30C1},
	{0xFF82, 0x30C4},
	{0xFF83, 0x30C6},
	{0xFF84, 0x30C8},
	{0xFF85, 0x30CA},
	{0xFF86, 0x30CB},
	{0xFF87, 0x30CC},
	{0xFF88, 0x30CD},
	{0xFF89, 0x30CE},
	{0xFF8A, 0x30CF},
	{0xFF8B, 0x30D2},
	{0xFF8C, 0x30D5},
	{0xFF8D, 0x30D8},
	{0xFF8E, 0x30DB},
	{0xFF8F, 0x30DE},
	{0xFF90, 0x30DF},
	{0xFF91, 0x30E0},
	{0xFF92, 0x30E1},
	{0xFF93, 0x30E2},
	{0xFF94, 0x30E4},
	{0xFF95, 0x30E6},
	{0xFF96, 0x30E8},
	{0xFF97, 0x30E9},
	{0xFF98, 0x30EA},
	{0xFF99, 0x30EB},
	{0xFF9A, 0x30EC},
	{0xFF9B, 0x30ED},
	{0xFF9C, 0x30EF},
	{0xFF9D, 0x30F3},
	{0xFF9E, 0x309B},
	{0xFF9F, 0x309C},
	{0xFFA0, 0x3164},
	{0xFFC2, 0x314F},
	{0xFFC3, 0x3150},
	{0xFFC4, 0x3151},
	{0xFFC5, 0x3152},
	{0xFFC6, 0x3153},
	{0xFFC7, 0x3154},
	{0xFFCA, 0x3155},
	{0xFFCB, 0x3156},
	{0xFFCC, 0x3157},
	{0xFFCD, 0x3158},
	{0xFFCE, 0x3159},
	{0xFFCF, 0x315A},
	{0xFFD2, 0x315B},
	{0xFFD3, 0x315C},
	{0xFFD4, 0x315D},
	{0xFFD5, 0x315E},
	{0xFFD6, 0x315F},
	{0xFFD7, 0x3160},
	{0xFFDA, 0x3161},
	{0xFFDB, 0x3162},
	{0xFFDC, 0x3163},
	{0xFFE8, 0x2502},
	{0xFFE9, 0x2190},
	{0xFFEA, 0x2191},
	{0xFFEB, 0x2192},
	{0xFFEC, 0x2193},
	{0xFFED, 0x25A0},
	{0xFFEE, 0x25CB},
}

var toDBC_ = []dict{
	{0x2190, 0xFFE9},
	{0x2191, 0xFFEA},
	{0x2192, 0xFFEB},
	{0x2193, 0xFFEC},
	{0x3000, 0x0020},
	{0x3001, 0xFF64},
	{0x3002, 0xFF61},
	{0x300C, 0xFF62},
	{0x300D, 0xFF63},
	{0x309B, 0xFF9E},
	{0x309C, 0xFF9F},
	{0x30A1, 0xFF67},
	{0x30A2, 0xFF71},
	{0x30A3, 0xFF68},
	{0x30A4, 0xFF72},
	{0x30A5, 0xFF69},
	{0x30A6, 0xFF73},
	{0x30A7, 0xFF6A},
	{0x30A8, 0xFF74},
	{0x30A9, 0xFF6B},
	{0x30AA, 0xFF75},
	{0x30AB, 0xFF76},
	{0x30AD, 0xFF77},
	{0x30AF, 0xFF78},
	{0x30B1, 0xFF79},
	{0x30B3, 0xFF7A},
	{0x30B5, 0xFF7B},
	{0x30B7, 0xFF7C},
	{0x30B9, 0xFF7D},
	{0x30BB, 0xFF7E},
	{0x30BD, 0xFF7F},
	{0x30BF, 0xFF80},
	{0x30C1, 0xFF81},
	{0x30C3, 0xFF6F},
	{0x30C4, 0xFF82},
	{0x30C6, 0xFF83},
	{0x30C8, 0xFF84},
	{0x30CA, 0xFF85},
	{0x30CB, 0xFF86},
	{0x30CC, 0xFF87},
	{0x30CD, 0xFF88},
	{0x30CE, 0xFF89},
	{0x30CF, 0xFF8A},
	{0x30D2, 0xFF8B},
	{0x30D5, 0xFF8C},
	{0x30D8, 0xFF8D},
	{0x30DB, 0xFF8E},
	{0x30DE, 0xFF8F},
	{0x30DF, 0xFF90},
	{0x30E0, 0xFF91},
	{0x30E1, 0xFF92},
	{0x30E2, 0xFF93},
	{0x30E3, 0xFF6C},
	{0x30E4, 0xFF94},
	{0x30E5, 0xFF6D},
	{0x30E6, 0xFF95},
	{0x30E7, 0xFF6E},
	{0x30E8, 0xFF96},
	{0x30E9, 0xFF97},
	{0x30EA, 0xFF98},
	{0x30EB, 0xFF99},
	{0x30EC, 0xFF9A},
	{0x30ED, 0xFF9B},
	{0x30EF, 0xFF9C},
	{0x30F2, 0xFF66},
	{0x30F3, 0xFF9D},
	{0x30FB, 0xFF65},
	{0x30FC, 0xFF70},
	{0x314F, 0xFFC2},
	{0x3150, 0xFFC3},
	{0x3151, 0xFFC4},
	{0x3152, 0xFFC5},
	{0x3153, 0xFFC6},
	{0x3154, 0xFFC7},
	{0x3155, 0xFFCA},
	{0x3156, 0xFFCB},
	{0x3157, 0xFFCC},
	{0x3158, 0xFFCD},
	{0x3159, 0xFFCE},
	{0x315A, 0xFFCF},
	{0x315B, 0xFFD2},
	{0x315C, 0xFFD3},
	{0x315D, 0xFFD4},
	{0x315E, 0xFFD5},
	{0x315F, 0xFFD6},
	{0x3160, 0xFFD7},
	{0x3161, 0xFFDA},
	{0x3162, 0xFFDB},
	{0x3163, 0xFFDC},
	{0x3164, 0xFFA0},
	{0xFF5F, 0x2985},
	{0xFF60, 0x2986},
	{0xFFE0, 0x00A2},
	{0xFFE1, 0x00A3},
	{0xFFE2, 0x00AC},
	{0xFFE3, 0x00AF},
	{0xFFE4, 0x00A6},
	{0xFFE5, 0x00A5},
	{0xFFE6, 0x20A9},
}

// 判断一个字符是否是半角字符
func isSBC(c rune) bool {
	switch {
	case c <= 0x007E:
	case c == 0x00A2 || c == 0x00A3:
	case c == 0x00A5 || c == 0x00A6:
	case c == 0x00AC || c == 0x00AF:
	case c == 0x20A9:
	case c == 0x2985 || c == 0x2986:
	case c >= 0xFF61 && c <= 0xFF9F:
	case c >= 0xFFA0 && c <= 0xFFBE:
	case c >= 0xFFC2 && c <= 0xFFC7:
	case c >= 0xFFCA && c <= 0xFFCF:
	case c >= 0xFFD2 && c <= 0xFFD7:
	case c >= 0xFFDA && c <= 0xFFDC:
	case c >= 0xFFE8 && c <= 0xFFEE:
	default:
		return false
	}
	return true
}

// 如果字符c为全角字符，返回对应半角字符（如有）；否则返回c。
func toDBC(c rune) rune {

	switch {
	case c <= 0x218F:
		return c
	case c <= 0x2193:
	case c == 0x2502:
		return 0xFFE8
	case c == 0x25A0:
		return 0xFFED
	case c == 0x25CB:
		return 0xFFEE
	case c <= 0x2FFF:
		return c
	case c <= 0x30FC:
	case c <= 0x3130:
		return c
	case c <= 0x314E:
		return c + 0xCE70
	case c <= 0x3164:
	case c <= 0xFF00:
		return c
	case c <= 0xFF5E:
		return c - 0xFEE0
	case c >= 0xFFE7:
		return c
	}
	i := sort.Search(len(toDBC_), func(i int) bool { return toDBC_[i].x >= c })
	if toDBC_[i].x == c {
		return toDBC_[i].y
	} else {
		return c
	}
}

// 如果字符c为半角字符，返回对应全角字符（如有）；否则返回c。
func toSBC(c rune) rune {

	switch {
	case c <= 0x001F:
		return c
	case c == 0x0020:
		return 0x3000
	case c <= 0x007E:
		return c + 0xFEE0
	case c <= 0x00AF:
	case c == 0x20A9:
		return 0xFFE6
	case c == 0x2985:
		return 0xFF5F
	case c == 0x2986:
		return 0xFF60
	case c <= 0xFF60:
		return c
	case c <= 0xFFA0:
	case c <= 0xFFBE:
		return c - 0xCE70
	case c >= 0xFFEF:
		return c
	}
	i := sort.Search(len(toSBC_), func(i int) bool { return toSBC_[i].x >= c })
	if toSBC_[i].x == c {
		return toSBC_[i].y
	} else {
		return c
	}
}

// 如果字符c为全角ASCII字符，返回对应的半角字符；否则返回c。
func toASCIIDBC(r rune) rune {
	if r == 0x3000 {
		return 0x3000
	}
	if r >= 0xFF01 && r <= 0xFF5E {
		return r - 0xFEE0
	}
	return r
}

// 如果字符c为半角ASCII字符，返回对应的全角字符；否则返回c。
func toASCIISBC(r rune) rune {
	if r == 0x0020 {
		return 0x3000
	}
	if r >= 0x0021 && r <= 0x007E {
		return r + 0xFEE0
	}
	return r
}

//全角转换半角
func dBCtoSBCNew(s string) string {
	buffer := bytes.Buffer{}
	for _, i := range s {
		inside_code := i
		if inside_code == 12288 {
			inside_code = 32
		} else {
			inside_code -= 65248
		}
		if inside_code < 32 || inside_code > 126 {
			buffer.WriteRune(i)
		} else {
			buffer.WriteRune(inside_code)
		}
	}
	return buffer.String()
}

//半角转换全角
func sBCtoDBC(s string) string {
	buffer := bytes.Buffer{}
	for _, i := range s {
		buffer.WriteRune(toASCIISBC(i))
	}
	return buffer.String()
}
