package front

import (
	"time"

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
			// TODO: add redirect to register steps
		}

		// TODO: add feedback
		return c.Redirect("/")
	}

	return c.Status(fiber.StatusInternalServerError).SendString("method not found")
}

// func (a *Front) Profile(c *fiber.Ctx) error {
// 	employeeID := c.Params("employee_id")
// 	if !govalidator.IsUUIDv4(employeeID) {
// 		return c.Status(fiber.StatusBadRequest).SendString("guest_id is not a valid uuid")
// 	}

// 	employee, err := a.Service.FindEmployee(c.Context(), &employeeID)
// 	if err != nil {
// 		return c.Status(fiber.StatusNotFound).SendString("employee not found")
// 	}
// 	// currSession, err := sessionStore.Get(c)
// 	// if err != nil {
// 	// 	return err
// 	// }
// 	// sessionUser := currSession.Get("User").(fiber.Map)
// 	// release the currSession
// 	// err = currSession.Save()
// 	// if err != nil {
// 	// 	return err
// 	// }

// 	// if sessionUser["Name"] == "" {
// 	// 	return c.Status(fiber.StatusBadRequest).SendString("User is empty")
// 	// }
// 	// username := sessionUser["Name"].(string)

// 	return c.Render("profile", fiber.Map{
// 		"Employee": employee,
// 		// "balance":   accounts[username],
// 		"csrfToken": c.Locals("token"),
// 	})
// }

// // TODO: removes only test
// func (a *Front) CreateEmployee(c *fiber.Ctx) error {
// 	// var req CreateGuestRequest

// 	// if err := c.BodyParser(&req); err != nil {
// 	// 	return c.Status(fiber.StatusInternalServerError).JSON(HTTPResponse{Msg: err.Error()})
// 	// }
// 	for i := 0; i < 30; i++ {
// 		fake, _ := faker.New("en")
// 		name := fake.Name()
// 		position := fake.JobTitle()
// 		email := fake.Email()
// 		password := fake.Characters(10)
// 		description := fake.Sentence(100, true)

// 		_, err := a.Service.CreateEmployee(c.Context(), name, position, email, password, description, time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC))
// 		if err != nil {
// 			return c.Status(fiber.StatusForbidden).SendString(err.Error())
// 		}
// 	}

// 	return c.Status(fiber.StatusOK).SendString("ok")
// }
