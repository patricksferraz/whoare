package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/patricksferraz/whoare/utils"
	uuid "github.com/satori/go.uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Employee struct {
	Base            `json:",inline" valid:"-"`
	Name            string            `json:"name" gorm:"column:name;not null" valid:"required"`
	Email           string            `json:"email" gorm:"column:email;not null" valid:"email"`
	Password        string            `json:"-" gorm:"column:password;not null" valid:"required"`
	Position        string            `json:"position" gorm:"column:position;not null" valid:"required"`
	Presentation    string            `json:"presentation" gorm:"column:presentation" valid:"required"`
	HireDate        time.Time         `json:"hire_date" gorm:"column:hire_date" valid:"required"`
	Skills          []*Skill          `json:"-" gorm:"many2many:employees_skills" valid:"-"`
	EmployeesSkills []*EmployeesSkill `json:"-" gorm:"foreignKey:EmployeeID" valid:"-"`
}

func NewEmployee(name, position, email, password, presentation string, hireDate time.Time) (*Employee, error) {
	e := Employee{
		Name:         name,
		Position:     position,
		Email:        email,
		Password:     password,
		Presentation: presentation,
		HireDate:     hireDate,
	}
	e.ID = uuid.NewV4().String()
	e.CreatedAt = time.Now()

	if err := e.isValid(); err != nil {
		return nil, err
	}

	e.Password = utils.HashedPassword(password)

	return &e, nil
}

func (e *Employee) isValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}

func (e *Employee) SetName(name string) error {
	e.Name = name
	e.UpdatedAt = time.Now()
	err := e.isValid()
	return err
}

func (e *Employee) SetEmail(email string) error {
	e.Email = email
	e.UpdatedAt = time.Now()
	err := e.isValid()
	return err
}

func (e *Employee) SetPosition(position string) error {
	e.Position = position
	e.UpdatedAt = time.Now()
	err := e.isValid()
	return err
}

func (e *Employee) SetPresentation(presentation string) error {
	e.Presentation = presentation
	e.UpdatedAt = time.Now()
	err := e.isValid()
	return err
}

func (e *Employee) SetHireDate(hireDate time.Time) error {
	e.HireDate = hireDate
	e.UpdatedAt = time.Now()
	err := e.isValid()
	return err
}
