package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.HandleFunc("/form", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%v\n%v\n%v\n", r.FormValue("name"), r.FormValue("hobby"), r.FormValue("age"))
	})

	fmt.Println("serving at 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
