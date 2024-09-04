package main

import "fmt"

func ZipString(s string) string {
	c := 1
	C := ""
	for i := 0; i < len(s); i++ {
		if i!=0 && s[i]==s[i-1] {
			continue
		}else if len(s) > i+c {
			for s[i] == s[i+c] {
				c++
			}
		}
		C += string(48+c) + string(s[i])
		c = 1

	}
	return C
}

// func Calcul(s string) string {
// 	h := 0
// 	c := ""
// 	for i := 0; i < len(s); i++ {
// 		for j := i + 1; j < len(s); j++ {
// 			if s[i] == s[j] && !Contien(s[i], c) {
// 				h++
// 			}
// 		}
// 		if !Contien(s[i], c) {
// 			c += string(49+h) + string(s[i])
// 		}
// 		h = 0
// 	}

// 	return c
// }

func Contien(b byte, s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == b {

			return true
		}
	}
	return false
}

func main() {
	fmt.Println(ZipString("YouuungFe999999977llllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))
}
