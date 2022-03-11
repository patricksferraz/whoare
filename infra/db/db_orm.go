package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/patricksferraz/whoare/domain/entity"
	_ "gorm.io/driver/sqlite"
)

type DbOrm struct {
	Db *gorm.DB
}

func NewDbOrm(dsnType, dsn string) (*DbOrm, error) {
	pg := &DbOrm{}

	err := pg.connect(dsnType, dsn)
	if err != nil {
		return nil, err
	}

	return pg, nil
}

func (p *DbOrm) connect(dsnType, dsn string) error {
	db, err := gorm.Open(dsnType, dsn)
	if err != nil {
		return fmt.Errorf("error connecting to database: %v", err)
	}

	p.Db = db

	return nil
}

func (p *DbOrm) Debug(enable bool) {
	p.Db.LogMode(enable)
}

func (p *DbOrm) Migrate() {
	p.Db.AutoMigrate(
		&entity.Employee{},
		&entity.Skill{},
		&entity.EmployeesSkill{},
	)
}
