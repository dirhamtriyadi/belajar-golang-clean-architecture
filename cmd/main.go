package main

import (
	"log"
	"net/http"

	httpHandler "github.com/dirhamtriyadi/belajar-golang-clean-architecture/internal/delivery/http"
	"github.com/dirhamtriyadi/belajar-golang-clean-architecture/internal/infra/sqlite"
	"github.com/dirhamtriyadi/belajar-golang-clean-architecture/internal/repository"
	"github.com/dirhamtriyadi/belajar-golang-clean-architecture/internal/usecase"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	db, err := sqlite.InitDB()

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	productRepo := repository.NewProductRepository(db)
	productUsecase := usecase.NewProductUsecase(productRepo)
	httpHandler.NewProductHandler(router, productUsecase)

	log.Println("Server started at :8080")
	http.ListenAndServe(":8080", router)
}
