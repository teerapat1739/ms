package service

import (
	"errors"
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

func (s *partyService) GetParty() ([]model.Party, error) {
	res, err := s.repo.GetParty()
	if err != nil {
		return res, err
	}
	for i := 0; i < len(res); i++ {
		res[i].FillDefaultParty()
	}
	return res, nil
}
func (s *partyService) GetPartyByID(id int) (model.Party, error) {
	res, err := s.repo.GetPartyByID(id)
	if err != nil {
		return res, err
	}
	res.FillDefaultParty()
	return res, nil
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
	if err := s.isAvailiableToJoin(m.PartyID); err != nil {

		return err
	}

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

func (s *partyService) isAvailiableToJoin(partyId int) error {
	count, err := s.repo.CountMemberByPartyID(partyId)
	if err != nil {
		return err
	}
	res, err := s.repo.GetPartyByID(partyId)
	if err != nil {
		return err
	}
	if int(count) <= res.Size {
		return nil
	}
	return errors.New("ปาร์ตี้นี้เต็มแล้ว")
}
