package payment

// Repository - interface
type Repository interface {
	Payment(email, id string) error
}
