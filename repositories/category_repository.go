package repositories

import (
	"database/sql"
	"errors"
	"gokasir-api/models"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) GetAll() ([]models.Category, error) {
	query := "SELECT id, name FROM categories"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := make([]models.Category, 0)
	for rows.Next() {
		var c models.Category
		err := rows.Scan(&c.ID, &c.Name)
		if err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}

	return categories, nil
}

func (r *CategoryRepository) Create(category *models.Category) error {
	query := "INSERT INTO categories (name) VALUES ($1) RETURNING id"
	err := r.db.QueryRow(query, category.Name).Scan(&category.ID)
	return err
}

func (r *CategoryRepository) GetByID(id int) (*models.Category, error){
	query := "SELECT id, name FROM categories WHERE id = $1"
	var c models.Category
	err := r.db.QueryRow(query, id).Scan(&c.ID, &c.Name)
	if err == sql.ErrNoRows{
		return nil, errors.New("Category not found")
	}
	if err != nil {
		return  nil, err
	}
	return &c, nil
}

func (r *CategoryRepository) Update(category *models.Category) error{
	query := "UPDATE categories SET name= $1 WHERE id = $2"
	result, err := r.db.Exec(query, category.Name, category.ID)
	if err != nil{
		return  err
	}

	rows, err := result.RowsAffected()
	if err != nil{
		return  err
	}

	if rows == 0 {
		return  errors.New("Category Not Found")
	}
	return  nil
}

func (r *CategoryRepository) Delete(id int) error{
	query := "DELETE FROM categories WHERE id = $1"
	result, err := r.db.Exec(query, id)

	if err != nil {
		return  err
	}

	rows, err := result.RowsAffected()
	if err != nil{
		return  err
	}

	if rows == 0 {
		return  errors.New("Category Not Found")
	}
	return err
}