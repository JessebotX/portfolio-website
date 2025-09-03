package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	port := ":3000"
	staticFilesDir := filepath.Join("frontend", "dist")

	http.Handle("/", http.FileServer(http.Dir(staticFilesDir)))

	fmt.Printf("Serving on address %s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
