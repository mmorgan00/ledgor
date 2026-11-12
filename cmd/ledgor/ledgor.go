package main

import (
	"fmt"
	"github.com/mmorgan00/ledgor/internal/httpapi"
	"github.com/mmorgan00/ledgor/internal/store"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Starting ledger")
	_, err := store.Connect(os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println("Failed to connect %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Connected to DB")
	mux := http.NewServeMux()
	httpapi.Build(mux)
	srv := &http.Server{Addr: ":5000", Handler: mux}
	log.Fatal(srv.ListenAndServe())
}
