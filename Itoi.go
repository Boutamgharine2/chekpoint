package main

import "fmt"

func Itoa(n int) string {
	base := "0123456789"
	res := ""
	isnegativ := false
	if n < 0 {
		n *= -1
		isnegativ = true
	}
	

	if n == 0 {
		return "0"
	}
	for n > 0 {
		res = string(base[n%10]) + res
		n /= 10
	}
	if isnegativ {
		res = "-"+ res
		
	}
	return res
}

func main() {
	fmt.Println(Itoa(12347852326))
	fmt.Println(Itoa(0))
	fmt.Println(Itoa(-1234))
	fmt.Println(Itoa(987654321))
}
