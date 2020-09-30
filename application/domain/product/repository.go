package product

//ProductRepository represents Product repository abstraction
type ProductRepository interface {
	Find(id int) (*Product, error)
	Create(product *Product) error
	InitClient()
}
