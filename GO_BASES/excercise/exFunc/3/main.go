/*
3 cathegories A,B,C
Salary of A= 3000 per houer + 50% of salary of A
Salary of B= 1500 per houer + 20% of salary of B
Salary of C= 1000 per houer
*/
package main

/*
3000 per houer = 50 per minutes
1500 per houer = 25 per minutes
1000 per houer = 16.66 per minutes
*/
import "fmt"

func salary(category string, minutes int) float32 {
	switch category {
	case "A":
		return 50*float32(minutes) + 0.5*50*float32(minutes)
	case "B":
		return 25*float32(minutes) + 0.2*25*float32(minutes)
	default:
		return 16.66 * float32(minutes)
	}
}
func main() {
	fmt.Println(salary("A", 60))
	fmt.Println(salary("B", 60))
	fmt.Println(salary("C", 60))
}
