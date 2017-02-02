package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/chilts/sandpeople"
	"github.com/gomiddleware/mux"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("homeHandler(): entry")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Home\n"))
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("myHandler(): entry")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("You are logged in.\n"))
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("userHandler(): entry")
	w.WriteHeader(http.StatusOK)

	user := sandpeople.GetUser(r)
	if user == nil {
		w.Write([]byte("No user detected.\n"))
		return
	}

	w.Write([]byte(fmt.Sprintf("User=%#v\n", user)))
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("adminHandler(): entry")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("You are admin.\n"))
}

func main() {
	m := mux.New()

	m.Get("/", homeHandler)
	m.Get("/my/", sandpeople.RequireUser("/"), myHandler)
	m.Get("/user/", sandpeople.MakeUser, userHandler)
	m.Get("/admin/", sandpeople.HasPerm("admin", "/"), adminHandler)

	log.Println("Starting server")
	log.Fatal(http.ListenAndServe(":8080", m))
}
