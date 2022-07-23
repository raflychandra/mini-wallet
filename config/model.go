package config

var (
	env = map[string]string{
		"local": "config.local.yml",
		"prod":  "config.prod.yml",
		"test":  "config.test.yml",
	}
)

type Config struct {
	Application Application `yaml:"app"`
	Wallet      Wallet      `yaml:"wallet"`
}

type Application struct {
	Env  string `yaml:"env"`
	Name string `yaml:"name"`
	Port string `yaml:"port"`
	Host string `yaml:"host"`
}

type Wallet struct {
	KeySignature string `yaml:"key-signature"`
}
