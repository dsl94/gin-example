//+build wireinject

package main

import (
	"example-product/product"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

func initProductApi(db *gorm.DB) product.ProductAPI {
	wire.Build(product.ProvideProductRepository, product.ProvideProductService, product.ProvideProductAPI)

	return product.ProductAPI{}
}
