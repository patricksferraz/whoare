package front

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/fiber/v2/utils"
)

type Middleware struct {
	Session *session.Store
}

func NewMiddleware(session *session.Store) *Middleware {
	return &Middleware{
		Session: session,
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

func (m *Middleware) RequireLogin(c *fiber.Ctx) error {
	// currSession, err := m.Session.Get(c)
	// if err != nil {
	// 	return err
	// }
	// user := currSession.Get("User")
	// defer currSession.Save()

	// if user == nil {
	// 	return c.Redirect("/login")
	// }

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
