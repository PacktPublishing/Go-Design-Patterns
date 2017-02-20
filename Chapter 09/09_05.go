package future

type SuccessFunc func(string)
type FailFunc func(error)
type ExecuteStringFunc func() (string, error)

type MaybeString struct {
}

func (s *MaybeString) Success(f SuccessFunc) *MaybeString {
	return nil
}

func (s *MaybeString) Fail(f FailFunc) *MaybeString {
	return nil
}

func (s *MaybeString) Execute(f ExecuteStringFunc) {
}
