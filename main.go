package main

import (
	"github.com/amirhossein2831/Vault_Provider/src/handler"
	"github.com/amirhossein2831/Vault_Provider/src/pkg"
	"github.com/go-chi/chi/v5"
	"github.com/hashicorp/vault/api"
	"log"
	"net/http"
)

func main() {
	// Initial the vault
	pkg.NewVault(&api.Config{
		Address: "http://127.0.0.1:8200",
	})

	// Read the variable
	err := pkg.GetVault().SetToken("root").ReadVault("secret/data/app")
	if err != nil {
		return
	}

	// Initial the router
	router := chi.NewRouter()

	// Initial the route
	router.Get("/users", handler.GetUsers)
	router.Get("/users/secret", handler.GetUser)

	// Initial the server
	go func() {
		err := http.ListenAndServe(":2000", router)
		if err != nil {
			log.Fatal("failed to start server")
		}
	}()
	println("server started successfully on port 2000")
	select {}
}
