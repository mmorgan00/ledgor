package main

import (
	"context"
	"fmt"
	"github.com/mmorgan00/ledgor/internal/httpapi"
	"github.com/mmorgan00/ledgor/internal/store"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Starting ledger")

	ctx := context.Background()
	st, err := store.New(ctx, os.Getenv("DATABASE_URL")) // New, not Connect
	if err != nil {
		log.Fatalf("Failed to connect: %v", err) // Fatalf formats
	}
	defer st.Close(ctx)

	mux := http.NewServeMux()
	httpapi.New(st).Build(mux)

	srv := &http.Server{Addr: ":5000", Handler: mux}
	log.Fatal(srv.ListenAndServe())
}
