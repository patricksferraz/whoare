package front

import (
	"errors"
	"fmt"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/patricksferraz/whoare/domain/service"
)

type Middleware struct {
	Session *session.Store
	Service *service.Service
}

func NewMiddleware(session *session.Store, service *service.Service) *Middleware {
	return &Middleware{
		Session: session,
		Service: service,
	}
}

func (m *Middleware) CsrfProtection() func(*fiber.Ctx) error {
	return csrf.New(csrf.Config{
		// Next: func(c *fiber.Ctx) bool {
		// 	return false
		// },
		KeyLookup:      "form:_csrf",
		CookieName:     "csrf_",
		CookieSameSite: "Strict",
		Expiration:     1 * time.Hour,
		KeyGenerator:   utils.UUID,
		ContextKey:     "token",
	})
}

func (m *Middleware) RequirePassword(c *fiber.Ctx) error {
	var req PasswordBody

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

	e, err := m.Service.FindEmployee(c.Context(), &employeeID)
	if err != nil {
		return c.Render("views/errors/error", fiber.Map{
			"Status": fmt.Sprintf("%d - %s", fiber.StatusInternalServerError, fiber.ErrInternalServerError),
			"Error":  err.Error()},
		)
	}

	if err := e.CompareHashAndPassword(req.Password); err != nil {
		return c.Render("views/errors/error", fiber.Map{
			"Status": fmt.Sprintf("%d - %s", fiber.StatusInternalServerError, fiber.ErrInternalServerError),
			"Error":  errors.New("invalid password")},
		)
	}

	return c.Next()
}

// func (m *Middleware) InitSession(c *fiber.Ctx) error {
// 	currSession, err := m.Session.Get(c)
// 	if err != nil {
// 		return err
// 	}

// 	r := currSession.Get("data")
// 	if r == nil {
// 		currSession.Set("data", Session{})
// 		currSession.Save()
// 	}

// 	return c.Next()
// }
