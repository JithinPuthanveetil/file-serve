package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/file-serve/handler/config"

	"github.com/file-serve/handler/v1"
)

func init() {
	pathFlag := flag.String("p", "", "path to the directory")
	flag.Parse()
	config.Root = *pathFlag
}

func main() {
	r := v1.Route()
	fmt.Println("Server listening on 3000")
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		fmt.Println("failed to start server:", err)
	}
}
