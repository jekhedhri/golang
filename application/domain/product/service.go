package product

//ProductService represents Product service abstraction
type ProductService interface {
	Find(id int) (*Product, error)
	Create(product *Product) error
}
