package main

import (
	"log"
	"os"

	"book"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//Method to setup logging
func setupLogging() {
	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()
	log.SetOutput(file)
}

//init database
func initDatabase() {
	var err error
	book.DBConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		//log.Fatal(err,)
		//panic("failed to connect to sql db")
		log.Fatal(err)
	}
	log.Println("database connection created successfully")
	book.DBConn.AutoMigrate(&book.Book{})
	log.Println("Data auto-migrated successfully")

}

//return nil, errors.Wrap(error, "Error in opening. Please place the file in the folder")
func setupRoutes(app *fiber.App) {
	app.Get("/api/book", book.GetBooks)
	app.Get("/api/book/:id", book.GetBook)
	app.Post("/api/book", book.NewBook)
	app.Delete("/api/book/:id", book.DeleteBook)

}
func main() {
	initDatabase()
	app := fiber.New()
	//app.Get("/", helloworld)
	setupRoutes(app)
	app.Listen(8081)
	defer book.DBConn.Close()
}
func helloworld(c *fiber.Ctx) {
	c.Send("Hello Vathi")
}
