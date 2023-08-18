package main

import "fmt"

func main() {
	fmt.Println(skobki("()((()))()"))
	fmt.Println(skobki("()(((()))()"))
	fmt.Println(skobki("(()"))
	fmt.Println(skobki(")"))
}

// ()((()))() - ok
// ()(((()))() - не ок

func skobki(skb string) bool {
	stek := make([]byte, 0, len(skb))

	for i := 0; i < len(skb); i++ {
		char := skb[i]
		if char == '(' {
			stek = append(stek, char)
		} else {
			// pop
			if len(stek) == 0 {
				return false
			}
			stek = stek[1:]
		}
	}

	if len(stek) > 0 {
		return false
	}

	return true
}
