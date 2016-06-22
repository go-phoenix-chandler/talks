// +build OMIT

package sample

//START OMIT

type SpecialError struct { /* fields of error information */
}

func (e *SpecialError) Error() string { return /* error message */ }

var ErrBad = &SpecialError{}

func DoSomething() error {
	if badStuffHappens {
		return ErrBad
	}
	return nil
}

//END OMIT
