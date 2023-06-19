package main

import (
	"fmt"
	"net/http"
)

func main() {
	r := NewRouter()
	fmt.Println("Listen on port 8081")
	err := http.ListenAndServe(":8081", r.Mux)

	if err != nil {
		panic(err)
	}

}
