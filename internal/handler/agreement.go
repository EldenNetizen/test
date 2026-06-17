package handler

import (
	"github.com/EldenNetizen/test/internal/model"
	"github.com/EldenNetizen/test/internal/service"
)

type AgreementHandler struct {
	as *service.AgreementService
}

func NewAgreementHandler() *AgreementHandler {
	return &AgreementHandler{as: service.NewAgreementService()}
}

func (ah *AgreementHandler) CreateAgreement(agreement *model.Agreement) error {
	return ah.as.CreateAgreement(agreement)
}

func (ah *AgreementHandler) GetAgreement(id uint64) (*model.Agreement, error) {
	return ah.as.GetAgreement(id)
}

func (ah *AgreementHandler) DeleteAgreement(id uint64) (bool, error) {
	return ah.as.DeleteAgreement(id)
}
