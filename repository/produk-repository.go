package repository

import (
	"github.com/yogaabdi80/go-crud-example/config"
	"github.com/yogaabdi80/go-crud-example/model"
)

var product *model.Product

func GetProducts() ([]model.Product, error) {
	db := config.CreateConnection()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM products ORDER BY id DESC")
	products := []model.Product{}
	defer rows.Close()
	if err!=nil{
		return nil, err
	}
	for rows.Next(){
		product := model.Product{}
		err = rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
		)
		products = append(products, product)
	}
	err = rows.Err()
	if err != nil{
		return nil, err
	}
	return products,nil
}

func GetProduct(p *model.Product) (*model.Product, error){

	db := config.CreateConnection()
	defer db.Close()

	err := db.QueryRow("SELECT name, price FROM products WHERE id=$1",
		p.ID).Scan(&p.Name, &p.Price)
	if err != nil{
		return nil, err
	}
	return p, err
}

func CreateProduct(p *model.Product) (*model.Product, error){
	db := config.CreateConnection()
	defer db.Close()
	err := db.QueryRow(
        "INSERT INTO products(name, price) VALUES($1, $2) RETURNING *",
		p.Name, p.Price).Scan(&p.ID,&p.Name, &p.Price)
	if err != nil{
		return nil, err
	}
	return p, err
}

func UpdateProduct(p *model.Product) (*model.Product, error){
	db := config.CreateConnection()
	defer db.Close()
	err := db.QueryRow(
        "update products set name = $2, price = $3 where id = $1 RETURNING *", p.ID,
		p.Name, p.Price).Scan(&p.ID,&p.Name, &p.Price)
	if err != nil{
		return nil, err
	}
	return p, err
}