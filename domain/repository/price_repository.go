package repository

import (
	"go_stock_analysis/domain/model"
)

type PriceRepository interface {
	Create(price *model.Price) (*model.Price, error)
	FindByID(id int) (*model.Price, error)
	Update(price *model.Price) (*model.Price, error)
	Delete(price *model.Price) error
}
