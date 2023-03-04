package Section_1_3

import (
	"bufio"
	"fmt"
	"os"
)

//Modifique dup2 para que exiba os nomes de todos os arquivos em que cada linha duplicada ocorre
func Ex01() {
	var lineCounters = make(map[string]map[string]int)

	readLinesAndParseMap := func(f *os.File, fileName string) {
		buffer := bufio.NewScanner(f)
		buffer.Split(bufio.ScanLines)
		for buffer.Scan() {
			line := buffer.Text()
			if _, exists := lineCounters[line]; !exists {
				lineCounters[line] = make(map[string]int)
			}
			lineCounters[line][fileName]++
		}
	}

	for _, fileName := range os.Args[1:] {
		f, err := os.Open(fileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		}
		defer f.Close()

		readLinesAndParseMap(f, fileName)
	}

	for line, m := range lineCounters {
		var counter int
		for _, qtd := range m {
			counter += qtd
		}
		if counter <= 1 {
			continue
		}

		fmt.Printf("%s line appears %d times\n", line, counter)
		for k, v := range m {
			fmt.Printf("%d time(s) in the file %s\n", v, k)
		}
		fmt.Println()
	}
}
