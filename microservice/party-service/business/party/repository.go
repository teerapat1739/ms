package party

import "party-service/business/party/model"

// Repository - interface
type Repository interface {
	Find(id string) (model.Party, error)
	FindAll(selector map[string]interface{}) ([]model.Party, error)
	Delete(kudo model.Party) error
	Update(kudo model.Party) error
	Create(party model.Party) error
	JoinParty(m model.Member) error
}
