package routes

import (
	"neuza/backend/model"
)
type RoutesConfig struct {
	Model        interface{}
	BasePath     string
	RequiresAuth map[string]bool
}

var routesConfig = map[Route]RoutesConfig{
	"articles": {
		Model:   	model.Article{},
		BasePath: "/articles",
		RequiresAuth: map[string]bool{
			"GetAll": false,
			"GetOne": false,
			"Create": true,
			"Update": true,
			"Delete": true,
		},
	},
}