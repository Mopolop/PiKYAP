package main

import (
	"errors"
	"fmt"
)

type Rectangle struct {
	width  float64
	height float64
}

func validateDimensions(width, height float64) error {
	if width <= 0 || height <= 0 {
		return errors.New("ширина и высота должны быть положительными числами")
	}
	return nil
}

func getArea(r Rectangle) float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

type ColoredRectangle struct {
	Rectangle Rectangle
	color     string
}

func main() {

	rect1 := Rectangle{width: 5, height: 10}
	fmt.Println("Rectangle 1:", rect1)

	width := 7.5
	height := 12.0
	if err := validateDimensions(width, height); err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	rect2 := Rectangle{width: width, height: height}
	fmt.Println("Rectangle 2:", rect2)

	var rect3 Rectangle
	if err := validateDimensions(rect3.width, rect3.height); err != nil {
		fmt.Println("Ошибка в Rectangle 3:", err)
	} else {
		fmt.Println("Rectangle 3:", rect3)
	}

	fmt.Println("Area of rect1:", getArea(rect1))
	fmt.Println("Perimeter of rect1:", rect1.Perimeter())

	fmt.Println("Area of rect2:", getArea(rect2))
	fmt.Println("Perimeter of rect2:", rect2.Perimeter())

	coloredRect := ColoredRectangle{Rectangle: Rectangle{width: 3, height: 4}, color: "red"}
	if err := validateDimensions(coloredRect.Rectangle.width, coloredRect.Rectangle.height); err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Colored Rectangle:", coloredRect)
		fmt.Println("Area of coloredRect:", getArea(coloredRect.Rectangle))
		fmt.Println("Perimeter of coloredRect:", coloredRect.Rectangle.Perimeter())
	}
	rect4 := Rectangle{width: 5, height: 10}
	rect5 := Rectangle{width: 5, height: 10}
	fmt.Println("rect4 == rect5:", rect4 == rect5)
}
