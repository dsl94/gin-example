package product

type ProductService struct {
	ProductRepository ProductRepository
}

func ProvideProductService(p ProductRepository) ProductService {
	return ProductService{ProductRepository: p}
}

func (p *ProductService) FindAll() []Product {
	return p.ProductRepository.FindAll()
}

func (p *ProductService) FindById(id uint) Product  {
	return p.ProductRepository.FindById(id)
}

func (p *ProductService) Save(product Product) Product {
	p.ProductRepository.Save(product)

	return product
}

func (p *ProductService) Delete(product Product) {
	p.ProductRepository.Delete(product)
}
