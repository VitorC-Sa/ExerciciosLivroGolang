package Section_1_7

import (
	"LivroGo/Exercises/1_FirstChapter/Section_1_4"
	"fmt"
	"log"
	"net/http"
	"sync"
)

func Server1() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "URL Path = %q\n", r.URL.Path)
	}

	fmt.Println("Starting server...")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
	fmt.Println("Ending server...")
}

func Server2() {
	var mu sync.Mutex
	var totalCount int
	var counters = make(map[string]int)

	handler := func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Path

		mu.Lock() //Evita que duas requisições simultâneas tentem alterar o valor da variável ao mesmo tempo (racing condition)
		counters[url]++
		totalCount++
		mu.Unlock()

		fmt.Fprintf(w, "URL Path = %q\n", url)
	}

	counter := func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		fmt.Fprintf(w, "Total Count %d\n", totalCount)
		for k, v := range counters {
			fmt.Fprintf(w, "%d\t%s\n", v, k)
		}
		mu.Unlock()
	}

	fmt.Println("Starting server...")
	http.HandleFunc("/count", counter)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
	fmt.Println("Ending server...")
}

func Server3() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s\t%s\t%s\n", r.Method, r.URL, r.Proto)
		for k, v := range r.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
		fmt.Fprintf(w, "Host = %q\n", r.Host)
		fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
		if err := r.ParseForm(); err != nil {
			log.Println(err)
		}
		for k, v := range r.Form {
			fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
		}
	}

	fmt.Println("Starting server...")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
	fmt.Println("Ending server...")
}

func ServerLissajous() {
	fmt.Println("Starting server...")
	defer fmt.Println("Ending server...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		Section_1_4.Lissajous(w)
	})
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
