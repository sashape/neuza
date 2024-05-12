package routes

import (
	"neuza/backend/database"
	"neuza/backend/model"

	"github.com/gofiber/fiber/v2"
)
type Route string

type Methods struct {
	GetOne func(c *fiber.Ctx) error
	GetAll func(c *fiber.Ctx) error
	Create func(c *fiber.Ctx) error
	Update func(c *fiber.Ctx) error
	Delete func(c *fiber.Ctx) error
}

func SetupRoutes(app *fiber.App, database *database.Database, route Route) {
	app.Get(routesConfig[route].BasePath, authenticate(routesConfig[route].RequiresAuth["GetAll"]), func(c *fiber.Ctx) error {
		articles := model.GetAll(routesConfig[route].Model, database, c)
		return c.JSON(articles)
})
	// app.Get(config.BasePath+"/:id", authenticate(config.RequiresAuth["GetOne"]), routes.GetOne)
	// app.Post(config.BasePath, authenticate(config.RequiresAuth["Create"]), routes.Create)
	// app.Put(config.BasePath+"/:id", authenticate(config.RequiresAuth["Update"]), routes.Update)
	// app.Delete(config.BasePath+"/:id", authenticate(config.RequiresAuth["Delete"]), routes.Delete)

}
// func getAll(c *fiber.Ctx) error {
// 	id := c.Params("id")
// 	var article Article
// 	db.First(&article, id)
// 	if article.ID == 0 {
// 		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
// 			"message": "Article not found",
// 		})
// 	}
// 	return c.JSON(article)
// }
// func getOne(c *fiber.Ctx) error {
// 	id := c.Params("id")
// 	var article Article
// 	db.First(&article, id)
// 	if article.ID == 0 {
// 		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
// 			"message": "Article not found",
// 		})
// 	}
// 	return c.JSON(article)
// }

// func createArticle(c *fiber.Ctx) error {
// 	article := new(Article)
// 	if err := c.BodyParser(article); err != nil {
// 		return responses.BadRequest(c, "Data validation error "+err.Error())
// 	}
// 	db.Create(&article)
// 	return c.JSON(article)
// }

// func updateArticle(c *fiber.Ctx) error {
// 	id := c.Params("id")
// 	var article Article
// 	db.First(&article, id)
// 	if article.ID == 0 {
// 		return responses.NotFound(c, "Article not found")
// 	}
// 	updatedArticle := new(Article)
// 	if err := c.BodyParser(updatedArticle); err != nil {
// 		return responses.BadRequest(c, "Data validation error")
// 	}
// 	db.Model(&article).Updates(updatedArticle)

// 	return responses.OK(c, article)
// }

// func deleteArticle(c *fiber.Ctx) error {
// 	id := c.Params("id")
// 	var article Article
// 	db.First(&article, id)
// 	if article.ID == 0 {
// 		return responses.NotFound(c, "Article not found")
// 	}
// 	db.Delete(&article)
// 	return responses.OK(c, "Article deleted")
// }

func authenticate(requireAuth bool) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if requireAuth {
			panic("need auth")
		}

		return c.Next()
	}
}
