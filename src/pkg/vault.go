package pkg

import (
	"fmt"
	"github.com/hashicorp/vault/api"
)

type Vault struct {
	Config    *api.Config
	Client    *api.Client
	PushVar   map[string]interface{}
	PulledVar map[string]string
}

func NewVault(config *api.Config) *Vault {
	return &Vault{Config: config,
		PushVar:   map[string]interface{}{},
		PulledVar: map[string]string{},
	}
}

func (v *Vault) Connect() error {
	client, err := api.NewClient(v.Config)
	if err != nil {
		return err
	}

	v.Client = client

	return nil
}

func (v *Vault) SetToken(token string) {
	v.Client.SetToken(token)
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
		return nil
	}

	return fmt.Errorf("no secret found in path: %s", path)
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
