package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	_ "zen/books/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Books struct {
	ID             int     `json:"id"`
	Title          string  `json:"title"`
	ISBN           string  `json:"isbn"`
	Language       string  `json:"langauge,omitempty"`
	BookPublishers string  `json:"book_publishers,omitempty"`
	BookGenre      string  `json:"book_genre,omitempty"`
	BookAuthor     string  `json:"book_author,omitempty"`
	Price          float64 `json:"prices,omitempty"`
	Status         bool    `json:"status,omitempty"`
}

var books = []Books{
	{ID: 1, Title: "My First Book of ABC", ISBN: "9789380069401", Language: "English", BookPublishers: "Om Books", BookGenre: "Educational", BookAuthor: "Om Books Editorial Team", Price: 203.12, Status: true},
	{ID: 2, Title: "Fundamentals of Wavelets", ISBN: "8989380069401", Language: "English", BookPublishers: "Texh Books", BookGenre: "technical", BookAuthor: "Jaideva Goswami", Price: 803.22, Status: true},
	{ID: 3, Title: "Integration of the Indian States", ISBN: "2789380069901", Language: "English", BookPublishers: "Orient", BookGenre: "History", BookAuthor: "V P Menon", Price: 299.99, Status: true},
	{ID: 4, Title: "The Trial", ISBN: "3789380269403", Language: "English", BookPublishers: "Random House", BookGenre: "Fiction", BookAuthor: "Frank Kafka", Price: 403.24, Status: true},
	{ID: 5, Title: "Slaughterhouse Five", ISBN: "9989380069409", Language: "English", BookPublishers: "Random House", BookGenre: "Fiction", BookAuthor: "Kurt Vonnegut", Price: 345.99, Status: true},
	{ID: 6, Title: "Godaan", ISBN: "7989380079409", Language: "Hindi", BookPublishers: "Rupa", BookGenre: "Fiction", BookAuthor: "Premchand", Price: 45.99, Status: true},
	{ID: 7, Title: "Aavarana", ISBN: "9781536682830", Language: "Kannada", BookPublishers: "Sahitya Bhandara", BookGenre: "Fiction", BookAuthor: "S L Bhyrappa", Price: 295.99, Status: true},
}

// @title Books API
// @host localhost:5000
// @Basepath /
// @schemes http
func main() {
	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	//url := ginSwagger.URL("http://localhost:5000/swagger/swagger.json")

	router.GET("/books", getBooks)

	router.GET("/books/:id", getBookById)

	router.POST("/books", createBook)

	router.PUT("/books/:id", updateBook)

	router.PATCH("/books/:id", patchBook)

	router.DELETE("/books/:id", deleteBook)

	router.GET("/books/search", searchBook)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run("localhost:5000")
}

// List godoc
//
//	@Summary		Get All Books
//	@Description	Get All Books
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	Books
//	@Router			/books [get]
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

// @Summary		Get Book by ID
// @Description	get book by ID
// @ID				int
// @Accept			json
// @Produce		json
// @Success		200	{string}	string	"ok"
// @Router			/books/{some_id} [get]
func getBookById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Wrong format for id"})
		return
	}

	for _, a := range books {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}

// @Summary		Add a new book
// @Description	Add a new book
// @Accept			json
// @Produce		json
// @Success		201	{object}	Books	"Created"
// @Router			/books/{some_id} [post]
// @Param data body Books true "body data"
func createBook(c *gin.Context) {
	var newBook Books

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	for _, a := range books {
		if a.ID == newBook.ID {
			c.IndentedJSON(http.StatusConflict, gin.H{"message": "ID already exists"})
			return
		}
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

// @Summary		Update a book
// @Description	Update a book
// @ID				int
// @Accept			json
// @Produce		json
// @Success		200	{object}	Books	"ok"
// @Router			/books/{some_id} [put]
// @Param data body Books true "body data"
func updateBook(c *gin.Context) {
	var newBook Books

	if !strings.Contains(c.Request.Header.Get("Content-Type"), "json") {
		c.IndentedJSON(http.StatusUnsupportedMediaType, gin.H{"message": "Invalid Body Type"})
		return
	}

	if err := c.BindJSON(&newBook); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Malformed Request"})
		return
	}

	paramId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Wrong format for id"})
		return
	}

	if err != nil && paramId != newBook.ID {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID in URL does not match ID in request body"})
		return
	}

	counter := 0

	for _, a := range books {
		if a.ID == newBook.ID {
			books[counter] = newBook
			c.IndentedJSON(http.StatusOK, newBook)
			return
		}
		counter++
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "ID not found"})
}

// @Summary		Partially update a book
// @Description	Partially update a book
// @ID				int
// @Accept			json
// @Produce		json
// @Success		200	{object}	Books	"ok"
// @Router			/books/{some_id} [patch]
// @Param request body Books true "same as post body"
func patchBook(c *gin.Context) {
	//var newBook Books
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Wrong format for id"})
		return
	}

	if !strings.Contains(c.Request.Header.Get("Content-Type"), "json") {
		c.IndentedJSON(http.StatusUnsupportedMediaType, gin.H{"message": "Invalid Body Type"})
		return
	}

	// if err := c.BindJSON(&newBook); err != nil {
	// 	c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Malformed Request"})
	// 	return
	// }

	jsonData, err := c.GetRawData()

	fmt.Println(jsonData)

	if err != nil {
		fmt.Println("somethings wrong with the data sent")
	}

	x := map[string]string{}
	json.Unmarshal([]byte(jsonData), &x)

	//keys := make([]string, 0, len(x))
	//values := make([]string, 0, len(x))

	counter := 0

	for _, a := range books {
		if a.ID == id {
			for k, v := range x {
				switch k {
				case "title":
					books[counter].Title = v

				case "isbn":
					books[counter].ISBN = v

				case "langauge":
					books[counter].Language = v

				case "book_publishers":
					books[counter].BookPublishers = v

				case "book_genre":
					books[counter].BookGenre = v

				case "book_author":
					books[counter].BookAuthor = v

				case "prices":
					price, err := strconv.ParseFloat(v, 64)
					if err != nil {
						books[counter].Price = price
					} else {
						c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "price was of invalid type"})
					}

				case "status":
					status, err := strconv.ParseBool(v)
					if err != nil {
						books[counter].Status = status
					} else {
						c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "status was of invalid type"})
					}

				default:
					c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "unexpected field"})
					return
				}
			}
			c.IndentedJSON(http.StatusOK, books[counter])
			return
		}
		counter++
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "ID not found"})
}

// @Summary		Delete a Book
// @Description	Delete a book
// @ID				int
// @Accept			json
// @Produce		json
// @Success		204
// @Router			/books/{some_id} [delete]
func deleteBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Wrong format for id"})
		return
	}

	counter := 0

	for _, a := range books {
		if a.ID == id {
			books = append(books[:counter], books[counter+1:]...)
			c.Writer.WriteHeader(http.StatusNoContent)
			return
		}
		counter++
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "ID not found"})
}

// @Summary		Search Books
// @Description	Get Books by search criteria
// @Accept			json
// @Produce		json
// @Param			some_id	path		int		true	"Some ID"
// @Success		200		{string}	string	"ok"
// @Router			/books/search [get]
func searchBook(c *gin.Context) {
	c.AbortWithStatus(503)
}
