package repo

import (
	"context"

	"github.com/patricksferraz/whoare/domain/entity"
)

type RepoInterface interface {
	CreateEmployee(ctx context.Context, employee *entity.Employee) error
	FindEmployee(ctx context.Context, employeeID *string) (*entity.Employee, error)
	SaveEmployee(ctx context.Context, employee *entity.Employee) error
	FindEmployeesByName(ctx context.Context, filter *entity.FilterEmployee) ([]*entity.Employee, *int64, error)
	FindEmployeesBySkill(ctx context.Context, filter *entity.FilterEmployee) ([]*entity.Employee, *int64, error)

	CreateSkill(ctx context.Context, skill *entity.Skill) error
	FindSkill(ctx context.Context, skillID *string) (*entity.Skill, error)
	SaveSkill(ctx context.Context, skill *entity.Skill) error
}
