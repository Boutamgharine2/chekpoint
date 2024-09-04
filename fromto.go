package main

import "fmt"

func FromTo(from int, to int) string {
	res := ""
	if (from > 99 || from < 0) || (to > 99 || to < 0) {
		return "Invalid\n"
	}
	if to >= from {
		for i := from; i <= to; i++ {
			if i < 10 {
				res += "0" + Itoa(i)
			} else {
				res += Itoa(i)
			}
			if i != to {
				res += "," + " "
			}
		}
	}else if to < from {
		for i := from; i >= to; i-- {
			if i < 10 {
				res += "0" + Itoa(i)
			} else {
				res += Itoa(i)
			}
			if i != to {
				res += "," + " "
			}
		}
	}

	res += "\n"
	return res
}

func main() {
	fmt.Print(FromTo(1, 10))
	fmt.Print(FromTo(10, 1))
	fmt.Print(FromTo(10, 10))
	fmt.Print(FromTo(100, 10))
}

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
		res = "-" + res
	}
	return res
}
