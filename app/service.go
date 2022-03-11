package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/manveru/faker"
	"github.com/patricksferraz/whoare/domain/entity"
	"github.com/patricksferraz/whoare/domain/service"
)

type AppService struct {
	Service *service.Service
	Session *session.Store
}

func NewAppService(service *service.Service, session *session.Store) *AppService {
	return &AppService{
		Service: service,
		Session: session,
	}
}

func (a *AppService) Index(c *fiber.Ctx) error {
	var req SearchEmployeesRequest
	var employees []*entity.Employee
	var count *int64

	err := c.QueryParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	filter := entity.NewFilterEmployee(req.Q, req.SearchBy, req.SortBy, req.PageSize, req.Page)
	switch req.SearchBy {
	case "skill":
		employees, count, err = a.Service.FindEmployeesBySkill(c.Context(), filter)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
	default:
		employees, count, err = a.Service.FindEmployeesByName(c.Context(), filter)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
	}

	// TODO: changes for equation
	var firstIdx int64 = 0
	var lastIdx int64 = 0
	if *count >= (filter.PageSize*(filter.Page-1))+1 {
		firstIdx = (filter.PageSize * (filter.Page - 1)) + 1
		lastIdx = *count
	}
	if *count > filter.PageSize*filter.Page {
		lastIdx = filter.PageSize * filter.Page
	}

	return c.Render("index", fiber.Map{
		"Employees": employees,
		"Filter":    filter,
		"G": fiber.Map{
			"Count":    *count,
			"FirstIdx": firstIdx,
			"LastIdx":  lastIdx,
			"Pages": func() interface{} {
				var pages []struct {
					Idx int64
				}
				l := (*count / filter.PageSize)
				if l > 0 {
					l += 1
				}
				for i := int64(1); i <= l; i++ {
					pages = append(pages, struct{ Idx int64 }{Idx: i})
				}
				return pages
			}(),
		},
	})
}

// TODO: removes only test
func (a *AppService) CreateEmployee(c *fiber.Ctx) error {
	// var req CreateGuestRequest

	// if err := c.BodyParser(&req); err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(HTTPResponse{Msg: err.Error()})
	// }
	for i := 0; i < 30; i++ {
		fake, _ := faker.New("en")
		name := fake.Name()
		position := fake.JobTitle()
		email := fake.Email()
		password := fake.Characters(10)

		_, err := a.Service.CreateEmployee(c.Context(), name, position, email, password)
		if err != nil {
			return c.Status(fiber.StatusForbidden).SendString(err.Error())
		}
	}

	return c.Status(fiber.StatusOK).SendString("ok")
}
