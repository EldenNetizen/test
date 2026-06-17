package model

import "encoding/json"

type Department struct {
	id   uint64
	name string
}

func NewDepartment(id uint64, name string) *Department {
	return &Department{
		id:   id,
		name: name,
	}
}

func (department *Department) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id":   department.id,
		"name": department.name,
	})
}

func (department *Department) UnmarshalJSON(data []byte) error {
	var aux struct {
		ID   uint64 `json:"id"`
		Name string `json:"name"`
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	department.id = aux.ID
	department.name = aux.Name
	return nil
}
