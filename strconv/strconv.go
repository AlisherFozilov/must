package strconvMust

import "strconv"

func Atoi(s string) (number int) {
	number, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return
}

func ParseBool(str string) bool {
	parseBool, err := strconv.ParseBool(str)
	if err != nil {
		panic(err)
	}
	return parseBool
}

func ParseFloat(s string, bitSize int) float64 {
	float, err := strconv.ParseFloat(s, bitSize)
	if err != nil {
		panic(err)
	}
	return float
}

func ParseInt(s string, base int, bitSize int) int64 {
	parseInt, err := strconv.ParseInt(s, base, bitSize)
	if err != nil {
		panic(err)
	}
	return parseInt
}

func ParseUint(s string, base int, bitSize int) uint64 {
	parseInt, err := strconv.ParseUint(s, base, bitSize)
	if err != nil {
		panic(err)
	}
	return parseInt
}

func Unquote(s string) string {
	unquote, err := strconv.Unquote(s)
	if err != nil {
		panic(err)
	}
	return unquote
}

func UnquoteChar(s string, quote byte) (value rune, multibyte bool, tail string) {
	char, multibyte, tail, err := strconv.UnquoteChar(s, quote)
	if err != nil {
		panic(err)
	}
	return char, multibyte, tail
}
