package main

import "fmt"

func main() {
	var x, y float32
	fmt.Print("Введите длины сторон прямоугольника (разделенные пробелом): ")
	fmt.Scan(&x, &y)
	fmt.Printf("Площадь: %.2f", RectangleSquare(x, y))
}

func RectangleSquare(x float32, y float32) float32 {
	return x * y
}
