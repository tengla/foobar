package main

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
)

func addHandleFunc(path string, fn func(rw http.ResponseWriter, r *http.Request)) {
	http.HandleFunc(path, func(rw http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path)
		auth := r.Header.Get("Authorization")
		var validID = regexp.MustCompile(`^Bearer abc$`)
		if !validID.MatchString(auth) {
			http.Error(rw, "Unauthorized", http.StatusUnauthorized)
			return
		}
		fn(rw, r)
	})
}

func main() {
	addHandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		message := map[string]string{
			"message": "Hello World",
		}
		rw.Header().Add("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(message)
	})
	log.Println("Start server at port 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
