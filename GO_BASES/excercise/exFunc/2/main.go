/*
Averge of N numbers integers, note there musnt be negative numbers
*/
package main

import "fmt"

func average(numbers []int) (average float64, err string) {
	for _, number := range numbers {
		if number < 0 {
			return 0, "negative number"
		}
		average += float64(number)
	}
	average /= float64(len(numbers))
	return
}
func main() {
	fmt.Println(average([]int{1, 2, 3, 4, 5}))
	fmt.Println(average([]int{1, 2, 3, 4, -5}))
}
