package depteststringer

type Stringer string

func NewStringer(s string) Stringer {
	return Stringer(s)
}

func (s Stringer) Get() string {
	return string(s)
}
