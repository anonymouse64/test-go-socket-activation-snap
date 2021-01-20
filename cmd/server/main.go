package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/coreos/go-systemd/activation"
)

func main() {
	listeners, err := activation.Listeners()
	if err != nil {
		log.Fatal(err)
	}

	if len(listeners) == 0 {
		log.Fatal(fmt.Errorf("error no systemd listeners"))
	}

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello!")
	})

	fmt.Printf("Starting server at port %+v\n", listeners[0])
	if err := http.Serve(listeners[0], nil); err != nil {
		log.Fatal(err)
	}
}
