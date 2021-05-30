package service

import (
	"party-service/business/party"
	"party-service/business/party/model"
)

type partyService struct {
	userId string
	repo   party.Repository
}

// userService - init user service
func NewUserService(partyRepo party.Repository, userId string) party.Service {
	return &partyService{
		userId: userId,
		repo:   partyRepo,
	}
}

func (s *partyService) GetKudos() ([]model.Party, error) {
	return s.repo.FindAll(map[string]interface{}{"userId": s.userId})
}

func (s *partyService) CreateParty(party model.Party) (model.Party, error) {
	var p model.Party
	err := s.repo.Create(party)
	if err != nil {
		return p, err
	}
	return p, nil
}

func (s *partyService) JoinParty(m model.Member) error {
	return s.repo.JoinParty(m)
}

func (s *partyService) RemoveKudo(githubRepo model.Party) (model.Party, error) {
	kudo := s.githubRepoToKudo(githubRepo)
	err := s.repo.Delete(kudo)
	if err != nil {
		return kudo, err
	}
	return kudo, nil
}

func (s *partyService) githubRepoToKudo(githubRepo model.Party) model.Party {
	return model.Party{}
}
