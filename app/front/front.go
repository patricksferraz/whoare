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
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	e, err := a.Service.SearchEmployees(c.Context(), &req.Q)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Render("index", fiber.Map{
		"Q":         req.Q,
		"Employees": e,
	})
}

func (a *Front) Register(c *fiber.Ctx) error {
	switch c.Method() {
	case "GET":
		return c.Render("register", fiber.Map{})

	case "POST":
		var req RegisterRequest

		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("error parsing body")
		}

		hireDate, _ := time.Parse("2006-01-02", req.HireDate) // TODO: handle error
		_, err := a.Service.CreateEmployee(c.Context(), req.Name, req.Position, req.Email, req.Password, req.Presentation, hireDate)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.Redirect("/")
	}

	return c.Status(fiber.StatusInternalServerError).SendString("method not found")
}

func (a *Front) Profile(c *fiber.Ctx) error {
	employeeID := c.Params("employee_id")
	if !govalidator.IsUUIDv4(employeeID) {
		return c.Status(fiber.StatusBadRequest).SendString("employee_id is not a valid uuid")
	}

	employee, err := a.Service.FindEmployee(c.Context(), &employeeID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("employee not found")
	}

	return c.Render("profile", fiber.Map{"Employee": employee})
}

func (a *Front) ProfileEdit(c *fiber.Ctx) error {
	employeeID := c.Params("employee_id")
	if !govalidator.IsUUIDv4(employeeID) {
		return c.Status(fiber.StatusBadRequest).SendString("employee_id is not a valid uuid")
	}

	employee, err := a.Service.FindEmployee(c.Context(), &employeeID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("employee not found")
	}

	switch c.Method() {
	case "GET":
		return c.Render("register", fiber.Map{"Employee": employee})

	case "POST":
		var req RegisterRequest

		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("error parsing body")
		}

		hireDate, _ := time.Parse("2006-01-02", req.HireDate) // TODO: handle error
		err := a.Service.UpdateEmployee(c.Context(), employeeID, req.Name, req.Position, req.Email, req.Password, req.Presentation, hireDate)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error()) // TODO: handle error (invalid password)
		}

		return c.Redirect(fmt.Sprintf("/profile/%s", employeeID))
	}

	return c.Status(fiber.StatusInternalServerError).SendString("method not found")
}
