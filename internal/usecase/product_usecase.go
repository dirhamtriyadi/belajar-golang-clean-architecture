package usecase

import (
	"github.com/dirhamtriyadi/belajar-golang-clean-architecture/internal/entity"
	"github.com/dirhamtriyadi/belajar-golang-clean-architecture/internal/repository"
)

type ProductUsecase interface {
	GetAllProducts() ([]*entity.Product, error)
	GetProductByID(id uint64) (*entity.Product, error)
	CreateProduct(product *entity.Product) (*entity.Product, error)
	UpdateProduct(product *entity.Product) (*entity.Product, error)
	DeleteProduct(id uint64) error
}

type productUsecase struct {
	productRepo repository.ProductRepository
}

func (u *productUsecase) GetAllProducts() ([]*entity.Product, error) {
	return u.productRepo.FindAll()
}

func (u *productUsecase) GetProductByID(id uint64) (*entity.Product, error) {
	return u.productRepo.FindByID(id)
}

func (u *productUsecase) CreateProduct(product *entity.Product) (*entity.Product, error) {
	return u.productRepo.Create(product)
}

func (u *productUsecase) UpdateProduct(product *entity.Product) (*entity.Product, error) {
	return u.productRepo.Update(product)
}

func (u *productUsecase) DeleteProduct(id uint64) error {
	return u.productRepo.Delete(id)
}

func NewProductUsecase(repo repository.ProductRepository) ProductUsecase {
	return &productUsecase{
		productRepo: repo,
	}
}
