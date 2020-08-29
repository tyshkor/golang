package main

import (
	"crypto/sha256"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var m map[string]int

func handlerauth(w http.ResponseWriter, r *http.Request) {
	q := r.Method

	if q == "POST" {
		str2 := r.FormValue("email")
		str3 := r.FormValue("pword")

		_, ok := m[str2]
		if ok {
			h := sha256.New()
			pw, _ := h.Write([]byte(str3))
			// fmt.Println(pw)
			// fmt.Print("m[str2] = ")
			// fmt.Println(m[str2])
			if m[str2] == pw {
				expirationTime := time.Now().Add(time.Second * 20)
				tokenString := str2 + " " + strconv.Itoa(pw)
				// var arr []string
				// arr = append(arr, str2)
				// arr = append(arr, strconv.Itoa(pw))

				http.SetCookie(w, &http.Cookie{
					Name:    "token",
					Value:   tokenString,
					Expires: expirationTime,
				})
				// fmt.Println(arr)
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
	q := r.Method

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

func handlerhomepage(w http.ResponseWriter, r *http.Request) {
	q := r.Method

	if q == "GET" {
		k, err := r.Cookie("token")
		if err != nil {
			fmt.Fprintf(w, "You don't have a token.")
		} else {
			// fmt.Println(k.Unparsed)
			name := strings.Split(k.Value, " ")[0]
			pw := strings.Split(k.Value, " ")[1]
			q, _ := strconv.Atoi(pw)
			if p, ok := m[name]; ok && p == q {
				fmt.Fprintf(w, "Your token is valid!")
			} else {
				fmt.Fprintf(w, "You don't have a token.")
			}
			// fmt.Println(m)
		}
	}

}

func main() {
	m = make(map[string]int)
	http.HandleFunc("/auth", handlerauth)
	http.HandleFunc("/register", handlerregister)
	http.HandleFunc("/homepage", handlerhomepage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
