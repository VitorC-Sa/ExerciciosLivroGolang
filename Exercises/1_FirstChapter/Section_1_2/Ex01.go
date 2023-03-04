package Section_1_2

import (
	"fmt"
	"os"
)

//Modifique o programa echo para exibir também os.Args[0], que é o nome do comando que o chamou
func Ex01() {
	s, sep := "", ""

	for _, arg := range os.Args {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
