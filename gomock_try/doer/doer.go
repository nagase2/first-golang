package doer

//go:generate mockgen -destination=../mocks/mock_doer.go -package=mocks nagase/mock/doer Doer

type Doer interface {
	DoSomething(int, string) (string, error)
}
