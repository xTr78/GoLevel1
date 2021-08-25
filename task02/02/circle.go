package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64

	fmt.Print("Введите площад круга: ")
	fmt.Scan(&x)

	fmt.Printf("Диаметр: %.2f\n", 2*math.Sqrt((x/math.Pi)))
	fmt.Printf("Длина окружности: %.2f", 2*math.Sqrt((x*math.Pi)))
}
