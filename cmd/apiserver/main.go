package main

import (
  "log"

  "github.com/f1rerabbit/golang_1/internal/app/apiserver"
)

// Путь к фаилу задать в качестве флага бинарника
var (
  configPath string
)

func init() {
  flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main()  {
  flag.Parse()


  config := apiserver.NewConfig()
  _, err := toml.DecodeFile(configPath, config)
  if err != nil {
    log.Fatal(err)
  }



  s := apiserver.New()
  if err := s.Start(); err != nil {
    log.Fatal(err)
  }
}
