package Section_1_5

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//Modifique Fetch() para exibir também o status code da resposta
func Ex03() {
	for _, url := range os.Args[1:] {
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
		fmt.Println(resp.StatusCode)
	}
}
