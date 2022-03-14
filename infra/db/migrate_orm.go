package db

import (
	"fmt"
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"gorm.io/gorm"
)

type MigrateOrm struct {
	Db *gorm.DB
	m  *gormigrate.Gormigrate
}

func NewMigrateOrm(db *gorm.DB) *MigrateOrm {
	m := MigrateOrm{
		Db: db,
	}
	m.load()
	return &m
}

func (m *MigrateOrm) load() {
	m.m = gormigrate.New(m.Db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "202203140038",
			Migrate: func(db *gorm.DB) error {
				type Base struct {
					ID        string    `gorm:"type:uuid;primaryKey"`
					CreatedAt time.Time `gorm:"column:created_at;autoUpdateTime"`
					UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime"`
				}
				type Employee struct {
					Base         `json:",inline"`
					Name         string    `gorm:"column:name;not null"`
					Email        string    `gorm:"column:email;not null"`
					Password     string    `gorm:"column:password;not null"`
					Position     string    `gorm:"column:position;not null"`
					Presentation string    `gorm:"column:presentation"`
					HireDate     time.Time `gorm:"column:hire_date"`
				}
				type XP int
				type EmployeesSkill struct {
					EmployeeID string `gorm:"column:employee_id;type:uuid;not null;unique_index:unique_employee_skill;primaryKey"`
					SkillID    string `gorm:"column:skill_id;type:uuid;not null;unique_index:unique_employee_skill;primaryKey"`
					XP         XP     `gorm:"column:xp;not null"`
				}
				type Skill struct {
					Base `json:",inline"`
					Name string `json:"name" gorm:"column:name;type:varchar(255)"`
				}

				return db.AutoMigrate(&Employee{}, &Skill{}, &EmployeesSkill{})
			},
			Rollback: func(db *gorm.DB) error {
				return db.Migrator().DropTable("employees", "skills", "employees_skills")
			},
		},
	})
}

func (m *MigrateOrm) Migrate() error {
	if err := m.m.Migrate(); err != nil {
		return fmt.Errorf("could not migrate: %v", err)
	}
	return nil
}
