package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	username = "root"
	password = "root"
	hostname = "127.0.0.1:3306"
	dbname   = "bookshop"
)

// Book struct to represent the book member fields
type Book struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	Author        string  `json:"author"`
	Price         float32 `json:"price"`
	Pages         int     `json:"pages"`
	DatePublished string  `json:"date_published"`
}

// BooksCollections used to have an array of book
type BooksCollections struct {
	Books []Book `json:"books"`
}

var db *sql.DB

// dataSourceName used to prepare the DSN for SQL
func dataSourceName(dbName string) string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
	log.Println("DataSourceName = ", dsn)
	return dsn
}

// dbConnection use to establish the connection and create the database
func dbConnection(dsn string) (*sql.DB, error) {
	// open the database - DSN gere is "root:root@tcp(127.0.0.1:3306)/"
	database, err := sql.Open("mysql", dataSourceName(""))
	if err != nil {
		log.Printf("Error %s when opening DB\n", err)
		return nil, err
	}

	// create a database
	result, err := database.Exec("CREATE DATABASE IF NOT EXISTS " + dbname)
	if err != nil {
		log.Printf("Error %s when creating DB\n", err)
		return nil, err
	}

	no, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error %s when fetching rows", err)
		return nil, err
	}
	log.Printf("rows affected %d\n", no)

	// close the database still we wanted to create the new table
	database.Close()

	database, err = sql.Open("mysql", dataSourceName(dbname))
	if err != nil {
		log.Printf("Error %s when open the new table", err)
		return nil, err
	}

	database.SetMaxOpenConns(10)
	database.SetConnMaxIdleTime(10)

	// assign database to global for all API handler usage
	db = database

	return database, nil
}

func main() {
	fmt.Println("MySQL DB operations in GO using gin http web framework")
	/*
		// Open up our database connection.
		// Set up a database on my local machine using MySQL Command Line. The database is called books
		database, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/books")
		// handle if there is an error opening the connection
		if err != nil {
			panic(err.Error())
		}
	*/

	db, err := dbConnection(dbname)
	if err != nil {
		log.Println("Received error from dbConnection = ", err)
		return
	}

	// create the book table if not exists
	query := "create table IF NOT EXISTS book(id varchar(36) primary key, name varchar(255) not null, author varchar(255) not null, price decimal(5,2) not null, pages int, date_published date)"
	result, err := db.Exec(query)
	if err != nil {
		log.Println("Received error while creating table = ", err)
		return
	}
	no, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error %s when fetching rows", err)
		return
	}
	log.Printf("rows affected %d\n", no)

	// defer the close till after the main function has finished executing
	defer db.Close()

	/* // Execute the query
	results, err := db.Query("SELECT * FROM book")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	books := BooksCollections{}
	for results.Next() {
		var book Book
		// for each row, scan the result into our tag composite object
		err = results.Scan(&book.ID, &book.Name, &book.Author, &book.Price, &book.Pages, &book.DatePublished)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		books.Books = append(books.Books, book)
		// and then print out the book attributes
		fmt.Println(book)

		fmt.Println(books)
	} */

	r := route()
	_ = r.Run(":8080")
}

func route() *gin.Engine {
	r := gin.Default()

	// use the ping command to test the API
	// "curl -X GET http://localhost:8080/ping"
	r.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Hello.. Welcome!")
	})

	// REST CRUD operations for books MySQL database
	r.GET("/books", listBooks)
	r.POST("/books", createBook)
	r.DELETE("/books/:id", deleteBook)
	r.PATCH("/books/:id", patchBook)

	return r
}

// listBooks handler to list of books in json format
func listBooks(c *gin.Context) {
	// Execute the query
	results, err := db.Query("SELECT * FROM book")
	if err != nil {
		//panic(err.Error()) // proper error handling instead of panic in your app
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	books := BooksCollections{}
	for results.Next() {
		var book Book
		// for each row, scan the result into our tag composite object
		err = results.Scan(&book.ID, &book.Name, &book.Author, &book.Price, &book.Pages, &book.DatePublished)
		if err != nil {
			//panic(err.Error()) // proper error handling instead of panic in your app
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		books.Books = append(books.Books, book)
	}

	fmt.Println("List API Success\n", books)
	c.JSON(http.StatusOK, books)
}

// createBook handler to creat the book entry into database
func createBook(c *gin.Context) {
	var book Book

	c.BindJSON(&book)
	fmt.Println(book)

	// prepare and execute query to insert the book values
	query := "INSERT into book values(?,?,?,?,?,?)"
	statement, err := db.Prepare(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	id := uuid.New()
	result, err := statement.Exec(id, book.Name, book.Author, book.Price, book.Pages, book.DatePublished)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	fmt.Printf("The last inserted row id: %d\n", lastId)

	statement.Close()

	book.ID = id.String()

	fmt.Println("Create API Success\n", book)
	c.JSON(http.StatusCreated, book)
}

// deleteBook handler to delete the book entry using the uuid
func deleteBook(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("Delete api called and ID = ", id)

	_, err := db.Exec("delete from book where id = ?", id)
	if err != nil {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	fmt.Println("Delete API Success")
	c.JSON(http.StatusNoContent, nil)
}

// patchBook handler to patch the book entry using uuid
func patchBook(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("PATCH api called and ID = ", id)

	// get the book entry first using uuid. If not respond with not_found
	results, err := db.Query("SELECT * FROM book where BINARY id = UNHEX(?)", id)
	if err != nil {
		fmt.Println("PATCH api failed ID not found")
		c.JSON(http.StatusNotFound, nil)
		return
	}
	var book Book
	err = results.Scan(&book.ID, &book.Name, &book.Author, &book.Price, &book.Pages, &book.DatePublished)

	// get the patch request body and extract the values
	var updates Book
	c.BindJSON(&updates)
	fmt.Println("Update req are: ", updates)

	// use the input values if need sto be updated
	if updates.Name != "" {
		book.Name = updates.Name
	}
	if updates.Author != "" {
		book.Author = updates.Author
	}
	book.Price = updates.Price
	book.Pages = updates.Pages
	if updates.DatePublished != "" {
		book.DatePublished = updates.DatePublished
	}

	// prepare the update query and execute it
	query := "update book set name=?, author=?, price=?, pages=?, date_published=? where id=?"
	statement, err := db.Prepare(query)
	if err != nil {
		fmt.Println("Patch API Prepare failed")
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	result, err := statement.Exec(book.Name, book.Author, book.Price, book.Pages, book.DatePublished, id)
	if err != nil {
		fmt.Println("Patch API Exec failed")
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	fmt.Println("result = ", result)

	value, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Patch RowsAffected failed")
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	fmt.Println(" value = ", value)
	statement.Close()

	book.ID = id
	fmt.Println("Patch API Success\n", book)
	c.JSON(http.StatusOK, book)
}
