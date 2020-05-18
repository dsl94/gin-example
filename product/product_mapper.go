package product

func ToProduct(productDTO ProductDto) Product {
	return Product{Code: productDTO.Code, Price: productDTO.Price}
}

func ToProductDTO(product Product) ProductDto {
	return ProductDto{ID: product.ID, Code: product.Code, Price: product.Price}
}

func ToProductDTOs(products []Product) []ProductDto {
	productdtos := make([]ProductDto, len(products))

	for i, item := range products {
		productdtos[i] = ToProductDTO(item)
	}

	return productdtos
}