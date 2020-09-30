package mysql

import (
	"fmt"
	"log"
	"os"

	product "github.com/jekhedhri/tuto006/application/domain/product"
	"github.com/jinzhu/gorm"
)

// Client -
type Client struct {
	name string
	DB   *gorm.DB
}

// NewMySQLClient -
func NewMySQLClient() *Client {
	client := &Client{
		name: "MySqlClient",
	}
	return client
}

//InitDb is  InitDb
func (client *Client) InitDb() {

	//env := os.Getenv("ENV_environnement")
	DbDriver := os.Getenv("ENV_dbdriver")
	DbUser := os.Getenv("ENV_dbuser")
	DbPassword := os.Getenv("ENV_dbpassword")
	DbPort := os.Getenv("ENV_dbport")
	DbHost := os.Getenv("ENV_dbhost")
	DbName := os.Getenv("ENV_dbname")
	client.initMySQL(DbUser, DbPassword, DbHost, DbPort, DbName, DbDriver)
	client.createTables()
}

func (client *Client) initMySQL(DbUser, DbPassword, DbHost, DbPort, DbName, DbDriver string) {
	var err error
	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
	client.DB, err = gorm.Open(DbDriver, DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to %s database", DbDriver)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database", DbDriver)
	}
}

func (client *Client) createTables() {
	DbMode := os.Getenv("ENV_dbmode")
	if DbMode == "drop-create" {
		client.dropTableIfExists()
	}
	client.autoMigrate()

}

//TODO : entities by names
func (client *Client) dropTableIfExists() {
	var err error
	err = client.DB.Debug().DropTableIfExists(&product.Product{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
}

//TODO : entities by names
func (client *Client) autoMigrate() {
	var err error
	err = client.DB.Debug().AutoMigrate(&product.Product{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

}
