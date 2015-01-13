package main

import (
	"flag"
	"log"
	"net/http"
)

//go:generate esc -o static.go -prefix=static static
func main() {
	var (
		useLocalAssets bool
		address        string
	)

	flag.BoolVar(&useLocalAssets, "local", false, "Use local assets intead of the embeded assets")
	flag.StringVar(&address, "http", ":8080", "Address and port to bind the HTTP server to")
	flag.Parse()

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(FS(useLocalAssets)))

	log.Fatal(http.ListenAndServe(address, mux))
}
