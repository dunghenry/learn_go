package errors

import (
	"errors"
	"fmt"
)

func checkName(name string) error {
	newErr := errors.New("Invalid name")
	if name != "nam" {
		return newErr
	}
	return nil
}
func divide(a, b int) error {
	if b == 0 {
		return fmt.Errorf("%d / %d\nCannot divide a number by zero", a, b)
	}
	fmt.Println(float32(a / b))
	return nil
}

type DivisionByZero struct {
	msg string
}

func (z DivisionByZero) Error() string {
	return "Number cannot be division by zero"
}
func divide1(a, b float64) (float32, error) {
	if b == 0 {
		return 0, &DivisionByZero{}
	} else {
		return float32(a / b), nil
	}
}
func Error() {
	msg := "Hello"
	myErr := errors.New("Wrong message")
	name := "nam"
	err := checkName(name)
	age := -14
	error := fmt.Errorf("%d is negative\nAge can't be negative", age)
	if age < 0 {
		fmt.Println(error)
	} else {
		fmt.Println("Age: %d", age)
	}
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Valid name")
	}
	if msg != "Hi" {
		fmt.Println(myErr)
	}
	er := divide(4, 1)
	if er != nil {
		fmt.Printf("error: %s", er)
	} else {
		fmt.Println("Valid division")
	}
	num1 := 5.0
	num2 := 2.0
	rs, err := divide1(num1, num2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(rs)
	}
}
