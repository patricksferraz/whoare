package service

import (
	"context"
	"time"

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

func (s *Service) CreateEmployee(ctx context.Context, name, position, email, password, presentation string, hireDate time.Time) (*string, error) {
	e, err := entity.NewEmployee(name, position, email, password, presentation, hireDate)
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

func (s *Service) SearchEmployees(ctx context.Context, query *string) ([]*entity.Employee, error) {
	e, err := s.Repo.SearchEmployees(ctx, query)
	if err != nil {
		return nil, err
	}

	return e, nil
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
