package service

import (
	"log"
	"net/http"
)

func Serve(addr string) {
	mux := http.NewServeMux()
	handler := &distilHandler{}

	mux.HandleFunc("POST /solve", handler.Solve)
	mux.HandleFunc("POST /manifest", handler.Encrypt)

	server := http.Server{
		Addr:    addr,
		Handler: mux,
	}

	log.Println("Listen on:", addr)
	log.Fatalln(server.ListenAndServe())
}
