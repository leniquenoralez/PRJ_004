package routes

import (
	"gordo/app/router"
	"gordo/app/routes/content"
)

func RegisterRoutes(r *router.Router) {
	content.RegisterContentRoutes(r)
}
