// +build OMIT

package sample

import "encoding/json"

//START OMIT
func LoadStruct(data []byte) (output map[string]interface{}) {
	json.Unmarshal(data, &output)
	return output
}

func LoadArray(data []byte) (output []interface{}) {
	json.Unmarshal(data, &output)
	return output
}

//END OMIT
