package testutils

type Failer interface {
	SetNextFail(err error)
	NextFail() error
}
