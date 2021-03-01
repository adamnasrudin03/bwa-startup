package transaction

import (
	"bwa-startup/campaign"
	"errors"
)


type Service interface {
	GetTransactionByCampaignId(input GetTransactionByCampaignIdInput) ([]Transaction, error)
	GetTransactionByUserId(userID int) ([]Transaction, error)
	CreateTransaction(input CreateTransactionInput) (Transaction, error)
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

func (s *service)GetTransactionByUserId(userID int) ([]Transaction, error) {
	transactions, err := s.repository.GetByUserId(userID)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s *service) CreateTransaction(input CreateTransactionInput) (Transaction, error) {
	transaction := Transaction{}
	transaction.Amount = input.Amount
	transaction.CampaignID = input.CampaignID
	transaction.UserID = input.User.ID
	transaction.Status = "pending"

	newTransaction, err := s.repository.Save(transaction)
	if err != nil {
		return newTransaction, err
	}

	return newTransaction, nil
}
