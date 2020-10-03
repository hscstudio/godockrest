package main

import (
	"net/http"
	"github.com/labstack/echo"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"log"
	"strconv"
)

type Article struct {
	Id    int
	Title  string
	Content  string
}


func main() {
	db, err := sql.Open("mysql", "gouser:gopwd@tcp(mysql:3306)/gorest")
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/articles", func (c echo.Context) error {
		selDB, err := db.Query("SELECT * FROM articles")
		if err != nil {
				panic(err.Error())
		}
		art := Article{}
		res := []Article{}
		for selDB.Next() {
				var id int
				var title , content string
				err = selDB.Scan(&id, &title , &content  )
				if err != nil {
						panic(err.Error())
				}
				art.Id = id
				art.Title  = title 
				art.Content  = content  
				res = append(res, art)
		}
	
		return c.JSON(http.StatusOK, res)
	})

	e.GET("/articles/:id", func(c echo.Context) error {
		id := c.Param("id")
		var title string
		var content string
		if err := db.QueryRow("SELECT title, content FROM articles WHERE id = ? LIMIT 1", id).Scan(&title, &content); 
		err != nil {
			// log.Fatal(err)
			return c.String(http.StatusOK, "Error")
		}

		intId, err := strconv.Atoi(id)
		if err != nil {
			return c.String(http.StatusOK, "Error")
		}
		response := Article{Id: intId, Title: title, Content : content}
		return c.JSON(http.StatusOK, response)
	})

	e.Logger.Fatal(e.Start(":81"))
}
