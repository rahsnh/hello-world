package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// Read dataservice URL from env variable
	dataServicehost := os.Getenv("DATASERVICE_HOST")
	port := os.Getenv("PORT")

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		// Call dataservice
		resp, err := http.Get(dataServicehost + "/data")
		if err != nil {
			log.Println("Error calling dataservice:", err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "Error calling dataservice")
			return
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)

		fmt.Fprintf(w, "Hello World! Dataservice says: %s", string(body))
	})

	fmt.Printf("Starting hello-world on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
