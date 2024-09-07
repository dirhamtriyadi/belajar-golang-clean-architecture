package repository

import (
	"database/sql"

	"github.com/dirhamtriyadi/belajar-golang-clean-architecture/internal/entity"

	_ "github.com/mattn/go-sqlite3"
)

type ProductRepository interface {
	FindAll() ([]*entity.Product, error)
	FindByID(id uint64) (*entity.Product, error)
	Create(product *entity.Product) (*entity.Product, error)
	Update(product *entity.Product) (*entity.Product, error)
	Delete(id uint64) error
}

type productRepository struct {
	db *sql.DB
}

func (r *productRepository) FindAll() ([]*entity.Product, error) {
	rows, err := r.db.Query("SELECT id, name, price FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*entity.Product
	for rows.Next() {
		var product entity.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Price); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}

func (r *productRepository) FindByID(id uint64) (*entity.Product, error) {
	row := r.db.QueryRow("SELECT id, name, price FROM products WHERE id = ?", id)
	var product entity.Product
	if err := row.Scan(&product.ID, &product.Name, &product.Price); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) Create(product *entity.Product) (*entity.Product, error) {
	result, err := r.db.Exec("INSERT INTO products (name, price) VALUES (?, ?)", product.Name, product.Price)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	product.ID = uint64(id)
	return product, nil
}

func (r *productRepository) Update(product *entity.Product) (*entity.Product, error) {
	_, err := r.db.Exec("UPDATE products SET name = ?, price = ? WHERE id = ?", product.Name, product.Price, product.ID)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *productRepository) Delete(id uint64) error {
	_, err := r.db.Exec("DELETE FROM products WHERE id = ?", id)
	return err
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./products.db")
	if err != nil {
		return nil, err
	}

	query := `
    CREATE TABLE IF NOT EXISTS products (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT,
        price REAL
    );
    `
	_, err = db.Exec(query)
	if err != nil {
		return nil, err
	}

	return db, nil
}
