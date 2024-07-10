package main

import (
	"github.com/amirhossein2831/Vault_Provider/src/handler"
	"github.com/amirhossein2831/Vault_Provider/src/pkg"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-chi/chi/v5"
	"github.com/hashicorp/vault/api"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	// Load env var
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Init  env var
	vaultAdd := os.Getenv("VAULT_ADDRESS")
	vaultToken := os.Getenv("VAULT_TOKEN")
	vaultPath := os.Getenv("VAULT_PATH")
	appPort := os.Getenv("APP_PORT")

	// Initial the vault
	pkg.NewVault(&api.Config{
		Address: vaultAdd,
	})

	// Read the variable
	err = pkg.GetVault().SetToken(vaultToken).ReadVault(vaultPath)
	if err != nil {
		log.Fatal(err)
	}

	// Initial the router
	router := chi.NewRouter()

	// Initial the route
	router.Get("/users", handler.GetUsers)
	router.Get("/users/secret", handler.GetUser)

	// Initial the server
	go func() {
		err := http.ListenAndServe(spew.Sprintf(":%v", appPort), router)
		if err != nil {
			log.Fatal("failed to start server")
		}
	}()
	println("server started successfully on port " + appPort)
	select {}
}
