package main

import (
	"fmt"
)

func ZipString(s string) string {
	c := ""
	res := ""
	resF := ""

	for i := 0; i < len(s); i++ {

		if s[i] != ' ' {
			c += string(s[i])
		}
		if (i == len(s)-1 && c != "") || (s[i] == ' ' && c != "") {
	

			res = res+""+ Calcul(c)
			resF = res 
			res = ""
		}
	}
	return resF
}

func Calcul(s string) string {
	h := 0
	c := ""
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] && !Contien(s[i], c) {
				h++
			}
		}
		if !Contien(s[i], c) {
			c += string(49+h) + string(s[i])
		}
		h = 0
	}

	return c 
}

func Contien(b byte, s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == b {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick"))
	// fmt.Println(ZipString("Helloo Therre!"))
}
