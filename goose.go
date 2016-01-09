package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	port    int
	webroot string
)

func main() {

	flag.IntVar(&port, "p", 8080, "Port to bind")
	flag.StringVar(&webroot, "r", "", "root")
	flag.Parse()

	if webroot == "" {
		if len(os.Args) > 1 {
			webroot = os.Args[1]
		} else {
			cwd, _ := os.Getwd()
			webroot = cwd
		}
	}

	fmt.Printf("Serving from %s\n", webroot)
	host := fmt.Sprintf(":%d", port)
	fmt.Printf("Binding to %s\n", host)

	//panic(http.ListenAndServe(host, http.FileServer(http.Dir(webroot))))
	//fmt.Println(webroot)
	fileServer := http.FileServer(http.Dir(webroot))
	log.Fatal(http.ListenAndServe(host, fileServer))
}
