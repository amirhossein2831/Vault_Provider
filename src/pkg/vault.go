package pkg

import (
	"fmt"
	"github.com/hashicorp/vault/api"
	"log"
	"sync"
)

var (
	instance *Vault
	once     sync.Once
)

type Vault struct {
	Config    *api.Config
	Client    *api.Client
	PushVar   map[string]interface{}
	PulledVar map[string]string
}

// NewVault NewVault: create a singleton object from Vault
func NewVault(config *api.Config) *Vault {
	once.Do(func() {
		instance = &Vault{
			Config:    config,
			PushVar:   map[string]interface{}{},
			PulledVar: map[string]string{},
		}

		err := instance.Connect()
		if err != nil {
			log.Fatal("failed to connect to vault")
		}
		log.Println("successfully connected to vault")
	})
	return instance
}

// GetVault return the singleton object from Vault
func GetVault() *Vault {
	return instance
}

func (v *Vault) Connect() error {
	client, err := api.NewClient(v.Config)
	if err != nil {
		return err
	}

	v.Client = client

	return nil
}

func (v *Vault) WriteVault(path string) error {
	_, err := v.Client.Logical().Write(path, v.PushVar)
	if err != nil {
		return err
	}
	return nil
}

func (v *Vault) ReadVault(path string) error {
	secret, err := v.Client.Logical().Read(path)
	if err != nil {
		return err
	}

	if secret != nil {
		if data, ok := secret.Data["data"].(map[string]interface{}); ok {
			for key, value := range data {
				v.PulledVar[key] = value.(string)
			}
		}
		log.Println("successfully read vault data")
		return nil
	}

	return fmt.Errorf("no secret found in path: %s", path)
}

func (v *Vault) SetToken(token string) *Vault {
	v.Client.SetToken(token)
	return v
}

func (v *Vault) SetVar(key string, value string) {
	v.PushVar[key] = value
}

func (v *Vault) GetVar(key string) (string, error) {
	value, ok := v.PulledVar[key]
	if !ok {
		return "", fmt.Errorf("no secret found in path: %s", key)
	}
	return value, nil
}
