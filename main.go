package main

import (
	"context"
	"fmt"
	"io"
	"kvdb/db"
	"kvdb/server"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func run(ctx context.Context, w io.Writer, args []string) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	store := db.New()
	srv := server.NewServer(store)

	handler := srv.Route()

	log.Println("Server is running on port 8080")
	return http.ListenAndServe(":8080", handler)
}

func main() {
	ctx := context.Background()
	if err := run(ctx, os.Stdout, os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
}
