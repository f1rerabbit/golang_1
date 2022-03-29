package apiserver

import (
  "github.com/sirupsen/logrus"
)



// Apiserver

type APIServer struct {
  config *Config
// Поле в качестве апи сервера
  logger *logrus.Logger
}

// New

func New(config *Config) *APIServer  {
  return &APIServer {
  config: config,
  // инициализация аписервера
  logger: logrus.New(),
  }
}

// Start
func (s *APIServer) Start() error  {
  if err := s.configureLogger(); err != nil {
    return err
  }

  s.logger.info("starting api server")

  return nil
}

// Функция которая конфиргурирует наш логгер

func (s *APIServer) confureLogger() error {
  level, err := logrus.ParseLeverl(s, config.Loglevel)
  if err != nil {
    return err
  }

  s.logger.SetLevel(level)

  return nil

}
