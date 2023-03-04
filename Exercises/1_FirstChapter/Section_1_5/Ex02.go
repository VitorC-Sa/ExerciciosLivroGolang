package Section_1_5

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

//Modifique Fetch() para que o prefixo "http://" seja adicionado a cada URL de argumento, caso esteja faltando
func Ex02() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			url = fmt.Sprintf("http://%s", url)
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			continue
		}
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			continue
		}
		defer resp.Body.Close()
		fmt.Println(string(b))
	}
}
