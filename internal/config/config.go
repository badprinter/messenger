package config

type BaseConfig struct {
	Net  NetCofnig
	User UserConfig
}

type NetCofnig struct {
	Host string
	Port string
}

type UserConfig struct {
	Name string
}
