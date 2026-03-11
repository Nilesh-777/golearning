package main

import "fmt"

func main() {
	cards := []string{"Hello", "My", "Name", "Is", "Nilesh"}

	for i, card := range cards {
		fmt.Println(i, card)
	}
}
