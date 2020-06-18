package book

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

var (
	//DBConn of the connection- This should be in database,go
	DBConn *gorm.DB
)

//Book is a ORM model
type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

//GetBooks is to get books
func GetBooks(c *fiber.Ctx) {
	c.Send("All books")
	db := DBConn
	var book []Book
	db.Find(&book)
	c.JSON(book)
}

//GetBook is to get a single book
func GetBook(c *fiber.Ctx) {
	//c.Send("All Single book")
	id := c.Params("id")
	db := DBConn
	var book Book
	db.Find(&book, id)
	c.JSON(book)
}

//NewBook is to create a new book
func NewBook(c *fiber.Ctx) {
	//c.Send("Add New book")
	//	id := c.Params("id")
	db := DBConn
	/*var book Book
	book.Title = "Golang Essentials"
	book.Author = "Vathirajan"
	book.Rating = 4*/

	//we have hardcoded previously

	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		c.Status(503).Send(err)
		return
	}

	db.Create(&book)
	c.JSON(book)
}

//DeleteBook is to delete a book
func DeleteBook(c *fiber.Ctx) {
	//c.Send("Delete a book ")
	id := c.Params("id")
	db := DBConn
	var book Book
	db.First(&book, id)
	if book.Title == "" {
		c.Status(500).Send("No Book Found")
		return
	}
	db.Delete(&book, id)
	c.Send("Book deleted successfully")
}
