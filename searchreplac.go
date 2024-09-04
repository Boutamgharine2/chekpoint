 package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	// 	c:=""
// 	// 	v:= os.Args[1:]
// 	// 	if len(v)> 3 {
// 	// 		return
// 	// 	}
// 	// 	for i:=0;i<len(v[0]);i++ {
// 	// 	h,e:=verifie(v[0],v[1])
// 	// 	if h && e!= 0 {/*  */
// 	// 		c+=string(v[1])
// 	// 		i++
// 	// 	}
// 	// 	c+=string(v[0][i])

// 	// }
// 	//  fmt.Println(c)
// 	if len(os.Args) != 4 {
// 		return
// 	}
// 	s := os.Args[1]
// 	a := os.Args[2]
// 	b := os.Args[3]
// 	fmt.Println(search(s, a, b))
// }

// // func verifie(s string, r string) bool {
// // 	for i := 0; i < len(s); i++ {
// // 		if string(s[i]) == r {
// // 			return true
// // 		}
// // 	}
// // 	return false
// // }

// func search(s,a,b string ) string {
// 	for i := 0; i < len(s); i++ {
// 		if string(s[i]) == a {
// 			s = s[0:i] + b + s[i+1:]
// 		}
// 	}
// 	return s
// }
