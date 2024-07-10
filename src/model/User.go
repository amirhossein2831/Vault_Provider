package model

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
