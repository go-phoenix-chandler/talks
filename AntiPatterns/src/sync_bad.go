// +build OMIT

package sample

//START OMIT
func MainFunc() string {
	holder := map[string]string{}
	go func() {
		holder["foo"] = "bar"
	}()

	go func() {
		holder["foo"] = "something else"
	}()
	return holder["foo"]
}

//END OMIT
