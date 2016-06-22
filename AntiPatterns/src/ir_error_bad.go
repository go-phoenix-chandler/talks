// +build OMIT

package sample

import "strings"

//START OMIT

func BadFunc() error {
	return error.New("bad things happened")
}

func CallBadFunc() error {
	// If this is all your doing it works
	if err := BadFunc(); err != nil {
		return err
	}
}

func ButIfNeedMore() error {
	//This is slow and only works if you do not change things
	if err := BadFunc(); strings.Contains(err.Error(), "I know have to look for your text") {
		//do something else
	}
}

//END OMIT
