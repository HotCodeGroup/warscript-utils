package testutils

// Failer возвращает установленную заранее ошибку
type Failer struct {
	nextFail error
}

// SetNextFail установка ошибки
func (f *Failer) SetNextFail(err error) {
	f.nextFail = err
}

// NextFail получение текущей ошибки
func (f *Failer) NextFail() error {
	err := f.nextFail
	f.nextFail = nil
	return err
}
