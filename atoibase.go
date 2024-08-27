package main

import "fmt"

func ItoaBase(value, base int) string {
	base16 := "0123456789ABCDEF"
	var (
		firstResult string
		result      string
		absolute    int
	)
	absolute = value
	if value < 0 {
		absolute = value * -1
	}
	for absolute > 0 {
		firstResult += string(base16[absolute%base])
		absolute /= base
	}
	for i := len(firstResult) - 1; i >= 0; i-- {
		result += string(firstResult[i])
	}
	if value < 0 {
		result = "-" + result
	}
	return result
}

func main() {
	result := ItoaBase(-11221212065437733299, 16)
	fmt.Println(result)
}
