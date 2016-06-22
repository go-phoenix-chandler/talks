// +build OMIT

package sample

import "encoding/json"

//START OMIT

type JSONData struct {
	Name string          `json:"command"`
	Body json.RawMessage `json:"body"`
}

func LoadStruct(data []byte) (output JSONData) {
	json.Unmarshal(data, &output)
	return output
}

func LoadArray(data []byte) (output []JSONData)

//END OMIT
