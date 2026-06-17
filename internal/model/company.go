package model

import "encoding/json"

type Company struct {
	id   uint64
	name string
}

func NewCompany(id uint64, name string) *Company {
	return &Company{
		id:   id,
		name: name,
	}
}

func (company *Company) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id":   company.id,
		"name": company.name,
	})
}

func (company *Company) UnmarshalJSON(data []byte) error {
	var aux struct {
		ID   uint64 `json:"id"`
		Name string `json:"name"`
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	company.id = aux.ID
	company.name = aux.Name
	return nil
}
