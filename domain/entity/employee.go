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
	Base     `json:",inline" valid:"-"`
	Name     string    `json:"name" gorm:"column:name;not null" valid:"required"`
	Email    string    `json:"email" gorm:"column:email;not null" valid:"email"`
	Password string    `json:"password" gorm:"column:password;not null" valid:"required"`
	Position string    `json:"position" gorm:"column:position;not null" valid:"required"`
	HireDate time.Time `json:"hire_date" gorm:"column:hire_date" valid:"-"`
	Skills   []*Skill  `json:"-" gorm:"many2many:employees_skills" valid:"-"`
}

func NewEmployee(name, position, email, password string) (*Employee, error) {
	e := Employee{
		Name:     name,
		Position: position,
		Email:    email,
		Password: utils.HashedPassword(password),
	}
	e.ID = uuid.NewV4().String()
	e.CreatedAt = time.Now()

	if err := e.isValid(); err != nil {
		return nil, err
	}

	return &e, nil
}

func (e *Employee) isValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}

type FilterEmployee struct {
	Q        string `json:"q" valid:"-"`
	SearchBy string `json:"search_by" valid:"-"`
	SortBy   string `json:"sort_by" valid:"-"`
	PageSize int64  `json:"page_size" valid:"-"`
	Page     int64  `json:"page" valid:"-"`
}

func NewFilterEmployee(q, searchBy, sortBy string, pageSize, page int64) *FilterEmployee {
	if page == 0 {
		page = 1
	}

	if pageSize == 0 {
		pageSize = 10
	}

	if sortBy == "" {
		sortBy = "asc"
	}

	e := FilterEmployee{
		Q:        q,
		SearchBy: searchBy,
		SortBy:   sortBy,
		PageSize: pageSize,
		Page:     page,
	}

	return &e
}
