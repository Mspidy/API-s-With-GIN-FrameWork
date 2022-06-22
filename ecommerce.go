package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Accessories struct {
	Id        int    `json:"id" form:"id"`
	BrandName string `json:"brand_name" form:"brand_name"`
	Gadget    string `json:"gadget" form:"gadget"`
	Price     int    `json:"price" form:"price"`
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
	var a Accessories

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		//query := "Select * from user;"
		//defer db.Close()
		// result := db.QueryRow("SELECT * FROM accessdata ")
		// defer db.Close()
		// fmt.Println(result)
		c.String(http.StatusOK, "Working")
		// c.JSON(http.StatusOK, a)
	})

	router.POST("/access", func(c *gin.Context) {
		id := c.Request.FormValue("id")
		brandname := c.Request.FormValue("brand_name")
		gadget := c.Request.FormValue("gadget")
		price := c.Request.FormValue("price")

		_, err := db.Exec("INSERT INTO accessdata(id, brand_name, gadget, price) VALUES(?,?,?,?)", id, brandname, gadget, price)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("insert accessdata Id{}", id)
		msg := fmt.Sprintf("insert successful %d", id)
		c.JSON(http.StatusOK, gin.H{
			"msg": msg,
		})
	})

	router.Run(":8800")
}
