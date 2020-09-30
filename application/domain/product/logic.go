package product

import (
	"errors"
)

var (
	//ErrProductNotFound represents error returned when fetched product is not found
	ErrProductNotFound = errors.New("Product Not Found")
	//ErrProductNotCreated represents error returned when product was not created
	ErrProductNotCreated = errors.New("Product Not Created")
)

type productService struct {
	productRepo ProductRepository
}

type productRepository struct {
	name string
}

// NewProductService represents ProductService constructor
func NewProductService(productRepo ProductRepository) ProductService {
	return &productService{
		productRepo,
	}
}

// NewProductRepository represents ProductRepository constructor
func NewProductRepository(name string) ProductRepository {
	return &productRepository{
		name,
	}
}

func (s *productService) Find(id int) (*Product, error) {
	return s.productRepo.Find(id)
}

func (s *productService) Create(product *Product) error {
	return s.productRepo.Create(product)
}

func (r *productRepository) Find(id int) (*Product, error) {
	return r.Find(id)
}

func (r *productRepository) Create(product *Product) error {
	return r.Create(product)
}

func (r *productRepository) InitClient() {
	r.InitClient()
}
