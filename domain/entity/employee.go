package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
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
	TerminationDate time.Time         `json:"termination_date" gorm:"column:termination_date" valid:"-"`
	IsActive        bool              `json:"is_active" gorm:"column:is_active;not null" valid:"-"`
	Skills          []*Skill          `json:"-" gorm:"many2many:employees_skills" valid:"-"`
	EmployeesSkills []*EmployeesSkill `json:"-" gorm:"foreignKey:EmployeeID" valid:"-"`
}

func NewEmployee(name, position, email, password, presentation string, hireDate time.Time) (*Employee, error) {
	e := Employee{
		Name:         name,
		Position:     position,
		Email:        email,
		Presentation: presentation,
		HireDate:     hireDate,
		IsActive:     true,
	}
	e.ID = uuid.NewV4().String()
	e.CreatedAt = time.Now()

	if err := e.HashedPassword(password); err != nil {
		return nil, err
	}

	if err := e.isValid(); err != nil {
		return nil, err
	}

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

func (e *Employee) Activate(hireDate time.Time) error {
	e.IsActive = true
	e.HireDate = hireDate
	e.TerminationDate = time.Time{}
	e.UpdatedAt = time.Now()
	err := e.isValid()
	return err
}

func (e *Employee) Deactivate(terminationDate time.Time) error {
	e.IsActive = false
	e.TerminationDate = terminationDate
	e.UpdatedAt = time.Now()
	err := e.isValid()
	return err
}

func (e *Employee) HashedPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	e.Password = string(hash)

	return nil
}

func (e *Employee) CompareHashAndPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(e.Password), []byte(password))
	return err
}
