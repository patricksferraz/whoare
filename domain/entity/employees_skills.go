package entity

import (
	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
	govalidator.TagMap["xp"] = govalidator.Validator(func(str string) bool {
		res := str == JUNIOR.String()
		res = res || str == FULL.String()
		res = res || str == SENIOR.String()
		res = res || str == SPECIALIST.String()
		return res
	})
}

type EmployeesSkill struct {
	EmployeeID string `gorm:"column:employee_id;type:uuid;not null;unique_index:unique_employee_skill;primaryKey" valid:"uuid"`
	SkillID    string `gorm:"column:skill_id;type:uuid;not null;unique_index:unique_employee_skill;primaryKey" valid:"uuid"`
	XP         XP     `json:"xp" gorm:"column:xp;not null" valid:"xp,optional"`
}

func NewEmployeesSkill(employeeID, skillID string) (*EmployeesSkill, error) {
	e := EmployeesSkill{
		EmployeeID: employeeID,
		SkillID:    skillID,
	}

	if err := e.isValid(); err != nil {
		return nil, err
	}

	return &e, nil
}

func (e *EmployeesSkill) isValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}

func (e *EmployeesSkill) AddXp(xp XP) error {
	e.XP = xp
	err := e.isValid()
	return err
}

type XP int

const (
	JUNIOR XP = iota + 1
	FULL
	SENIOR
	SPECIALIST
)

func (t XP) String() string {
	switch t {
	case JUNIOR:
		return "JUNIOR"
	case FULL:
		return "FULL"
	case SENIOR:
		return "SENIOR"
	case SPECIALIST:
		return "SPECIALIST"
	}
	return ""
}
