package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server Builder")

	filePath := http.FileServer(http.Dir("./static"))

	http.Handle("/", filePath)

	http.HandleFunc("/form", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		http.ServeFile(w, r, "./static/form.html")
		name := r.FormValue("name")
		email := r.FormValue("email")
		password := r.FormValue("password")
		cpass := r.FormValue("cpass")
		fmt.Println(name, email, password, cpass)
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
		http.ServeFile(w, r, "./static/index.html")
	})

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		// now send home.html
		http.ServeFile(w, r, "./static/home.html")
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
