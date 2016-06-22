// +build OMIT

package sample

//START OMIT
func DoSomethingThatMightPanic() (err error) {
	defer func() {
		if perr := recover(); perr != nil {
			err = perr
		}
	}()
	err = externalpkg.OMGPanic()
	return err
}

//END OMIT
