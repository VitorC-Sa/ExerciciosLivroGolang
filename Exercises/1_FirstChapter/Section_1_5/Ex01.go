package Section_1_5

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

//A chamada de função io.Copy(dst, src) lê de src e escreve em dst. Use-a no lugar de ioutil.ReadAll para corpiar o body da resposta para os.Stdout
func Ex01() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			continue
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			continue
		}
	}
}
