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
type RealType struct {
	Name  string `json:"name"`
	Value string `json:"doesnothavetomatch"`
}

func LoadJSON(data []byte) (output RealType) {
	json.Unmarshal(data, &output)
	return output
}

type DesiredInterface interface {
	DoSomething()
}

func MultiHandling(anthing DesiredInterface) {
	anything.DoSomething()
}

//END OMIT
