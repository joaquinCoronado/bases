package errors

import "fmt"

type BasesError struct {
	Number string
	Base int
	ValidSymbols []string
}

func (e *BasesError) Error() string{
	return fmt.Sprintf(
		"\n\n %v is not valid to base %v \n\n Valid symbols to base %v: %v \n\n",
		e.Number,
		e.Base,
		e.Base,
		e.ValidSymbols,
	)
}
