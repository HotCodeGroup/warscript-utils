package testutils

type Failer struct {
	nextFail error
}

func (f *Failer) SetNextFail(err error) {
	f.nextFail = err
}

func (f *Failer) NextFail() error {
	err := f.nextFail
	f.nextFail = nil
	return err
}
