package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Println("Введите коэффициенты квадратного уравнения ax^2 + bx + c = 0:")
	fmt.Print("a = ")
	fmt.Scanln(&a)
	fmt.Print("b = ")
	fmt.Scanln(&b)
	fmt.Print("c = ")
	fmt.Scanln(&c)

	if a == 0 {
		fmt.Println("Уравнение не является квадратным")
		return
	}

	d := b*b - 4*a*c

	if d > 0 {
		x1 := (-b + math.Sqrt(d)) / (2 * a)
		x2 := (-b - math.Sqrt(d)) / (2 * a)
		fmt.Printf("Корни уравнения: x1 = %.2f, x2 = %.2f\n", x1, x2)
	} else if d == 0 {
		x := -b / (2 * a)
		fmt.Printf("Корни уравнения: x = %.2f\n", x)
	} else {
		fmt.Println("Уравнение не имеет действительных корней, хотите посчитать комплексные корни?(y/n)")
		var e byte
		fmt.Scanf("%c", &e)
		if e == 'n' {
			return
		} else {
			realPart := -b / (2 * a)
			imagPart := math.Sqrt(math.Abs(d)) / (2 * a)
			x1 := complex(realPart, imagPart)
			x2 := complex(realPart, -imagPart)
			fmt.Printf("Корни уравнения: x1 = %.3v, x2 = %.3v\n", x1, x2)
		}
	}
}
