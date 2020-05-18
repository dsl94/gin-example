package product

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type ProductAPI struct {
	ProductService ProductService
}

func ProvideProductAPI(p ProductService) ProductAPI {
	return ProductAPI{ProductService: p}
}

func (p *ProductAPI) FindAll(c *gin.Context) {
	products := p.ProductService.FindAll()

	c.JSON(http.StatusOK, gin.H{"products": ToProductDTOs(products)})
}

func (p *ProductAPI) FindById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	product := p.ProductService.FindById(uint(id))

	c.JSON(http.StatusOK, gin.H{"product": ToProductDTO(product)})
}

func (p *ProductAPI) Create(c *gin.Context) {
	var productDTO ProductDto
	err := c.BindJSON(&productDTO)
	if err != nil {
		c.Status(http.StatusBadRequest)
		log.Fatalln(err)
		return
	}

	createdProduct := p.ProductService.Save(ToProduct(productDTO))

	c.JSON(http.StatusOK, gin.H{"product" :ToProductDTO(createdProduct)})
}

func (p *ProductAPI) Update(c *gin.Context) {
	var productDTO ProductDto
	err := c.BindJSON(&productDTO)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}

	id, _ :=  strconv.Atoi(c.Param("id"))
	product := p.ProductService.FindById(uint(id))
	if product == (Product{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	product.Code = productDTO.Code
	product.Price = productDTO.Price
	p.ProductService.Save(product)

	c.Status(http.StatusOK)
}

func (p *ProductAPI) Delete(c *gin.Context) {
	id, _ :=  strconv.Atoi(c.Param("id"))
	product := p.ProductService.FindById(uint(id))
	if product == (Product{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	p.ProductService.Delete(product)

	c.Status(http.StatusOK)
}
