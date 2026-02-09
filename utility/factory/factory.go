package factory

import (
	"errors"
	"fmt"
)

// interface - not init object dc
type Shape interface {
	Draw()
}

// factory - noi quyet dinh tao object nao
// Client ko can biết là circle or square, chỉ biết đó là Shape
func ShapeFactory(t string) (Shape, error) {
	switch t {
	case "circle":
		return Circle{}, nil
	case "square":
		return Square{}, nil
	default:
		return nil, errors.New("invalid type")
	}
}

func FactoryPattern() {
	shape1, err1 := ShapeFactory("circle")
	if err1 != nil {
		fmt.Println("err1:", err1)
	} else {
		shape1.Draw()
	}

	shape2, err2 := ShapeFactory("square")
	if err2 != nil {
		fmt.Println("err2:", err2)
	} else {
		shape2.Draw()
	}

	shape3, err3 := ShapeFactory("toni")
	if err3 != nil {
		fmt.Println("err3:", err3)
	} else {
		shape3.Draw()
	}
}
