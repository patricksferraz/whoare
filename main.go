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
	"gorm.io/gorm/logger"
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

	l := logger.Error
	if *conf.Db.Debug {
		l = logger.Info
	}

	orm, err := db.NewDbOrm(*conf.Db.Dsn, l)
	if err != nil {
		log.Fatal(err)
	}

	if err = orm.Migrate(); err != nil {
		log.Fatal(err)
	}
	log.Printf("migration did run successfully")

	rest.StartFront(orm)
}
