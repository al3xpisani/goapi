package database

import (
	"database/sql"

	"github.com/al3xpisani/goapi/internal/entity"
)

type ProductDB struct {
	db *sql.DB
}

func NewProductDB(db *sql.DB) *ProductDB {
	return &ProductDB{db: db}
}

func (cd *ProductDB) GetProducts() ([]*entity.Product, error) {
	rows, err := cd.db.Query("SELECT id, name, price, category_id, image_url, description from products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*entity.Product
	for rows.Next() {
		var product entity.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.CategoryID, &product.Image_url, &product.Description); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}

func (cd *ProductDB) GetProduct(id string) (*entity.Product, error) {
	var product entity.Product
	err := cd.db.QueryRow("SELECT id, name, price, category_id, image_url, description FROM products WHERE id = ? ", id).Scan(&product.ID, &product.Name, &product.Price, &product.CategoryID, &product.Image_url, &product.Description)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (cd *ProductDB) CreateProduct(product *entity.Product) (string, error) {
	_, err := cd.db.Exec("INSERT INTO products (id, name, price, category_id, image_url, description) VALUES (?,?,?,?,?,?)", product.ID, product.Name, product.Price, product.CategoryID, product.Image_url, product.Description)
	if err != nil {
		return "", err
	}
	return product.ID, nil
}

func (pd *ProductDB) GetProductsByCategoryID(categoryID string) ([]*entity.Product, error) {
	rows, err := pd.db.Query("select id, name, price, category_id, image_url, description from products where category_id = ? ", categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []*entity.Product
	for rows.Next() {
		var product entity.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.CategoryID, &product.Image_url, &product.Description); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}
