package main

import (
	"log"
	"net/http"
)

func main() {
	file := http.FileServer(http.Dir("/mnt/chromeos/MyFiles/Downloads"))
	http.Handle("/file/", http.StripPrefix("/file/", file))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println(err)
	}
}
