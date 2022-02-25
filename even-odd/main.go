package main

import "fmt"

func main() {
	var s []int
	for i := 0; i < 10; i++ {
		s = append(s, i+1)
	}
	for j := range s {
		if (s[j] % 2 == 0) {
			fmt.Printf("%v is even\n", s[j])
		} else {
			fmt.Printf("%v is odd\n", s[j])
			fmt.Print()
		}
	}
}