package front

import (
	"embed"
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html"
	"github.com/patricksferraz/whoare/domain/service"
	"github.com/patricksferraz/whoare/infra/db"
	"github.com/patricksferraz/whoare/infra/repo"
)

//go:embed views/*
var views embed.FS

//go:embed public/*
var public embed.FS

var sessionStore = session.New()

func init() {
	sessionStore.RegisterType(Session{})
}

func StartFront(orm *db.DbOrm) {
	engine := html.NewFileSystem(http.FS(views), ".html")

	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "views/layouts/main",
	})

	repository := repo.NewRepository(orm)
	service := service.NewService(repository)
	front := NewFront(service, sessionStore)
	// middleware := NewMiddleware(sessionStore)

	app.Get("/", front.Index)
	app.Get("/register", front.GetRegister)
	app.Post("/register", front.PostRegister)
	app.Get("/profile/:employee_id", front.Profile)
	app.Get("/profile/:employee_id/edit", front.GetProfileEdit)
	app.Post("/profile/:employee_id/edit", front.PostProfileEdit)
	app.Post("/profile/:employee_id/deactivate", front.ProfileDeactivate)
	app.Post("/profile/:employee_id/activate", front.ProfileActivate)

	app.Use(recover.New())

	app.Use("/static", filesystem.New(filesystem.Config{
		Root:       http.FS(public),
		PathPrefix: "public",
		Browse:     false,
	}))

	app.Use(func(c *fiber.Ctx) error {
		return c.Render("errors/error", fiber.Map{
			"Status": fmt.Sprintf("%d - %s", fiber.StatusNotFound, fiber.ErrNotFound),
			"Error":  c.Context().Request.URI()},
		) // => 404 "Not Found"
	})

	err := app.Listen("0.0.0.0:8080")
	if err != nil {
		log.Fatal("cannot start app", err)
	}
}
