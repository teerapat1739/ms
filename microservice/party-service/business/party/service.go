package party

import (
	"party-service/business/party/model"
)

// Service - interface
type Service interface {
	GetParty() ([]model.Party, error)
	GetPartyByID(id int) (model.Party, error)
	CreateParty(githubRepo model.Party) (model.Party, error)
	JoinParty(m model.Member) error
}
