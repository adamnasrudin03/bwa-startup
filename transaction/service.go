package transaction

type Service interface {
	GetTransactionByCampaignId(input GetTransactionByCampaignIdInput) ([]Transaction, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetTransactionByCampaignId(input GetTransactionByCampaignIdInput) ([]Transaction, error) {
	transactions, err := s.repository.GetCampaignById(input.ID)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}