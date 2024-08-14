package main

import (
	"fmt"
	"os"
)

func main() {
	c := ""
	v := os.Args[1:]
	if len(v) != 1 {
		return
	}
	for i := 0; i < len(v[0]); i++ {
		if i <= len(v[0])-2 && v[0][i] == ' ' && v[0][i+1] == ' ' {
			continue
		} else if v[0][i] == ' ' && c == "" {
			continue
		}
		c += string(v[0][i])
	}
	if c[len(c)-1] == ' ' {
		c = c[:len(c)-2]
	}
	fmt.Println(c)
}
