package Section_1_3

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Dup1 exibe o texto de toda linha que aparece mais de uma vez na entrada-padrão, precedida por sua contagem.
func Dup1() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// Dup2 exibe a contagem e o texto das linhas que aparecem mais de uma vez na entrada. Ele lê de stdin ou de uma lista de arquivos nomeados.
func Dup2() {
	counts := make(map[string]int)
	files := os.Args[1:]

	countLines := func(f *os.File, counts map[string]int) {
		input := bufio.NewScanner(f)
		for input.Scan() {
			counts[input.Text()]++
		}
	}

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func Dup3() {
	counts := make(map[string]int)
	for _, arg := range os.Args[1:] {
		data, err := ioutil.ReadFile(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Fprintf(os.Stdout, "%d\t%s\n", n, line)
		}
	}
}
