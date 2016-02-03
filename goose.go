package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	flag "snilius.com/pflag"
)

const version = "1.0"
const usage = `
Usage: goose [options...] [webroot]
webroot defaults to current working directory

Options:
`

var (
	port   int
	export bool
)

func main() {
	flag.IntVarP(&port, "port", "p", 8080, "Port to bind")
	flag.BoolVarP(&export, "export", "e", false, "Bind server to 0.0.0.0")
	flag.Version = version
	flag.Usage = func() {
		flag.PrintVersion()
		fmt.Println(usage)
		flag.PrintDefaults()
	}
	flag.Parse()

	var webroot string
	if len(flag.Args()) > 0 {
		webroot = flag.Args()[0]
	} else {
		cwd, _ := os.Getwd()
		webroot = cwd
	}

	fmt.Printf("Serving from %s\n", webroot)
	host := fmt.Sprintf("127.0.0.1:%d", port)
	if export {
		host = fmt.Sprintf("0.0.0.0:%d", port)
	}
	fmt.Printf("Binding to %s\n", host)

	fileServer := http.FileServer(http.Dir(webroot))
	log.Fatal(http.ListenAndServe(host, handlers.LoggingHandler(os.Stdout, fileServer)))
}
