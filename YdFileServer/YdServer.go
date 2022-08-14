package main

import (
	"fmt"
	"net/http"
	"log"
)

func main() {
	fmt.Println("YdFileServer\nbuild by ljlVink 20220814\napp start on port 8000")
	http.HandleFunc("/", index)
    log.Fatal(http.ListenAndServe(":8000", http.FileServer(http.Dir("/"))))
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/index")
	fmt.Fprintf(w, "{\"data\":\"Network test\"}")
}
