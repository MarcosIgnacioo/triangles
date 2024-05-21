package main

import (
	"errors"
	"fmt"
)

const (
	EQUILATERAL = "EQUILATERO"
	ISOSCELES   = "ISOSCELES"
	SCALENE     = "ESCALENO"
)

func main() {
	eq, _ := getTriangleType(1, 2, 2)
	fmt.Println(*eq)
	eq, _ = getTriangleType(1, 1, 1)
	fmt.Println(*eq)
	eq, _ = getTriangleType(2, 3, 1)
	fmt.Println(*eq)
	eq, _ = getTriangleType(2, 3, 2)
	fmt.Println(*eq)
	eq, _ = getTriangleType(2, 3, 2)

	eq, err := getTriangleType(2, -1, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(*eq)
	}
}

func getTriangleType(x float64, y float64, z float64) (*string, error) {
	if x <= 0 || y <= 0 || z <= 0 {
		return nil, errors.New("Una de las medidas son invalidas")
	}
	var triangleType string
	if isEquilateral(x, y, z) {
		triangleType = EQUILATERAL
		return &triangleType, nil
	}
	if isIsosceles(x, y, z) {
		triangleType = ISOSCELES
		return &triangleType, nil
	}
	triangleType = SCALENE
	return &triangleType, nil
}

func isIsosceles(x float64, y float64, z float64) bool {
	return (((x == y) && x != z) || (x == z) && x != y || (y == z) && y != x)
}

func isEquilateral(x float64, y float64, z float64) bool {
	return x == y && x == z
}
