/*
Notas: No puede existir dos main.go en la misma carpeta
//declaracion del paquete
*/
package main
//Importar multiples dependencias
import (
	"fmt"
	"github.com/fatih/color"

	"go-bases/intro/calculator"
)
//import "fmt"
func main(){
	color.Red("Hello world!")
	fmt.Println(calculator.Value)
}