package apiserver

// Apiserver

type APIServer struct {
  config * Config
}

// New

func New(config * Config) *APIServer  {
  return &APIServer
  config: config,
}

func (s *APIServer) Start() error  {
  return nil
}
