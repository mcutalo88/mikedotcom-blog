package config

import (
  "os"
  "log"
  "github.com/spf13/viper"
)

var Vip = viper.New()

func Load() error {
  Vip.SetConfigType("json")

  if os.Getenv("GO_ENV") != "" {
    switch os.Getenv("GO_ENV") {
      case "dev":
        Vip.SetConfigName("dev")
      case "prod":
        Vip.SetConfigName("prod")
      default:
        Vip.SetConfigName("dev")
    }
  } else {
    Vip.SetConfigName("dev")
  }

  Vip.AddConfigPath("./config")
  Vip.AutomaticEnv()

  if err := Vip.ReadInConfig(); err == nil {
      log.Println("Using config file:", Vip.ConfigFileUsed())
  } else {
      return err
  }
  return nil
}
