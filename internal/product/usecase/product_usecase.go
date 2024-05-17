package usecase

import (
	"eulabs/internal/entity"
	"eulabs/internal/product/repository"
)

type ProductUsecase struct {
	productRepo repository.ProductRepository
}

func NewProductUsecase(r repository.ProductRepository) *ProductUsecase {
	return &ProductUsecase{r}
}

func (u *ProductUsecase) CreateProduct(product *entity.Product) (*entity.Product, error) {
	return u.productRepo.Create(product)
}

func (u *ProductUsecase) FetchProducts() ([]*entity.Product, error) {
	p, err := u.productRepo.Fetch()
	if err != nil {
		return nil, err
	}
	if len(p) == 0 {
		return make([]*entity.Product, 0), nil
	}
	return p, nil
}

func (u *ProductUsecase) GetProductByID(id int) (*entity.Product, error) {
	return u.productRepo.GetByID(id)
}

func (u *ProductUsecase) UpdateProduct(product *entity.Product) error {
	return u.productRepo.Update(product)
}

func (u *ProductUsecase) DeleteProduct(id int) error {
	return u.productRepo.Delete(id)
}
