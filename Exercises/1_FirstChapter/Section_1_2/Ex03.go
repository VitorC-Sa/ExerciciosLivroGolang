package Section_1_2

import (
	"fmt"
	"os"
	"strings"
	"time"
)

//Meça a diferença de tempo de execução entre as três versões de "echo"
func Ex03() {
	start := time.Now()
	firstVersion()
	fmt.Printf("First Version: %v\n", time.Since(start))

	start = time.Now()
	secondVersion()
	fmt.Printf("Second Version: %v\n", time.Since(start))

	start = time.Now()
	thirdVersion()
	fmt.Printf("Third Version: %v\n", time.Since(start))
}

func firstVersion() string {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	return s
}

func secondVersion() string {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	return s
}

func thirdVersion() string {
	return strings.Join(os.Args[1:], " ")
}
