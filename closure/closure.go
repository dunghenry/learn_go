package closure

import "fmt"

func displayNumbers() func() int {
	number := 0
	return func() int {
		number++
		return number
	}
}
func Add(car string) func() []string {
	cars := []string{"Mazda CX-5"}
	return func() []string {
		cars = append(cars, car)
		// fmt.Println(cars[0])
		return cars
	}
}
func Closure() {
	num := displayNumbers()

	fmt.Println(num())
	fmt.Println(num())
	fmt.Println(num())
	cars := Add("BMW")
	cars = Add("Huyndai Santafe 2023")
	for _, car := range cars() {
		fmt.Println(car)
	}
}
