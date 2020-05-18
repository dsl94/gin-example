package product

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type ProductRepository struct {
	DB *gorm.DB
}

func ProvideProductRepository(DB *gorm.DB) ProductRepository  {
	return ProductRepository{DB: DB}
}

func (p *ProductRepository) FindAll() []Product  {
	var products []Product
	p.DB.Find(&products)

	return products
}

func (p *ProductRepository) FindById(id uint) Product {
	var product Product
	p.DB.Find(&product, id)

	return product
}

func (p *ProductRepository) Save(product Product) Product {
	p.DB.Save(&product)

	return product
}

func (p *ProductRepository) Delete(product Product) {
	p.DB.Delete(&product)
}

