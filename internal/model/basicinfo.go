package model

type BasicInfo struct {
	id         uint64
	code       string
	creator    *User
	createDate string
}
