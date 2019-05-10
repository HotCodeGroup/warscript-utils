package testutils

type Failer interface {
	SetNextFail()
	NextFail()
}
