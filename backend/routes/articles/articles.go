package articles
import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type Article struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
}
func setupDatabase() {
	dsn := "host=database user=user password=password1234 dbname=user port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Не удалось подключиться к базе данных")
	}

	database.AutoMigrate(&Article{})

	db = database
}

func SetupArticleRoutes(app *fiber.App) {
    app.Get("/articles", getAllArticles)
    app.Get("/articles/:id", getArticle)
    app.Post("/articles", createArticle)
    app.Put("/articles/:id", updateArticle)
    app.Delete("/articles/:id", deleteArticle)
}

func getAllArticles(c *fiber.Ctx) error {
	var articles []Article
	result := db.Find(&articles)

	if result.Error != nil {
		switch err := result.Error; err {
		case gorm.ErrRecordNotFound:
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Articles not found",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Server error",
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"answer": articles})
}

func getArticle(c *fiber.Ctx) error {
	id := c.Params("id")
	var article Article
	db.First(&article, id)
	if article.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Article not found",
		})
	}
	return c.JSON(article)
}

func createArticle(c *fiber.Ctx) error {
	article := new(Article)
	if err := c.BodyParser(article); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Data validation error",
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
			"message": "Article not found",
		})
	}
	updatedArticle := new(Article)
	if err := c.BodyParser(updatedArticle); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Data validation error",
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
			"message": "Article not found",
		})
	}
	db.Delete(&article)
	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"message": "Article deleted",
	})
}
