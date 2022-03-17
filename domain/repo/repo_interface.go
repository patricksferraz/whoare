package repo

import (
	"context"

	"github.com/patricksferraz/whoare/domain/entity"
)

type RepoInterface interface {
	CreateEmployee(ctx context.Context, employee *entity.Employee) error
	FindEmployee(ctx context.Context, employeeID *string) (*entity.Employee, error)
	SaveEmployee(ctx context.Context, employee *entity.Employee) error
	SearchEmployees(ctx context.Context, query *string) ([]*entity.Employee, error)

	CreateSkill(ctx context.Context, skill *entity.Skill) error
	FindSkill(ctx context.Context, skillID *string) (*entity.Skill, error)
	SaveSkill(ctx context.Context, skill *entity.Skill) error
	FindSkillByName(ctx context.Context, name *string) (*entity.Skill, error)

	AddEmployeeSkill(ctx context.Context, employeeSkill *entity.EmployeesSkill) error

	// NOTE: temporary
	DeleteEmployeeSkills(ctx context.Context, employeeID *string) error
}
