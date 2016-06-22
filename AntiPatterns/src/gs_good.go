// +build OMIT

package sample

import "fmt"

//START OMIT
type NameJSONNoExport struct {
	First string `json:"-"`
	Last  string `json:"-"`
}

type NonSimpleGS struct {
	first string
	last  string
}

func (n *NonSimpleGS) GetName() string {
	return fmt.Printf("%s %s", n.first, n.last)
}

//Even Better
func (n *NonSimpleGS) String() string {
	return fmt.Printf("%s %s", n.first, n.last)
}

func (n *NonSimpleGS) LoadFromJWT(token string) {
	//Do magic to load from some odd encrypted / encoded version of data
	//Verify add verification
}

//END OMIT
