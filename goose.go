package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	flag "github.com/victorhaggqvist/pflag"
)

const version = "0.1.1"
const usage = `
Usage: goose [options...] [webroot]
webroot defaults to current working directory

Options:
`

var (
	port   int
	export bool
	quiet  bool
)

func main() {
	flag.IntVarP(&port, "port", "p", 8080, "Port to bind")
	flag.BoolVarP(&export, "export", "e", false, "Bind server to 0.0.0.0")
	flag.BoolVarP(&quiet, "quiet", "q", false, "Run in quiet mode, ie. no logs")
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
	if !quiet {
		fileServer = handlers.LoggingHandler(os.Stdout, fileServer)
	}
	log.Fatal(http.ListenAndServe(host, fileServer))
}
