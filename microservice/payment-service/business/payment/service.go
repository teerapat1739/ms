package payment

// Service - interface
type Service interface {
	Payment(email, id string) error
}
