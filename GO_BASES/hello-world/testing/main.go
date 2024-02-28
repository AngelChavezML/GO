/*
Notas: No puede existir dos main.go en la misma carpeta

*/
package main
//Importar multiples dependencias
//Para nuestras dependencias usar nombre_del_modulo/ruta_de_carpeta_donde_esta
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
