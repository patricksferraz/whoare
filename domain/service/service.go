package service

import (
	"context"

	"github.com/patricksferraz/whoare/domain/entity"
	"github.com/patricksferraz/whoare/domain/repo"
)

type Service struct {
	Repo repo.RepoInterface
}

func NewService(repo repo.RepoInterface) *Service {
	return &Service{
		Repo: repo,
	}
}

func (s *Service) CreateEmployee(ctx context.Context, name, position, email, password string) (*string, error) {
	e, err := entity.NewEmployee(name, position, email, password)
	if err != nil {
		return nil, err
	}

	if err = s.Repo.CreateEmployee(ctx, e); err != nil {
		return nil, err
	}

	return &e.ID, nil
}

func (s *Service) FindEmployee(ctx context.Context, employeeID *string) (*entity.Employee, error) {
	e, err := s.Repo.FindEmployee(ctx, employeeID)
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (s *Service) FindEmployeesByName(ctx context.Context, filter *entity.FilterEmployee) ([]*entity.Employee, *int64, error) {
	e, count, err := s.Repo.FindEmployeesByName(ctx, filter)
	if err != nil {
		return nil, nil, err
	}

	return e, count, nil
}

func (s *Service) FindEmployeesBySkill(ctx context.Context, filter *entity.FilterEmployee) ([]*entity.Employee, *int64, error) {
	e, count, err := s.Repo.FindEmployeesBySkill(ctx, filter)
	if err != nil {
		return nil, nil, err
	}

	return e, count, nil
}

func (s *Service) CreateSkill(ctx context.Context, name string) (*string, error) {
	e, err := entity.NewSkill(name)
	if err != nil {
		return nil, err
	}

	if err = s.Repo.CreateSkill(ctx, e); err != nil {
		return nil, err
	}

	return &e.ID, nil
}

func (s *Service) FindSkill(ctx context.Context, skillID *string) (*entity.Skill, error) {
	e, err := s.Repo.FindSkill(ctx, skillID)
	if err != nil {
		return nil, err
	}

	return e, nil
}
