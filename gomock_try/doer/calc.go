package doer



//go:generate mockgen -destination=../mocks/mock_calc.go -package=mocks nagase/doer Calculator

type Calculator interface {
	DoSomething(int, string) (string, error)
	Add(int, int) (int, error)
}
