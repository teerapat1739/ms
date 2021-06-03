package service

import (
	"payment-service/business/payment"
)

type paymentService struct {
	userId string
	repo   payment.Repository
}

// userService - init user service
func NewUserService(paymentRepo payment.Repository, userId string) payment.Service {
	return &paymentService{
		userId: userId,
		repo:   paymentRepo,
	}
}

func (s *paymentService) Payment(email, id string) error {
	return s.repo.Payment(email, id)
}
