package main

import "fmt"

func takeList(str []string) map[string]int {
	data := map[string]int{}
	for _, value := range str {
		for _, character := range []rune(value) {
			if character == 'a' || character == 'e' || character == 'i' || character == 'o' || character == 'u' ||
				character == 'A' || character == 'E' || character == 'I' || character == 'O' || character == 'U' {
				data[value]++
			}
		}
	}
	return data
}

func main() {
	str := []string{"Ana", "mere", "peree"}

	fmt.Println(takeList(str))
}
