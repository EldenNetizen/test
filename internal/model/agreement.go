package model

import "encoding/json"

type Agreement struct {
	ID                 uint64 `db:"id"`
	AgreementAmountSum string `db:"agreementAmountSum"`
}

type AgreementOptions func(agreement *Agreement)

func WithId(id uint64) AgreementOptions {
	return func(agreement *Agreement) {
		agreement.ID = id
	}
}

func WithAgreementAmountSum(agreementAmountSum string) AgreementOptions {
	return func(agreement *Agreement) {
		agreement.AgreementAmountSum = agreementAmountSum
	}
}

func NewAgreement(options ...AgreementOptions) *Agreement {
	agreement := &Agreement{}
	for _, option := range options {
		option(agreement)
	}
	return agreement
}

func (agreement *Agreement) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id":                 agreement.GetId(),
		"agreementAmountSum": agreement.GetAgreementAmountSum(),
	})
}

func (agreement *Agreement) UnmarshalJSON(data []byte) error {
	var aux struct {
		ID                 uint64 `json:"id"`
		AgreementAmountSum string `json:"agreementAmountSum"`
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	agreement.ID = aux.ID
	agreement.AgreementAmountSum = aux.AgreementAmountSum
	return nil
}

func (agreement *Agreement) GetId() uint64 {
	return agreement.ID
}

func (agreement *Agreement) GetAgreementAmountSum() string {
	return agreement.AgreementAmountSum
}
