package repository

import (
	"strconv"

	"github.com/EldenNetizen/test/internal/model"
	"github.com/EldenNetizen/test/pkg/database"
	"github.com/EldenNetizen/test/pkg/utils"
)

type AgreementRepository struct {
	sf *utils.SnowflakeIDGenerator
	rm *database.RedisManager
}

func NewAgreementRepository() *AgreementRepository {
	ar := &AgreementRepository{
		sf: utils.NewSnowflakeIDGenerator(1),
		rm: database.NewRedisManager(),
	}
	ar.rm.Connect()
	return ar
}

func (ar *AgreementRepository) CreateAgreement(agreement *model.Agreement) error {
	agreementJSON, err := agreement.MarshalJSON()
	if err != nil {
		return err
	}
	return ar.rm.Set(strconv.FormatUint(agreement.GetId(), 10), string(agreementJSON), 0)
}

func (ar *AgreementRepository) GetAgreement(key uint64) (*model.Agreement, error) {
	agreementJSON, err := ar.rm.Get(strconv.FormatUint(key, 10))
	if err != nil {
		return nil, err
	}
	agreement := &model.Agreement{}
	if err := agreement.UnmarshalJSON([]byte(agreementJSON)); err != nil {
		return nil, err
	}
	return agreement, nil
}

func (ar *AgreementRepository) DeleteAgreement(key uint64) (int64, error) {
	return ar.rm.Delete(strconv.FormatUint(key, 10))
}
