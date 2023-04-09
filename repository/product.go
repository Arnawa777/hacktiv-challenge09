package repository

import "challenge-08/models"

type ProductRepository interface {
	FindByID(id uint) *models.Product
}
