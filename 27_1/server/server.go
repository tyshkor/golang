package main

import (
	"crypto/sha256"
	"fmt"
	"log"
	"net/http"
)

var m map[string]int

func handlerauth(w http.ResponseWriter, r *http.Request) {
	// query := r.URL.Query()
	q := r.Method

	if q == "POST" {
		str2 := r.FormValue("email")
		str3 := r.FormValue("pword")

		// fmt.Println(str)

		_, ok := m[str2]
		if ok {
			h := sha256.New()
			pw, _ := h.Write([]byte(str3))
			fmt.Println(pw)
			fmt.Print("m[str2] = ")
			fmt.Println(m[str2])
			if m[str2] == pw {
				fmt.Fprintf(w, "You are signed in!")
			} else {
				fmt.Fprintf(w, "You are not signed in.")
			}
		} else {
			fmt.Fprintf(w, "You are not signed in.")
		}
		// fmt.Println(m)
	}

}

func handlerregister(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	q := r.Method

	fmt.Println(query)

	if q == "POST" {
		str2 := r.FormValue("email")
		str3 := r.FormValue("pword")
		// fmt.Println(str2)

		// fmt.Println(str3)
		h := sha256.New()
		q, _ := h.Write([]byte(str3))
		m[str2] = q
		// fmt.Println(m)
		// for b, a := range m {
		// 	fmt.Print(b + " = ")
		// 	fmt.Println(a)
		// }

	}

}

func main() {
	m = make(map[string]int)
	http.HandleFunc("/auth", handlerauth)
	http.HandleFunc("/register", handlerregister)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
