package main

// import "fmt"

// func main() {
// 	fmt.Println(RepeatAlpha("abc"))
// 	fmt.Println(RepeatAlpha("Choumi."))
// 	fmt.Println(RepeatAlpha(""))
// 	fmt.Println(RepeatAlpha("abacadaba 01!"))
// }

func RepeatAlpha(s string) string {
	str := ""
	for i := 0; i < len(s); i++ {
		str += canada(rune(s[i]))
	}
	return str
}

func canada(r rune) string {
	s := ""
	var v rune
	if r >= 'a' && r <= 'z' {
		v = r - 96
		for i := 0; i < int(v); i++ {
			s += string(r)
		}
	} else if r >= 'A' && r <= 'Z' {
		v = r - 64
		for i := 0; i < int(v); i++ {
			s += string(r)
		}
	} else {
		s += string(r)
	}
	return s
}
