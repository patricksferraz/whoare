package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Skill struct {
	Base      `json:",inline" valid:"-"`
	Name      string      `json:"name" gorm:"column:name;type:varchar(255)" valid:"-"`
	Employees []*Employee `json:"-" gorm:"many2many:employees_skils" valid:"-"`
}

func NewSkill(name string) (*Skill, error) {
	e := Skill{
		Name: name,
	}
	e.ID = uuid.NewV4().String()
	e.CreatedAt = time.Now()

	if err := e.isValid(); err != nil {
		return nil, err
	}

	return &e, nil
}

func (e *Skill) isValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}
