package main

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/Netflix/go-env"
	"github.com/joho/godotenv"
	rest "github.com/patricksferraz/whoare/app/front"
	"github.com/patricksferraz/whoare/config"
	"github.com/patricksferraz/whoare/infra/db"
)

func init() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	if os.Getenv("ENV") == "dev" {
		err := godotenv.Load(basepath + "/.env")
		if err != nil {
			log.Printf("Error loading .env files")
		}
	}
}

func main() {
	var conf config.Config

	_, err := env.UnmarshalFromEnviron(&conf)
	if err != nil {
		log.Fatal(err)
	}

	orm, err := db.NewDbOrm(*conf.Db.DsnType, *conf.Db.Dsn)
	if err != nil {
		log.Fatal(err)
	}

	if *conf.Db.Debug {
		orm.Debug(true)
	}

	if *conf.Db.Migrate {
		orm.Migrate()
	}
	defer orm.Db.Close()

	rest.StartFront(orm)
}
