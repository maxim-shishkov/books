package config

type postgresConfig struct {
	Host     string
	Port     int64
	User     string
	Password string
	DbName   string
}

var Postgres = &postgresConfig{}

func (pc *postgresConfig) Init() {
}
