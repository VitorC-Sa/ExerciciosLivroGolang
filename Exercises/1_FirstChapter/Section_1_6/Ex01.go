package Section_1_6

import (
	"fmt"
	"os"
	"time"
)

//Modifique fetch all para exibir sua saída em um arquivo para que ela possa ser examinada.
func Ex01() {
	start := time.Now()
	ch := make(chan string)

	f, err := os.Create("./Exercises/1_FirstChapter/Section_1_6/tmp/dat.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error on Ex01 - Section 1.6: %v", err)
		return
	}
	defer f.Close()

	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		s := <-ch
		_, err := f.WriteString(fmt.Sprintf("%s\n", s))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error Writing a string (%s) to File on Ex01 - Section 1.6: %v\n", s, err)
		}
	}

	f.WriteString(fmt.Sprintf("%.2fs elapsed", time.Since(start).Seconds()))
	f.Sync()
}
