package mocked

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/jekhedhri/tuto006/application/domain/product"
)

type mockedRepository struct {
	name   string
	client Client
}

// NewMockedRepository -
func NewMockedRepository() (product.ProductRepository, error) {
	repo := &mockedRepository{
		name:   "Mocked",
		client: *NewMockClient(),
	}
	return repo, nil
}

func (r *mockedRepository) Find(id int) (*product.Product, error) {
	fmt.Println("Find from Mocked datas")
	response, err := http.Get(r.client.baseURL + r.client.port + "/products/" + fmt.Sprint(id))
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	product := &product.Product{}
	json.Unmarshal(responseData, &product)
	return product, nil
}

func (r *mockedRepository) Create(redirect *product.Product) error {
	fmt.Println("Store from Mocked datas")
	return nil
}

func (r *mockedRepository) InitClient() {
	name := r.client.name
	fmt.Println(name)
}

func (r *mockedRepository) GetName() string {
	return r.name
}
