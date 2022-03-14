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
		return c.Render("errors/error", fiber.Map{"Status": fiber.StatusBadRequest, "Error": err.Error()})
	}

	e, err := a.Service.SearchEmployees(c.Context(), &req.Q)
	if err != nil {
		return c.Render("errors/error", fiber.Map{"Status": fiber.StatusInternalServerError, "Error": err.Error()})
	}

	return c.Render("index", fiber.Map{
		"Q":         req.Q,
		"Employees": e,
	})
}

func (a *Front) GetRegister(c *fiber.Ctx) error {
	return c.Render("register", fiber.Map{})
}

func (a *Front) PostRegister(c *fiber.Ctx) error {
	var req RegisterRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Render("errors/error", fiber.Map{"Status": fiber.StatusInternalServerError, "Error": err.Error()})
	}

	hireDate, err := time.Parse("2006-01-02", req.HireDate)
	if err != nil {
		return c.Render("errors/error", fiber.Map{"Status": fiber.StatusBadRequest, "Error": err.Error()})
	}

	_, err = a.Service.CreateEmployee(c.Context(), req.Name, req.Position, req.Email, req.Password, req.Presentation, hireDate)
	if err != nil {
		return c.Render("errors/error", fiber.Map{"Status": fiber.StatusInternalServerError, "Error": err.Error()})
	}

	return c.Redirect("/")
}

func (a *Front) Profile(c *fiber.Ctx) error {
	employeeID := c.Params("employee_id")
	if !govalidator.IsUUIDv4(employeeID) {
		return c.Render("errors/error", fiber.Map{"Status": fiber.StatusBadRequest, "Error": "employee_id is not a valid uuid"})
	}

	employee, err := a.Service.FindEmployee(c.Context(), &employeeID)
	if err != nil {
		return c.Render("errors/error", fiber.Map{"Status": fiber.StatusNotFound, "Error": err.Error()})
	}

	return c.Render("profile", fiber.Map{"Employee": employee})
}

func (a *Front) GetProfileEdit(c *fiber.Ctx) error {
	employeeID := c.Params("employee_id")
	if !govalidator.IsUUIDv4(employeeID) {
		return c.Render("errors/error", fiber.Map{"Status": fiber.StatusBadRequest, "Error": "employee_id is not a valid uuid"})
	}

	employee, err := a.Service.FindEmployee(c.Context(), &employeeID)
	if err != nil {
		return c.Render("errors/error", fiber.Map{"Status": fiber.StatusNotFound, "Error": err.Error()})
	}

	return c.Render("register", fiber.Map{"Employee": employee})
}

func (a *Front) PostProfileEdit(c *fiber.Ctx) error {
	var req RegisterRequest

	employeeID := c.Params("employee_id")
	if !govalidator.IsUUIDv4(employeeID) {
		return c.Render("errors/error", fiber.Map{"Status": fiber.StatusBadRequest, "Error": "employee_id is not a valid uuid"})
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Render("errors/error", fiber.Map{"Status": fiber.StatusInternalServerError, "Error": err.Error()})
	}

	hireDate, err := time.Parse("2006-01-02", req.HireDate)
	if err != nil {
		return c.Render("errors/error", fiber.Map{"Status": fiber.StatusBadRequest, "Error": err.Error()})
	}

	err = a.Service.UpdateEmployee(c.Context(), employeeID, req.Name, req.Position, req.Email, req.Password, req.Presentation, hireDate)
	if err != nil {
		return c.Render("errors/error", fiber.Map{"Status": fiber.StatusInternalServerError, "Error": err.Error()})
	}

	return c.Redirect(fmt.Sprintf("/profile/%s", employeeID))
}

func (a *Front) ProfileDelete(c *fiber.Ctx) error {
	var req DeleteRequest

	employeeID := c.Params("employee_id")
	if !govalidator.IsUUIDv4(employeeID) {
		return c.Render("errors/error", fiber.Map{"Status": fiber.StatusBadRequest, "Error": "employee_id is not a valid uuid"})
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Render("errors/error", fiber.Map{"Status": fiber.StatusInternalServerError, "Error": err.Error()})
	}

	err := a.Service.DeleteEmployee(c.Context(), &employeeID, &req.Password)
	if err != nil {
		return c.Render("errors/error", fiber.Map{"Status": fiber.StatusBadRequest, "Error": err.Error()})
	}

	return c.Redirect("/")
}
