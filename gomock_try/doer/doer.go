package doer

//go:generate mockgen -destination=../mocks/mock_doer.go -package=mocks nagase/doer Doer

type Doer interface {
	DoSomething(int, string) (string, error)
}



//go:generate mockgen -destination=../mocks/mock_fs.go -package=mocks github.com/spf13/afero Fs
