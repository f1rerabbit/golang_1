package apiserver

type Config struct {
  BindAddr string `toml:"bing_addr"`
}
//NewConfig
func NewConfig() *Config  {
    return &Config{
      BindAddr: `:8080`,
    }
}
