
package main

import "fmt"

func main() {
	fmt.Println("Hello world!")

	var num int = 200000
	dog := "Nala"
	var isEpic bool = false

	fmt.Println(isEpic)
	fmt.Println(num)
	fmt.Println(dog)

	if isEpic {
		fmt.Println("You are epic.")
	} else if 5 == 6 {
		fmt.Println("You are kinda epic.")
	} else {
		fmt.Println("You are not epic at all, whatsoever, sorry bruh")
	}
}
