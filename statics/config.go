package statics

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Config struct {
	Database Database
	Redis Redis
}

type Database struct {
	Ip string
	Port string
	Name string
	Username string
	Password string
}

type Redis struct {
	Ip string
	Port string
	Password string
	DB int
}

var config Config

func init()  {
	path := "config.json"
	for _, val := range os.Args {
		if strings.Contains(val, "--conifg=") {
			path = strings.Replace(val, "--conifg=", "", -1)
		}
	}
	f, err := os.OpenFile(path, os.O_RDONLY, 0600)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		log.Fatal(err)
	}
}

func GetConfig() Config {
	return config
}