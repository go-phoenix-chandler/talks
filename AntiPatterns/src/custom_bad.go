// +build OMIT

package sample

//START OMIT

type SpecialError struct { /* fields of error information */
}

func (e *SpecialError) Error() string { return /* error message */ }

var ErrBad = &SpecialError{}

func DoSomething() error {
	var specErr *SpecialError = nil
	if badStuffHappens {
		specErr = ErrBad
	}
	return specErr
}

//END OMIT
