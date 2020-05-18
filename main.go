package main

import (
	"example-product/product"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func initDB() *gorm.DB  {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=product password=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&product.Product{})

	return db
}

func main()  {
	db := initDB()
	defer db.Close()

	productApi := initProductApi(db)

	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/products", productApi.FindAll)
	r.GET("/products/:id", productApi.FindById)
	r.POST("/products", productApi.Create)
	r.PUT("/products/:id", productApi.Update)
	r.DELETE("/products/:id", productApi.Delete)

	err := r.Run()
	if err != nil {
		panic(err)
	}
}
