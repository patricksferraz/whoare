package config

type Config struct {
	Db struct {
		Debug *bool   `env:"DB_DEBUG,default=false"`
		Dsn   *string `env:"DSN,required=true"`
	}
}
