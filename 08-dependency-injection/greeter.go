package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func WebGreeter(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func IOGreeter() {
	Greet(os.Stdout, "Guilherme")
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(WebGreeter)))
}
