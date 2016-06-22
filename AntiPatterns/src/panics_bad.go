// +build OMIT

package sample

//START OMIT
func DoSomethingThatMightPanic() (err error) {
	err = externalpkg.OMGPanic()
	return err
}

//END OMIT
