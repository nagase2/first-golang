package user

import "nagase/mock/doer"

type User struct {
	Doer doer.Doer
}

func (u *User) Use() (string, error) {
	return u.Doer.DoSomething(123, "Hello GoMock")
}
