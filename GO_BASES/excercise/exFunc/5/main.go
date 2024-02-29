package main

/*
Dog eats 10
Cat eats 5
Hamster eats 0.250
Tarantula eats 0.150
*/
import "fmt"

const (
	dog       = "dog"
	cat       = "cat"
	hamster   = "hamster"
	tarantula = "tarantula"
)

// One function for each animal that return the mount of the food
func animalDog(mount float64) float64 {
	return 10 * mount
}
func animalCat(mount float64) float64 {
	return 5 * mount
}
func animalHamster(mount float64) float64 {
	return 0.250 * mount
}
func animalTarantula(mount float64) float64 {
	return 0.150 * mount
}

// Animal is a function that recive the name of an animal and return another function and a message (if the animal is not suported)
func animal(animal string) (func(float64) float64, string) {
	switch animal {
	case dog:
		return animalDog, "dog"
	case cat:
		return animalCat, "cat"
	case hamster:
		return animalHamster, "hamster"
	case tarantula:
		return animalTarantula, "tarantula"
	default:
		return nil, "animal not suported"
	}
}
func main() {
	animalDog, msg := animal(dog)
	animalCat, msg := animal(cat)
	animalHamster, msg := animal(hamster)
	animalTarantula, msg := animal(tarantula)
	var amount float64
	if msg == "animal not suported" {
		fmt.Println("Animal not suported")
		return
	}
	amount += animalDog(10)
	amount += animalCat(10)
	amount += animalHamster(10)
	amount += animalTarantula(10)
	fmt.Println("The amount of food is", amount)

}
