package entity

import (
	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type EmployeesSkill struct {
	EmployeeID string `gorm:"column:employee_id;type:uuid;not null;unique_index:unique_employee_skill;primaryKey" valid:"uuid"`
	SkillID    string `gorm:"column:skill_id;type:uuid;not null;unique_index:unique_employee_skill;primaryKey" valid:"uuid"`
	// TODO: add seniority type
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
