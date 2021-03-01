package transaction

import (
	"bwa-startup/campaign"
	"errors"
)


type Service interface {
	GetTransactionByCampaignId(input GetTransactionByCampaignIdInput) ([]Transaction, error)
}

type service struct {
	repository        Repository
	campaignRepository campaign.Repository
}

func NewService(repository Repository, campaignRepository campaign.Repository) *service {
	return &service{repository, campaignRepository}
}

func (s *service) GetTransactionByCampaignId(input GetTransactionByCampaignIdInput) ([]Transaction, error) {
	campaign, err := s.campaignRepository.FindById(input.ID)
	if err != nil {
		return []Transaction{}, err
	}
	
	if campaign.UserID != input.User.ID {
		return []Transaction{}, errors.New("Not an owner of the campaign")
	}
	
	
	transactions, err := s.repository.GetByCampaignId(input.ID)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}