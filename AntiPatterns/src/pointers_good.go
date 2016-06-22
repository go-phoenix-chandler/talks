// +build OMIT

package sample

//START OMIT
func MainFunc() string {
	return <-End(Start())
}

func Start() chan map[string]string {
	start := make(chan map[string]string)
	go func() {
		holder := map[string]string{}
		holder["foo"] = "bar"
		start <- holder
	}()
	return start
}

func End(data chan map[string]string) chan map[string]string {
	end := make(chan map[string]string)
	go func() {
		holder := <-data
		holder["foo"] = "something else"
		end <- holder
	}()
	return end
}

//END OMIT
