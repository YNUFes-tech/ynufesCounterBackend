package setting

import (
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

var (
	setting *Setting = nil
)

type (
	Setting struct {
		Infrastructure Infrastructure `yaml:"infrastructure"`
		Service        Service        `yaml:"service"`
		ThirdParty     ThirdParty     `yaml:"third_party"`
	}
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Failed to load .env file... %v", err)
	}
	envLocation := os.Getenv("ENV_LOCATION")
	log.Println("ENV_LOCATION: " + envLocation)
	reader, err := os.Open(envLocation)
	if err != nil {
		dir, _ := os.Getwd()
		log.Fatalln(nil, "failed to open setting file: %v, %v\n", dir, err)
	}
	decoder := yaml.NewDecoder(reader)
	setting = &Setting{}
	err = decoder.Decode(setting)
	if err != nil {
		panic(err)
	}
}

func Get() Setting {
	if setting == nil {
		panic("setting is nil")
	}
	return *setting
}
