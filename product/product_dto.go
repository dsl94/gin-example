package product

type ProductDto struct {
	ID    uint   `json:"id,string,omitempty"`
	Code  string `json:"code"`
	Price uint   `json:"price"`
}
