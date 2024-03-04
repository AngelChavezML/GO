package main

import (
	"errors"
	"fmt"
)

type ErrorSalaryCustom struct {
	Message string
}

const (
	ErrorSalaryCustomMessage = "Error: the salary entered does not reach the taxable minimum"
	ErrorSalaryLessThan10000 = "Error: salary is less than 10000"
	ErrHours                 = "Error: the worker cannot have worked less than 80 hours per month"
)

// var ErrSalary = ErrorSalaryCustom{Message: ErrorSalaryLessThan10000} //Ejercicio 2
var ErrSalary = errors.New(ErrorSalaryLessThan10000) //Ejercicio 3
func (e ErrorSalaryCustom) Error() string {
	return e.Message
}

func (e ErrorSalaryCustom) New(message string) error {
	return ErrorSalaryCustom{Message: message}
}

func Impuestos(salary float64) (float64, error) {

	/*if salary > 150000 {
		return salary * 0.15, nil
	}*/
	if salary < 100000 {
		//return 0, ErrSalary //Ejercicio 2
		//return 0, errors.New(ErrorSalaryLessThan10000) //Ejercicio 3
		err := fmt.Errorf("%w - Error: the minimum taxable amount is 150,000 and the salary entered is: %.2f", ErrSalary, salary)
		return 0, err
	}
	return 0, errors.New(ErrorSalaryCustomMessage)
}
func CalcularSalario(horasMes int, valorHora float32) (float32, error) { // Ejercicio 5
	var salario float32 = float32(horasMes) * valorHora
	if horasMes < 80 { // Ejercicio 5
		return 0, errors.New(ErrHours)
	}
	if salario > 150000 {
		return salario * 0.9, nil
	}
	if salario < 100000 {
		return 0, ErrSalary
	}
	//err := ErrorSalaryCustom{Message: ErrSalaryNoTaxable} - Ejercicio 1
	//err := errors.New(ErrSalaryNoTaxable) // Ejercicio 3
	err := fmt.Errorf("%w - Error: the minimum taxable amount is 150,000 and the salary entered is: %.2f", ErrSalary, salario)
	return salario, err
}

func main() {

	/*
		Ejercicio 1
		impuestos, error := Impuestos(float64(100000))
		if error != nil {
			fmt.Println(error)
		}
		fmt.Println(impuestos) */
	/* impuestos, error := Impuestos(float64(1000))
	if errors.Is(error, ErrSalary) {
		fmt.Println(error)
	}
	fmt.Println(impuestos) */
	salario, err := CalcularSalario(200, 20000)
	if err != nil {
		fmt.Println(err)
	}
	if errors.Is(err, ErrSalary) {
		fmt.Println(err)
	}
	fmt.Printf("El salario es: %.2f\n", salario)
}

/*
package main
import (
    "errors"
    "fmt"
)
type ErrorSalaryCustom struct {
    Message string
}
func (e ErrorSalaryCustom) Error() string {
    return e.Message
}
const (
    ErrSalaryNoTaxable     = "Error: the salary entered does not reach the taxable minimum"
    ErrSalaryLessThan10000 = "Error: salary is less than 10000"
    ErrHours               = "Error: the worker cannot have worked less than 80 hours per month"
)
// var ErrSalary = ErrorSalaryCustom{Message: ErrSalaryLessThan10000} - Ejercicio 2
var ErrSalary = errors.New(ErrSalaryLessThan10000) // Ejercicio 3
func CalcularSalario(horasMes int, valorHora float32) (float32, error) { // Ejercicio 5
    var salario float32 = float32(horasMes) * valorHora
    if horasMes < 80 { // Ejercicio 5
        return 0, errors.New(ErrHours)
    }
    if salario > 150000 {
        return salario * 0.9, nil
    }
    if salario < 100000 {
        return 0, ErrSalary
    }
    //err := ErrorSalaryCustom{Message: ErrSalaryNoTaxable} - Ejercicio 1
    //err := errors.New(ErrSalaryNoTaxable) // Ejercicio 3
    err := fmt.Errorf("%w - Error: the minimum taxable amount is 150,000 and the salary entered is: %.2f", ErrSalary, salario)
    return salario, err
}
func ImpuestoDeSalario(salario float32) (float32, error) {
    if salario > 150000 {
        return salario * 0.9, nil
    }
    if salario < 100000 {
        return 0, ErrSalary
    }
    //err := ErrorSalaryCustom{Message: ErrSalaryNoTaxable} - Ejercicio 1
    //err := errors.New(ErrSalaryNoTaxable) // Ejercicio 3
    err := fmt.Errorf("%w - Error: the minimum taxable amount is 150,000 and the salary entered is: %.2f", ErrSalary, salario)
    return salario, err
}
func main() {
    // Ejercicio 1
    /*
        var salary int = 100000
        salario, err := ImpuestoDeSalario(float32(salary))
        if err != nil {
            fmt.Println(err)
        } else {
            fmt.Println(salario)
        }
*/
// Ejercicio 2
/*  salario, err := ImpuestoDeSalario(100000)
        if err != nil {
            fmt.Println(err)
        }
        if errors.Is(err, ErrSalary) {
            fmt.Println(err)ß
        }
        fmt.Printf("El salario es: %.2f\n", salario)
    // Ejercicio 5
    salario, err := CalcularSalario(20, 20000)
    if err != nil {
        fmt.Println(err)
    }
    if errors.Is(err, ErrSalary) {
        fmt.Println(err)
    }
    fmt.Printf("El salario es: %.2f\n", salario)
}*/
