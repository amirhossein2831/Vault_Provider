package main

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

type User struct {
	Name           string `json:"name"`
	SecretPassword string `json:"password"`
	NationalId     string `json:"nationalId"`
}

func Users() []User {
	return []User{
		{Name: "amir", SecretPassword: "Password", NationalId: "12341234"},
		{Name: "ali", SecretPassword: "myPass", NationalId: "12121212"},
		{Name: "ali", SecretPassword: "varySecretPass", NationalId: "11111111"},
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	users := Users()

	marshal, err := json.Marshal(users)
	if err != nil {
		log.Fatal("failed to marshal users")
	}
	_, err = w.Write(marshal)
	if err != nil {
		log.Fatal("failed to write response")
	}
}
func main() {
	router := chi.NewRouter()
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("welcome"))
		if err != nil {
			log.Fatal("failed to write response")
		}
	})

	router.Get("/posts", GetUser)

	go func() {
		err := http.ListenAndServe(":2000", router)
		if err != nil {
			log.Fatal("failed to start server")
		}
	}()

	println("server started successfully on port 2000")
	select {}
}
