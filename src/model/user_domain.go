package model

import (
	"crypto/md5"
	"encoding/hex"
)

type UserDomainInterface interface {
	GetEmail() string
	GetName() string
	GetPassword() string
	GetAge() int
	GetID() string
	SetID(string)
	EncryptPassword()
}
type userDomain struct {
	id       string `json:"id"`
	email    string `json:"email"`
	password string `json:"password"`
	name     string `json:"name"`
	age      int    `json:"age"`
}

func NewUserDomain(
	email, password, name string,
	age int,
) UserDomainInterface {
	return &userDomain{
		email:    email,
		name:     name,
		password: password,
		age:      age,
	}
}

func (ud *userDomain) SetID(id string)     { ud.id = id }
func (ud *userDomain) GetID() string       { return ud.id }
func (ud *userDomain) GetEmail() string    { return ud.email }
func (ud *userDomain) GetName() string     { return ud.name }
func (ud *userDomain) GetPassword() string { return ud.password }
func (ud *userDomain) GetAge() int         { return ud.age }
func (ud *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.password))
	ud.password = hex.EncodeToString(hash.Sum(nil))
}
