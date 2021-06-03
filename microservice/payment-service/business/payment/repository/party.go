package repository

import (
	"fmt"
	"log"
	"payment-service/business/payment"

	"github.com/go-redis/redis"
)

type paymentRepository struct {
	client *redis.Client
}

// NewRepository - init user repository
func NewRepository(client *redis.Client) payment.Repository {
	return &paymentRepository{
		client: client,
	}
}

// Update updates an kudo.
func (r *paymentRepository) Payment(email, id string) error {
	msg := fmt.Sprintf("คุณ %s ได้ทำการจ่ายเงิน %s แล้ว", email, id)
	err := r.client.Publish("testttt", msg)
	log.Println("err", err)
	return err.Err()
}
