package main

import "fmt"

func main() {
	var x, rank int

	fmt.Print("Введите трехзначное число: ")
	fmt.Scan(&x)

	for i := 100; i >= 1; i = i / 10 {
		rank = x / i
		fmt.Println(rank)
		x = x - rank*i
	}
}
