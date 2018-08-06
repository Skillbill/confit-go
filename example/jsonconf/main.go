//
// load a configuration in json format with confit
//
package main

import (
	"github.com/Skillbill/confit-go"

	"encoding/json"
	"log"
	"os"
)

// Configurations  we'll need:
type Config struct {
	DB   DatabaseCfg
	Auth AuthCfg
}

type DatabaseCfg struct {
	Address  string
	Port     int
	Name     string
	User     string
	Password string
}

type AuthCfg struct {
	PrivateKeyId string `json:"private_key_id"`
	PrivateKey   string `json:"private_key"`
}

func loadConfig(buf *Config) error {
	id := os.Getenv("CONFIT_REPOID")
	secret := os.Getenv("CONFIT_REPOSECRET")
	ref := os.Getenv("CONFIT_REPOREF") // we can use specific configurations using git tags
	rsc := os.Getenv("CONFIT_RESOURCE")
	c := confit.Client{RepoId: id, Secret: secret, Ref: ref}
	p, err := c.LoadByPath(rsc)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(p, buf); err != nil {
		return err
	}
	return nil
}

func main() {
	cfg := Config{}
	err := loadConfig(&cfg)
	if err != nil {
		log.Fatalln("could not load configuration:", err)
	}
	// get the work done...
}
