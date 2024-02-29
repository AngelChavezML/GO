// If salary is greater than 50000 it has a descount of 17% and if is greatter than 150000 discount of 27%
package main

import "fmt"

func discount(salary float64) (nothing float64) {
	if salary > 50000 {
		return salary * 0.17
	}
	if salary > 150000 {
		return salary * 0.27
	}
	return
}
func main() {
	fmt.Println(discount(510000))
}
