// +build OMIT

package sample

import "encoding/json"

type TypeOne struct {
	Name string
}

type TypeTwo struct {
	Value string
}

//START OMIT
func LoadJSON(data []byte) (output interface{}) {
	json.Unmarshal(data, &output)
	return output
}

func MultiHandling(anthing interface{}) {
	switch v := anything.(type) {
	case TypeOne:
		//Do Something
	case TypeTwo:
		//Do Something
	}
}

//END OMIT
