package main

import (
	"github.com/gofiber/fiber"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
}

var db *gorm.DB

func setupDatabase() {
	// Подключение к базе данных PostgreSQL
	dsn := "host=neuza-database user=user password=password1234 dbname=user port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Не удалось подключиться к базе данных")
	}

	// Автомиграция таблицы Article
	database.AutoMigrate(&Article{})

	db = database
}

func main() {
	app := fiber.New()

	app.Get("/", homepage)

	setupDatabase()

	app.Get("/articles", getAllArticles)
	app.Get("/articles/:id", getArticle)
	app.Post("/articles", createArticle)
	app.Put("/articles/:id", updateArticle)
	app.Delete("/articles/:id", deleteArticle)

	app.Listen(8000)
}

func homepage(c *fiber.Ctx) {
	c.Send("NEUZA API")
	// hi
}
func getAllArticles(c *fiber.Ctx) error {
	var articles []Article
	db.Find(&articles)
	return c.JSON(articles)
}

func getArticle(c *fiber.Ctx) error {
	id := c.Params("id")
	var article Article
	db.First(&article, id)
	if article.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Статья не найдена",
		})
	}
	return c.JSON(article)
}

func createArticle(c *fiber.Ctx) error {
	article := new(Article)
	if err := c.BodyParser(article); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Ошибка валидации данных",
		})
	}
	db.Create(&article)
	return c.JSON(article)
}

func updateArticle(c *fiber.Ctx) error {
	id := c.Params("id")
	var article Article
	db.First(&article, id)
	if article.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Статья не найдена",
		})
	}
	updatedArticle := new(Article)
	if err := c.BodyParser(updatedArticle); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Ошибка валидации данных",
		})
	}
	db.Model(&article).Updates(updatedArticle)
	return c.JSON(article)
}

func deleteArticle(c *fiber.Ctx) error {
	id := c.Params("id")
	var article Article
	db.First(&article, id)
	if article.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Статья не найдена",
		})
	}
	db.Delete(&article)
	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"message": "Статья удалена",
	})
}
