package product

import (
	"log"
	"os"

	"github.com/jekhedhri/tuto006/application/domain/product"
	rr "github.com/jekhedhri/tuto006/application/infra/product/mocked"
	mr "github.com/jekhedhri/tuto006/application/infra/product/mysql"
)

// InjectProductService -
func InjectProductService() (product.ProductService, error) {
	repo := chooseRepo()
	repo.InitClient()
	service := product.NewProductService(repo)
	return service, nil
}

func chooseRepo() product.ProductRepository {
	infra := os.Getenv("Infra")
	switch infra {
	case "api":
		repo, err := rr.NewMockedRepository()
		if err != nil {
			log.Fatal(err)
		}
		return repo
	case "db":
		repo, err := mr.NewMySQLRepository()
		if err != nil {
			log.Fatal(err)
		}
		return repo
	}
	return nil
}
