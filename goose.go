package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

var (
	port    int
	webroot string
)

func main() {
	flag.IntVar(&port, "p", 8080, "Port to bind")
	flag.StringVar(&webroot, "r", "", "Webroot (default cwd)")
	flag.Parse()

	if webroot == "" {
		if len(flag.Args()) > 0 {
			webroot = flag.Args()[0]
		} else {
			cwd, _ := os.Getwd()
			webroot = cwd
		}
	}

	fmt.Printf("Serving from %s\n", webroot)
	host := fmt.Sprintf(":%d", port)
	fmt.Printf("Binding to %s\n", host)

	fileServer := http.FileServer(http.Dir(webroot))
	log.Fatal(http.ListenAndServe(host, handlers.LoggingHandler(os.Stdout, fileServer)))
}
