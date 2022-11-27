package user

import (
	"fmt"
	"nagase/doer"
	"time"
)

type User struct {
	Doer doer.Doer
	Calc doer.Calculator
}

var now = time.Now

func (u *User) Use() (string, error) {
	u.Calc.Add(12, 34)
	fmt.Println("🈲", now().Weekday())
	return u.Doer.DoSomething(123, "Hello GoMock")
}
