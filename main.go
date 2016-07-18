package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	mess = &Messenger{}
	port string
)

func init() {
	flag.StringVar(&port, "p", ":3000", "webserver port address")
}

func main() {
	flag.Parse()
	fmt.Printf("Webserver ready at %s\n", port)
	server := setupServer(port)
	log.Fatal(server.ListenAndServe())
}

func setupServer(addr string) *http.Server {
	// http.HandleFunc("/fbwebhook", mess.Handler)
	var appServeMux = http.NewServeMux()

	server := &http.Server{
		Addr:    addr,
		Handler: appServeMux,
	}
	appServeMux.Handle("/", handlerRoot())
	appServeMux.HandleFunc("/fbwebhook", mess.Handler)
	return server
}

func handlerRoot() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte("Hello World Update!"))
	})
}
