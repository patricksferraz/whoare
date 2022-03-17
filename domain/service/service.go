package service

import (
	"context"
	"errors"
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
	var e *entity.Skill

	e, err := s.Repo.FindSkillByName(ctx, &name)
	if err != nil {
		e, err = entity.NewSkill(name)
		if err != nil {
			return nil, err
		}

		if err = s.Repo.CreateSkill(ctx, e); err != nil {
			return nil, err
		}
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

	if err = e.CompareHashAndPassword(password); err != nil {
		return errors.New("invalid password")
	}

	if err = s.Repo.SaveEmployee(ctx, e); err != nil {
		return err
	}

	return nil
}

func (s *Service) DeactivateEmployee(ctx context.Context, employeeID, employeePassword *string, terminationDate time.Time) error {
	e, err := s.Repo.FindEmployee(ctx, employeeID)
	if err != nil {
		return err
	}

	if err = e.CompareHashAndPassword(*employeePassword); err != nil {
		return errors.New("invalid password")
	}

	if err := e.Deactivate(terminationDate); err != nil {
		return err
	}

	err = s.Repo.SaveEmployee(ctx, e)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) AddEmployeeSkill(ctx context.Context, employeeID, skillID *string, xp int) error {
	e, err := s.Repo.FindEmployee(ctx, employeeID)
	if err != nil {
		return err
	}

	skill, err := s.Repo.FindSkill(ctx, skillID)
	if err != nil {
		return err
	}

	es, err := entity.NewEmployeesSkill(e, skill, entity.XP(xp))
	if err != nil {
		return err
	}

	if err = s.Repo.AddEmployeeSkill(ctx, es); err != nil {
		return err
	}

	return nil
}

func (s *Service) ActivateEmployee(ctx context.Context, employeeID, employeePassword *string, hireDate time.Time) error {
	e, err := s.Repo.FindEmployee(ctx, employeeID)
	if err != nil {
		return err
	}

	if err = e.CompareHashAndPassword(*employeePassword); err != nil {
		return errors.New("invalid password")
	}

	if err := e.Activate(hireDate); err != nil {
		return err
	}

	err = s.Repo.SaveEmployee(ctx, e)
	if err != nil {
		return err
	}

	return nil
}
