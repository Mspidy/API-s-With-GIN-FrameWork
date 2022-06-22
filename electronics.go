package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Electronics struct {
	Subject string `json:"subject" form:"subject"`
	Brand   string `json:"brand" form:"brand"`
	Price   int    `json:"price" form:"price"`
}

func main() {
	db, err := sql.Open("mysql", "root:Prabhat@2022@tcp(127.0.0.1:3306)/accessories?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatalln(err)
	}
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Electronics API's Working fine")
	})

	router.POST("/access", func(c *gin.Context) {
		subject := c.Request.FormValue("subject")
		brand := c.Request.FormValue("brand")
		price := c.Request.FormValue("price")

		_, err := db.Exec("INSERT INTO electric(subject,brand,price) VALUES(?,?,?)", subject, brand, price)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("Insert Accessdata price{}", price)
		msg := fmt.Sprintf("Insert Successfully %d", price)
		c.JSON(http.StatusOK, gin.H{
			"msg": msg,
		})

	})
	router.Run(":9000")
}
