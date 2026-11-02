package main

import (
"fmt"
"log"
"net/http"
)

func main() {
fs := http.FileServer(http.Dir("./public"))
http.Handle("/", fs)

fmt.Println("Servidor rodando em http://localhost:8080")
log.Fatal(http.ListenAndServe(":8080", nil))
}
