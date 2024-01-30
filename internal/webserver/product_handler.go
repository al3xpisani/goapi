package webserver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/al3xpisani/goapi/internal/entity"
	"github.com/al3xpisani/goapi/internal/service"
	"github.com/go-chi/chi/v5"
)

type WebProductHandler struct {
	ProductService *service.ProductService
}

func NewWebProductHandler(ProductService *service.ProductService) *WebProductHandler {
	return &WebProductHandler{ProductService: ProductService}
}

func (wch *WebProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := wch.ProductService.GetProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(products)
}

func (wch *WebProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Id is required", http.StatusBadRequest)
		return
	}
	product, err := wch.ProductService.GetProduct(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(product)
}

func (wch *WebProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product entity.Product
	fmt.Println("******************************************", json.NewDecoder(r.Body).Decode(&product))
	err := json.NewDecoder(r.Body).Decode(&product)
	if err == nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := wch.ProductService.CreateProduct(product.Name, product.Description, product.CategoryID, product.Image_url, product.Price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(result)
}

func (wch *WebProductHandler) GetProductByCategoryID(w http.ResponseWriter, r *http.Request) {
	categoryID := chi.URLParam(r, "categoryID")
	if categoryID == "" {
		http.Error(w, "categoryID is required", http.StatusBadRequest)
		return
	}
	product, err := wch.ProductService.GetProductByCategoryID(categoryID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(product)
}
