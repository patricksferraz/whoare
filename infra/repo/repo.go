package repo

import (
	"context"
	"fmt"

	"github.com/patricksferraz/whoare/domain/entity"
	"github.com/patricksferraz/whoare/infra/db"
)

type Repository struct {
	Orm *db.DbOrm
}

func NewRepository(orm *db.DbOrm) *Repository {
	return &Repository{
		Orm: orm,
	}
}

func (r *Repository) CreateEmployee(ctx context.Context, employee *entity.Employee) error {
	err := r.Orm.Db.Create(employee).Error
	return err
}

func (r *Repository) FindEmployee(ctx context.Context, employeeID *string) (*entity.Employee, error) {
	var e entity.Employee

	r.Orm.Db.First(&e, "id = ?", *employeeID)

	if e.ID == "" {
		return nil, fmt.Errorf("no employee found")
	}

	return &e, nil
}

func (r *Repository) SaveEmployee(ctx context.Context, employee *entity.Employee) error {
	err := r.Orm.Db.Save(employee).Error
	return err
}

func (r *Repository) FindEmployeesByName(ctx context.Context, filter *entity.FilterEmployee) ([]*entity.Employee, *int64, error) {
	var e []*entity.Employee

	q := r.Orm.Db.
		Where("name LIKE ?", "%"+filter.Q+"%")
	count := q.Find(&e).RowsAffected

	q = q.Order("name " + filter.SortBy).
		Limit(filter.PageSize).
		Offset(filter.PageSize * (filter.Page - 1))
	err := q.Find(&e).Error
	if err != nil {
		return nil, nil, err
	}

	return e, &count, nil
}

func (r *Repository) FindEmployeesBySkill(ctx context.Context, filter *entity.FilterEmployee) ([]*entity.Employee, *int64, error) {
	var e []*entity.Employee

	q := r.Orm.Db.
		Preload("Skills", "name = ?", "%"+filter.Q+"%")
	count := q.Find(&e).RowsAffected

	q = q.Order("name " + filter.SortBy).
		Limit(filter.PageSize).
		Offset(filter.PageSize * (filter.Page - 1))
	err := q.Find(&e).Error
	if err != nil {
		return nil, nil, err
	}

	return e, &count, nil
}

func (r *Repository) CreateSkill(ctx context.Context, skill *entity.Skill) error {
	err := r.Orm.Db.Create(skill).Error
	return err
}

func (r *Repository) FindSkill(ctx context.Context, skillID *string) (*entity.Skill, error) {
	var e entity.Skill

	r.Orm.Db.First(&e, "id = ?", *skillID)

	if e.ID == "" {
		return nil, fmt.Errorf("no skill found")
	}

	return &e, nil
}

func (r *Repository) SaveSkill(ctx context.Context, skill *entity.Skill) error {
	err := r.Orm.Db.Save(skill).Error
	return err
}
