package main

import (
	"fmt"
	"os"
)

func main() {
	v := os.Args[1:]
	c := v[0]
	C := ""
	if len(v) != 1 {
		return
	}
	for i := 0; i < len(c); i++ {
		if c[i] == ' ' && C == "" {
			continue
		}
		if i <= len(c)-2 && string(c[i]) == string(' ') && string(c[i+1]) != string(' ') {
			C += string(' ') + string(' ')
		}

		C += string(c[i])

	}
	for C[len(C)-1] == ' ' {
		C = C[:len(C)-1]
	}
	fmt.Println(C)
}
