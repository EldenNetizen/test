package model

import "encoding/json"

type User struct {
	id         uint64
	name       string
	department *Department
	company    *Company
}

func NewUser(id uint64, name string, department *Department, company *Company) *User {
	return &User{
		id:         id,
		name:       name,
		department: department,
		company:    company,
	}
}

func (user *User) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id":         user.id,
		"name":       user.name,
		"department": user.department,
		"company":    user.company,
	})
}

func (user *User) UnmarshalJSON(data []byte) error {
	var aux struct {
		ID         uint64      `json:"id"`
		Name       string      `json:"name"`
		Department *Department `json:"department"`
		Company    *Company    `json:"company"`
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	user.id = aux.ID
	user.name = aux.Name
	user.department = aux.Department
	user.company = aux.Company
	return nil
}
