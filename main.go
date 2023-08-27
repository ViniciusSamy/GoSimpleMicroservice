// main.go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!")
	})

	port := "8080"
	fmt.Printf("Servidor rodando em http://localhost:%s\n", port)
	http.ListenAndServe(":"+port, nil)
}
