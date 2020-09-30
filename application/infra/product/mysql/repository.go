package mysql

import (
	"errors"
	"fmt"

	product "github.com/jekhedhri/tuto006/application/domain/product"

	// MySQL
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type mysqlRepository struct {
	name   string
	client Client
}

// NewMySQLRepository -
func NewMySQLRepository() (product.ProductRepository, error) {
	repo := &mysqlRepository{
		name:   "MySql",
		client: *NewMySQLClient(),
	}
	return repo, nil
}

func (r *mysqlRepository) InitClient() {
	r.client.InitDb()
}

func (r *mysqlRepository) Find(id int) (*product.Product, error) {
	fmt.Println("Find from Mysql")
	var err error
	result := product.Product{}
	r.client.DB.Debug().Model(product.Product{}).Where("id = ?", id).Take(&result)
	if err != nil {
		return &result, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &result, errors.New("Product Not Found")
	}

	return &result, err
}

func (r *mysqlRepository) Create(redirect *product.Product) error {
	fmt.Println("Store from Mysql")

	return nil
}

func (r *mysqlRepository) GetName() string {
	return r.name
}
