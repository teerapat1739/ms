package party

import "party-service/business/party/model"

// Service - interface
type Service interface {
	GetKudos() ([]model.Party, error)
	CreateParty(githubRepo model.Party) (model.Party, error)
}
