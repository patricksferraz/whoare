package config

type Config struct {
	Db struct {
		Debug   *bool   `env:"DB_DEBUG,default=false"`
		Migrate *bool   `env:"DB_MIGRATE,default=true"`
		DsnType *string `env:"DSN_TYPE,required=true"`
		Dsn     *string `env:"DSN,required=true"`
	}
}
