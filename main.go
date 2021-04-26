package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fs := http.FileServer(http.Dir("./_built_"))

	go func() {
		log.Fatal(http.ListenAndServe(":9000", fs))
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
}
