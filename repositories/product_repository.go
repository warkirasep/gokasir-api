package repositories

import (
	"database/sql"
	"errors"
	"gokasir-api/models"
)

type ProductRepository struct{
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository{
	return &ProductRepository{db: db}
}

func (r *ProductRepository) GetAll(name string) ([]models.Product, error) {
	query := "SELECT id, name, stock, price, category_id FROM products"

	var args []interface{}

	if name != "" {
		query += " WHERE name ILIKE $1"
		args = append(args, "%"+name+"%")
	}
	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]models.Product, 0)
	for rows.Next() {
		var p models.Product
		err := rows.Scan(&p.ID, &p.Name, &p.Stock, &p.Price, &p.CategoryID)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

func (r *ProductRepository) Create(product *models.Product) error {
	query := "INSERT INTO products (name, category_id, stock, price) VALUES ($1, $2, $3, $4) RETURNING id"
	err := r.db.QueryRow(query, product.Name, product.CategoryID, product.Stock, product.Price).Scan(&product.ID)
	return err
}

func (r *ProductRepository) GetByID(id int) (*models.Product, error){
	query := "SELECT id, name, stock, price, category_id FROM products WHERE id = $1"
	var p models.Product
	err := r.db.QueryRow(query, id).Scan(&p.ID, &p.Name,&p.Stock, &p.Price, &p.CategoryID)
	if err == sql.ErrNoRows{
		return nil, errors.New("Product not found")
	}
	if err != nil {
		return  nil, err
	}
	return &p, nil
}

func (r *ProductRepository) Update(product *models.Product) error{
	query := "UPDATE products SET name= $1, category_id=$2, stock=$3, price=$4 WHERE id = $5"
	result, err := r.db.Exec(query, product.Name, product.CategoryID, product.Stock, product.Price, product.ID)
	if err != nil{
		return  err
	}

	rows, err := result.RowsAffected()
	if err != nil{
		return  err
	}

	if rows == 0 {
		return  errors.New("Product Not Found")
	}
	return  nil
}

func (r *ProductRepository) Delete(id int) error{
	query := "DELETE FROM products WHERE id = $1"
	result, err := r.db.Exec(query, id)

	if err != nil {
		return  err
	}

	rows, err := result.RowsAffected()
	if err != nil{
		return  err
	}

	if rows == 0 {
		return  errors.New("Products Not Found")
	}
	return err
}