package config

import (
    "log"

    "github.com/BurntSushi/toml"
)

type Config struct {
  GraylogAddr string
}

var ServiceConfig Config

// Reads info from config file
func ReadConfig(conf string) {
  var cnf Config
  if _, err := toml.DecodeFile(conf, &cnf); err != nil {
	  log.Println(err)
  }

  ServiceConfig = cnf
}

func Get() Config {
  return ServiceConfig
}
