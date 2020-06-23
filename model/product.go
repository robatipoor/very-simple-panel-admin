package product

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/spf13/viper"
)

var db *gorm.DB

// Product model
type Product struct {
	gorm.Model   `json:"-"`
	Name         string `gorm:"type:varchar(100)" json:"name"`
	ProductModel string `gorm:"type:varchar(100)" json:"model"`
	Description  string `gorm:"type:varchar(255)" json:"desc"`
	Price        string `gorm:"type:varchar(100)" json:"price"`
}

func init() {
	url := viper.GetString("database.url")
	dialect := viper.GetString("database.dialect")
	var err error
	db, err = gorm.Open(dialect, url)
	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(&Product{})
}

func InsertProduct(p *Product) {
	db.Create(p)
}

func FindById(id uint) *Product {
	p := &Product{}
	db.Find(p, id)
	return p
}

func FindAll() *[]Product {
	p := &[]Product{}
	db.Find(p)
	return p
}

func Update(id uint, p *Product) {
	db.Model(&Product{Model: gorm.Model{ID: id}}).Update(p)
}

func Delete(id uint) {
	db.Delete(&Product{Model: gorm.Model{ID: id}})
}
