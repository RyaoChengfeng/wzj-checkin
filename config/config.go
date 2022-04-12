package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
)

func init() {
	configFile := "config.yml"

	if v, ok := os.LookupEnv("ENV"); ok {
		configFile = v + ".yml"
	}

	data, err := ioutil.ReadFile(fmt.Sprintf("env/%s", configFile))

	if err != nil {
		panic(err)
		return
	}

	config := &Config{}

	err = yaml.Unmarshal(data, config)

	if err != nil {
		log.Print("Unmarshal config error!")
		panic(err)
		return
	}

	C = config

	log.Print("Config " + configFile + " loaded.")
	if C.Debug {
		log.Printf("%+v\n", C)
	}
}
