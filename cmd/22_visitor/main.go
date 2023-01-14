package visitor

import "fmt"

/*
クライアント・コード
*/
func main() {
	square := &Square{side: 2}
	circle := &Circle{radius: 3}
	rectangle := &Rectangle{l: 2, b: 3}

	areaCalculator := &AreaCalculator{}
	square.accept(areaCalculator)
	circle.accept(areaCalculator)
	rectangle.accept(areaCalculator)

	fmt.Println()
	middleCoordinates := &MiddleCoordinates{}
	square.accept(middleCoordinates)
	circle.accept(middleCoordinates)
	rectangle.accept(middleCoordinates)
}

/*
実行結果
Calculating area for square
Calculating area for circle
Calculating area for rectangle

Calculating middle point coordinates for square
Calculating middle point coordinates for circle
Calculating middle point coordinates for rectangle
*/
