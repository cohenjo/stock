package config

type config struct {
	Port int
	Name string
}

// C is global configuration
var C = newConfiguration()

func newConfiguration() *config {
	return &config{}
}
