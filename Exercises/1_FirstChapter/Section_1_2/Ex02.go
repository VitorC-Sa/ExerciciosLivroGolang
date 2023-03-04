package Section_1_2

import (
	"fmt"
	"os"
)

//Modifique o programa para exibir o índice e o valor de cada um de seus argumentos, um por linha
func Ex02() {
	for i, arg := range os.Args {
		fmt.Println(i, arg)
	}
}
