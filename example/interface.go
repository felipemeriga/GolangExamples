package example

type AddHeadersFunc func() error

func (m AddHeadersFunc) AddHeader() error {
	return m()
}

type AddHeadersInterface interface {
	AddHeader() error
}

type Pattern struct {
	AddHeadersInterface
	Name string
}

type Extended struct {
	ID string
	*Pattern
}

func NewExtended() *Extended {
	extended := &Extended{
		ID: "",
		Pattern: &Pattern{
			AddHeadersInterface: nil,
			Name:                "",
		},
	}
	extended.AddHeadersInterface = AddHeadersFunc(
		func() error {
			return nil
		},
	)

	return extended
}
