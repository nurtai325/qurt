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
	err := http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hello")
		fmt.Println(r.URL.Path)
		fmt.Println(r.Header)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(body))
	}))
	if err != nil {
		panic(err)
	}
}
