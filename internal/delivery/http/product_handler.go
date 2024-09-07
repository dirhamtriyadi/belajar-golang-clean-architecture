package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dirhamtriyadi/belajar-golang-clean-architecture/internal/entity"
	"github.com/dirhamtriyadi/belajar-golang-clean-architecture/internal/usecase"

	"github.com/gorilla/mux"
)

type ProductHandler struct {
	productUsecase usecase.ProductUsecase
}

func NewProductHandler(router *mux.Router, usecase usecase.ProductUsecase) {
	handler := &ProductHandler{
		productUsecase: usecase,
	}

	router.HandleFunc("/products", handler.GetAllProducts).Methods("GET")
	router.HandleFunc("/products/{id}", handler.GetProductByID).Methods("GET")
	router.HandleFunc("/products", handler.CreateProduct).Methods("POST")
	router.HandleFunc("/products/{id}", handler.UpdateProduct).Methods("PUT")
	router.HandleFunc("/products/{id}", handler.DeleteProduct).Methods("DELETE")
}

func (h *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.productUsecase.GetAllProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(products)
}

func (h *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 64)
	product, err := h.productUsecase.GetProductByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product entity.Product
	json.NewDecoder(r.Body).Decode(&product)
	createdProduct, err := h.productUsecase.CreateProduct(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(createdProduct)
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var product entity.Product
	json.NewDecoder(r.Body).Decode(&product)
	updatedProduct, err := h.productUsecase.UpdateProduct(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(updatedProduct)
}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 64)
	err := h.productUsecase.DeleteProduct(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
