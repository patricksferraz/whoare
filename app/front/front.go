package front

import (
	"fmt"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/patricksferraz/whoare/domain/service"
)

type Front struct {
	Service *service.Service
	Session *session.Store
}

func NewFront(service *service.Service, session *session.Store) *Front {
	return &Front{
		Service: service,
		Session: session,
	}
}

func (a *Front) Index(c *fiber.Ctx) error {
	var req SearchRequest

	err := c.QueryParser(&req)
	if err != nil {
		return c.Render("views/errors/error", fiber.Map{
			"Status": fmt.Sprintf("%d - %s", fiber.StatusBadRequest, fiber.ErrBadRequest),
			"Error":  err.Error()},
		)
	}

	e, err := a.Service.SearchEmployees(c.Context(), &req.Q)
	if err != nil {
		return c.Render("views/errors/error", fiber.Map{
			"Status": fmt.Sprintf("%d - %s", fiber.StatusInternalServerError, fiber.ErrInternalServerError),
			"Error":  err.Error()},
		)
	}

	return c.Render("views/index", fiber.Map{
		"Q":         req.Q,
		"Employees": e,
	})
}

func (a *Front) GetRegister(c *fiber.Ctx) error {
	return c.Render("views/register", fiber.Map{})
}

func (a *Front) PostRegister(c *fiber.Ctx) error {
	var req RegisterRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Render("views/errors/error", fiber.Map{
			"Status": fmt.Sprintf("%d - %s", fiber.StatusInternalServerError, fiber.ErrInternalServerError),
			"Error":  err.Error()},
		)
	}

	hireDate, err := time.Parse("2006-01-02", req.HireDate)
	if err != nil {
		return c.Render("views/errors/error", fiber.Map{
			"Status": fmt.Sprintf("%d - %s", fiber.StatusBadRequest, fiber.ErrBadRequest),
			"Error":  err.Error()},
		)
	}

	e, err := a.Service.CreateEmployee(c.Context(), req.Name, req.Position, req.Email, req.Password, req.Presentation, hireDate)
	if err != nil {
		return c.Render("views/errors/error", fiber.Map{
			"Status": fmt.Sprintf("%d - %s", fiber.StatusBadRequest, fiber.ErrBadRequest),
			"Error":  err.Error()},
		)
	}

	// TODO: change to batch
	for _, s := range req.Skills {
		sID, err := a.Service.CreateSkill(c.Context(), s.Name)
		if err != nil {
			return c.Render("views/errors/error", fiber.Map{
				"Status": fmt.Sprintf("%d - %s", fiber.StatusBadRequest, fiber.ErrBadRequest),
				"Error":  err.Error(),
			})
		}

		err = a.Service.AddEmployeeSkill(c.Context(), e, sID, &s.Note, s.XP)
		if err != nil {
			return c.Render("views/errors/error", fiber.Map{
				"Status": fmt.Sprintf("%d - %s", fiber.StatusBadRequest, fiber.ErrBadRequest),
				"Error":  err.Error(),
			})
		}
	}

	return c.Redirect("/")
}

func (a *Front) Profile(c *fiber.Ctx) error {
	employeeID := c.Params("employee_id")
	if !govalidator.IsUUIDv4(employeeID) {
		return c.Render("views/errors/error", fiber.Map{
			"Status": fmt.Sprintf("%d - %s", fiber.StatusBadRequest, fiber.ErrBadRequest),
			"Error":  "employee_id is not a valid uuid"},
		)
	}

	employee, err := a.Service.FindEmployee(c.Context(), &employeeID)
	if err != nil {
		return c.Render("views/errors/error", fiber.Map{
			"Status": fmt.Sprintf("%d - %s", fiber.StatusNotFound, fiber.ErrNotFound),
			"Error":  err.Error()},
		)
	}

	return c.Render("views/profile", fiber.Map{"Employee": employee})
}

func (a *Front) GetProfileEdit(c *fiber.Ctx) error {
	employeeID := c.Params("employee_id")
	if !govalidator.IsUUIDv4(employeeID) {
		return c.Render("views/errors/error", fiber.Map{
			"Status": fmt.Sprintf("%d - %s", fiber.StatusBadRequest, fiber.ErrBadRequest),
			"Error":  "employee_id is not a valid uuid"},
		)
	}

	employee, err := a.Service.FindEmployee(c.Context(), &employeeID)
	if err != nil {
		return c.Render("views/errors/error", fiber.Map{
			"Status": fmt.Sprintf("%d - %s", fiber.StatusNotFound, fiber.ErrNotFound),
			"Error":  err.Error()},
		)
	}

	return c.Render("views/register", fiber.Map{"Employee": employee})
}

func (a *Front) PostProfileEdit(c *fiber.Ctx) error {
	var req RegisterRequest

	employeeID := c.Params("employee_id")
	if !govalidator.IsUUIDv4(employeeID) {
		return c.Render("views/errors/error", fiber.Map{
			"Status": fmt.Sprintf("%d - %s", fiber.StatusBadRequest, fiber.ErrBadRequest),
			"Error":  "employee_id is not a valid uuid"},
		)
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Render("views/errors/error", fiber.Map{
			"Status": fmt.Sprintf("%d - %s", fiber.StatusInternalServerError, fiber.ErrInternalServerError),
			"Error":  err.Error()},
		)
	}

	hireDate, err := time.Parse("2006-01-02", req.HireDate)
	if err != nil {
		return c.Render("views/errors/error", fiber.Map{
			"Status": fmt.Sprintf("%d - %s", fiber.StatusBadRequest, fiber.ErrBadRequest),
			"Error":  err.Error()},
		)
	}

	err = a.Service.UpdateEmployee(c.Context(), employeeID, req.Name, req.Position, req.Email, req.Password, req.Presentation, hireDate)
	if err != nil {
		return c.Render("views/errors/error", fiber.Map{
			"Status": fmt.Sprintf("%d - %s", fiber.StatusBadRequest, fiber.ErrBadRequest),
			"Error":  err.Error(),
		})
	}

	// TODO: changes to best practices
	err = a.Service.DeleteEmployeeSkills(c.Context(), &employeeID)
	if err != nil {
		return err
	}

	// TODO: change to batch
	for _, s := range req.Skills {
		sID, err := a.Service.CreateSkill(c.Context(), s.Name)
		if err != nil {
			return c.Render("views/errors/error", fiber.Map{
				"Status": fmt.Sprintf("%d - %s", fiber.StatusBadRequest, fiber.ErrBadRequest),
				"Error":  err.Error(),
			})
		}

		err = a.Service.AddEmployeeSkill(c.Context(), &employeeID, sID, &s.Note, s.XP)
		if err != nil {
			return c.Render("views/errors/error", fiber.Map{
				"Status": fmt.Sprintf("%d - %s", fiber.StatusBadRequest, fiber.ErrBadRequest),
				"Error":  err.Error(),
			})
		}
	}

	return c.Redirect(fmt.Sprintf("/profile/%s", employeeID))
}

func (a *Front) ProfileDeactivate(c *fiber.Ctx) error {
	var req DeactivateRequest

	employeeID := c.Params("employee_id")
	if !govalidator.IsUUIDv4(employeeID) {
		return c.Render("views/errors/error", fiber.Map{
			"Status": fmt.Sprintf("%d - %s", fiber.StatusBadRequest, fiber.ErrBadRequest),
			"Error":  "employee_id is not a valid uuid"},
		)
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Render("views/errors/error", fiber.Map{
			"Status": fmt.Sprintf("%d - %s", fiber.StatusInternalServerError, fiber.ErrInternalServerError),
			"Error":  err.Error()},
		)
	}

	terminationDate, err := time.Parse("2006-01-02", req.TaminationDate)
	if err != nil {
		return c.Render("views/errors/error", fiber.Map{
			"Status": fmt.Sprintf("%d - %s", fiber.StatusBadRequest, fiber.ErrBadRequest),
			"Error":  err.Error()},
		)
	}

	err = a.Service.DeactivateEmployee(c.Context(), &employeeID, &req.Password, terminationDate)
	if err != nil {
		return c.Render("views/errors/error", fiber.Map{
			"Status": fmt.Sprintf("%d - %s", fiber.StatusBadRequest, fiber.ErrBadRequest),
			"Error":  err.Error()},
		)
	}

	return c.Redirect("/")
}

func (a *Front) ProfileActivate(c *fiber.Ctx) error {
	var req ActivateRequest

	employeeID := c.Params("employee_id")
	if !govalidator.IsUUIDv4(employeeID) {
		return c.Render("views/errors/error", fiber.Map{
			"Status": fmt.Sprintf("%d - %s", fiber.StatusBadRequest, fiber.ErrBadRequest),
			"Error":  "employee_id is not a valid uuid"},
		)
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Render("views/errors/error", fiber.Map{
			"Status": fmt.Sprintf("%d - %s", fiber.StatusInternalServerError, fiber.ErrInternalServerError),
			"Error":  err.Error()},
		)
	}

	hireDate, err := time.Parse("2006-01-02", req.HireDate)
	if err != nil {
		return c.Render("views/errors/error", fiber.Map{
			"Status": fmt.Sprintf("%d - %s", fiber.StatusBadRequest, fiber.ErrBadRequest),
			"Error":  err.Error()},
		)
	}

	err = a.Service.ActivateEmployee(c.Context(), &employeeID, &req.Password, hireDate)
	if err != nil {
		return c.Render("views/errors/error", fiber.Map{
			"Status": fmt.Sprintf("%d - %s", fiber.StatusBadRequest, fiber.ErrBadRequest),
			"Error":  err.Error()},
		)
	}

	return c.Redirect("/")
}
