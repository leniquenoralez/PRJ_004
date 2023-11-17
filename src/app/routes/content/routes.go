package content

import "gordo/app/router"

func RegisterContentRoutes(r *router.Router) {
	r.Get("/api/content", r.HandleRequest(getAllContent))
}
