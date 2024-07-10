package handler

import (
	"encoding/json"
	"github.com/amirhossein2831/Vault_Provider/src/model"
	"github.com/amirhossein2831/Vault_Provider/src/pkg"
	"log"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := model.Users()

	marshal, err := json.Marshal(users)
	if err != nil {
		log.Fatal("failed to marshal users")
	}
	_, err = w.Write(marshal)
	if err != nil {
		log.Fatal("failed to write response")
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	userName, err := pkg.GetVault().GetVar("user_name")
	if err != nil {
		log.Fatal("failed to get user_name")
	}

	password, err := pkg.GetVault().GetVar("password")
	if err != nil {
		log.Fatal("failed to get password")
	}

	user := &model.User{
		Name:           userName,
		SecretPassword: password,
	}

	marshal, err := json.Marshal(user)
	if err != nil {
		log.Fatal("failed to marshal user")
	}
	_, err = w.Write(marshal)
	if err != nil {
		log.Fatal("failed to write response")
	}
}
