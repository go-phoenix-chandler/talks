// +build OMIT

package sample

import "encoding/json"

//START OMIT
func LoadStruct(data []byte) (output JSONData, err error) {
	err = json.Unmarshal(data, &output)
	return output, err
}

//END OMIT
