package model

import "encoding/json"

type Agreement struct {
	id                 uint64
	agreementAmountSum string
}

type AgreementOptions func(agreement *Agreement)

func WithId(id uint64) AgreementOptions {
	return func(agreement *Agreement) {
		agreement.id = id
	}
}

func WithAgreementAmountSum(agreementAmountSum string) AgreementOptions {
	return func(agreement *Agreement) {
		agreement.agreementAmountSum = agreementAmountSum
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
	agreement.id = aux.ID
	agreement.agreementAmountSum = aux.AgreementAmountSum
	return nil
}

func (agreement *Agreement) GetId() uint64 {
	return agreement.id
}

func (agreement *Agreement) GetAgreementAmountSum() string {
	return agreement.agreementAmountSum
}
