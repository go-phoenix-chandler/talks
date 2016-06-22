// +build OMIT

package sample

import "encoding/json"

//START OMIT
func LoadStruct(data []byte) (output JSONData) {
	json.Unmarshal(data, &output)
	return output
}

//END OMIT
