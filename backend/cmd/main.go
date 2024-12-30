package main

import (
	"fmt"
	"net/http"

	"github.com/Vitaljano/invly/backend/config"
)

func main() {
	conf := config.Load()

	mux := http.NewServeMux()

	addr := fmt.Sprintf(":%s", conf.Port)
	server := http.Server{
		Addr:    addr,
		Handler: mux,
	}

	fmt.Println("Server start on port", addr)
	server.ListenAndServe()
}
