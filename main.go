package main

import (
	"fmt"
	"log"
	"net/http"
)

// /home route handler
func homeHandler(w http.ResponseWriter, r *http.Request) {
	
	if r.URL.Path != "/home" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}

	http.ServeFile(w, r, "./static/form.html")

}

// /form route handler
func formHandler(w http.ResponseWriter, r *http.Request) {
	
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	waifu := r.FormValue("waifu")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Waifu = %s\n", waifu)

}


// /asscii route handler
func assciiHandler(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w, r, "./static/asscii.html")
}


// main function
func main(){

	fileServer := http.FileServer(http.Dir("./static"))
	
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/home", homeHandler)
	http.HandleFunc("/asscii", assciiHandler)

	fmt.Println("server started at port 6969")
	if err := http.ListenAndServe(":6969", nil); err != nil {
		log.Fatal(err)
	}

}
