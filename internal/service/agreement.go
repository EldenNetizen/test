package service

import (
	"github.com/EldenNetizen/test/internal/model"
	"github.com/EldenNetizen/test/internal/repository"
)

type AgreementService struct {
	ar *repository.AgreementRepository
}

func NewAgreementService() *AgreementService {
	return &AgreementService{
		ar: repository.NewAgreementRepository(),
	}
}

func (as *AgreementService) CreateAgreement(agreement *model.Agreement) error {
	if as.IsCreateAllowed(agreement) {
		return as.ar.CreateAgreement(agreement)
	}
	return nil
}

func (as *AgreementService) GetAgreement(id uint64) (*model.Agreement, error) {
	if !as.IsGetAllowed(id) {
		return nil, nil
	}
	agreement, err := as.ar.GetAgreement(id)
	if err != nil {
		return nil, err
	}
	return agreement, nil
}

func (as *AgreementService) DeleteAgreement(id uint64) (bool, error) {
	if !as.IsDeleteAllowed(id) {
		return false, nil
	}
	_, err := as.ar.DeleteAgreement(id)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (as *AgreementService) IsCreateAllowed(agreement *model.Agreement) bool {
	return true
}

func (as *AgreementService) IsDeleteAllowed(id uint64) bool {
	return true
}

func (as *AgreementService) IsGetAllowed(id uint64) bool {
	return true
}
