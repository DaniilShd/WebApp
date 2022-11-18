package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/DaniilShd/WebApp/pkg/handlers"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)

	fmt.Println(filepath.Dir("/template"))

	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting%s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)

}
