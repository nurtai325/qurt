package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	secret := os.Getenv("GIT_SECRET")
	if secret == "" {
		panic("secret is empty. set GIT_SECRET env variable")
	}
	http.HandleFunc("POST /githook/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hello")
		fmt.Println(r.Header)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(body))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
