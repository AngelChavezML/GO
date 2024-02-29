/*
Calculate the minimum,avergage and maximum
*/
package main

import "fmt"

const (
	minimumValue = "minimum"
	averageValue = "average"
	maximumValue = "maximum"
)

func main() {

	const (
		minimum = minimumValue
		average = averageValue
		maximum = maximumValue
	)

	minFunc, err := operation(minimum)

	averageFunc, err := operation(average)

	maxFunc, err := operation(maximum)

	minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)

	averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)

	maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println("Operation successful")
	}
	fmt.Println(minValue)
	fmt.Println(maxValue)
	fmt.Println(averageValue)
}
func operation(operation string) (func(...int) int, error) {
	switch operation {
	case minimumValue:
		return func(numbers ...int) int {
			min := numbers[0]
			for _, number := range numbers {
				if number < min {
					min = number
				}
			}
			return min
		}, nil
	case averageValue:
		return func(numbers ...int) int {
			sum := 0
			for _, number := range numbers {
				sum += number
			}
			return sum / len(numbers)
		}, nil
	case maximumValue:
		return func(numbers ...int) int {
			max := numbers[0]
			for _, number := range numbers {
				if number > max {
					max = number
				}
			}
			return max
		}, nil
	default:
		return nil, fmt.Errorf("operation %s not supported", operation)
	}
}
