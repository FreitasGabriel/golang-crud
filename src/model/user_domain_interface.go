package model

type UserDomainInterface interface {
	GetEmail() string
	GetName() string
	GetPassword() string
	GetAge() int
	GetID() string
	SetID(string)
	EncryptPassword()
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

func NewUserUpdateDomain(
	name string,
	age int,
) UserDomainInterface {
	return &userDomain{
		name: name,
		age:  age,
	}
}
