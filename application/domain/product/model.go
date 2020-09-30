package product

//Product reprensents Product model
type Product struct {
	ID    int    `gorm:"primary_key;auto_increment" json:"id"`
	Name  string `gorm:"size:255;not null;unique" json:"name"`
	Price string `gorm:"size:255;not null" json:"price"`
}
