// +build OMIT

package sample

//START OMIT

var ErrBad = error.New("bad things happened")

func BadFunc() error {
	return ErrBad
}

func ButIfNeedMore() error {
	//This is slow and only works if you do not change things
	err := BadFunc()
	switch err {
	case ErrBad:
		//Handle It
	default:
		if err != nil { //Unhandled Error
			return err
		}
	}
}

//END OMIT
