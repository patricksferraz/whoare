package service

import (
	"context"
	"errors"
	"time"

	"github.com/patricksferraz/whoare/domain/entity"
	"github.com/patricksferraz/whoare/domain/repo"
	"github.com/patricksferraz/whoare/utils"
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

func (s *Service) UpdateEmployee(ctx context.Context, employeeID, name, position, email, password, presentation string, hireDate time.Time) error {
	e, err := s.Repo.FindEmployee(ctx, &employeeID)
	if err != nil {
		return err
	}

	if err = e.SetName(name); err != nil {
		return err
	}

	if err = e.SetEmail(email); err != nil {
		return err
	}

	if err = e.SetPosition(position); err != nil {
		return err
	}

	if err = e.SetPresentation(presentation); err != nil {
		return err
	}

	if err = e.SetHireDate(hireDate); err != nil {
		return err
	}

	if err = utils.CompareHashAndPassword(e.Password, password); err != nil {
		return errors.New("invalid password")
	}

	if err = s.Repo.SaveEmployee(ctx, e); err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteEmployee(ctx context.Context, employeeID, employeePassword *string) error {
	e, err := s.Repo.FindEmployee(ctx, employeeID)
	if err != nil {
		return err
	}

	if err = utils.CompareHashAndPassword(e.Password, *employeePassword); err != nil {
		return errors.New("invalid password")
	}

	err = s.Repo.DeleteEmployee(ctx, e)
	if err != nil {
		return err
	}

	return nil
}
