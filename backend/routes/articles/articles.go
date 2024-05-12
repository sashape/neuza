package routes

import (
	// "neuza/backend/database"
	// "neuza/backend/responses"

	// "github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
}

// func SetupArticleRoutes(app *fiber.App, database *database.Database) {
// 	db = database.DB
// 	db.AutoMigrate(&Article{})
// 	app.Get("/articles", getAllArticles)
// 	app.Get("/articles/:id", getArticle)
// 	app.Post("/articles", createArticle)
// 	app.Put("/articles/:id", updateArticle)
// 	app.Delete("/articles/:id", deleteArticle)
// }

// func getAllArticles(c *fiber.Ctx) error {
// 	var articles []Article
// 	result := db.Find(&articles)

// 	if result.Error != nil {
// 		switch err := result.Error; err {
// 		case gorm.ErrRecordNotFound:
// 			return responses.NotFound(c, "Articles not found")
// 		default:
// 			return responses.InternalServerError(c, "Server error")
// 		}
// 	}

// 	return responses.OK(c, fiber.Map{"articles": articles})
// }

// func getArticle(c *fiber.Ctx) error {
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
// 		return responses.BadRequest(c, "Data validation error " + err.Error())
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
