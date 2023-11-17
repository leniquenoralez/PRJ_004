package app

import (
	"gordo/app/router"
	"gordo/app/routes"
)

func InitServer() {
	r := &router.Router{}
	r.InitRouter()
	routes.RegisterRoutes(r)
	r.Run(":5000")

}
