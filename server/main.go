package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/csoban373/context_example/server/handler"
)

func main() {
	http.HandleFunc("/", handler.Decorator(example))
	fmt.Println("starting server on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func example(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Println("received request")

	id, err := handler.GetValue(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("value stored %v\n", id)

	defer fmt.Println()
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("succesful request")
		fmt.Fprintln(w, "hello world")
	case <-ctx.Done():
		err := ctx.Err().Error()
		fmt.Println(err)
		http.Error(w, err, http.StatusInternalServerError)
	}
}
