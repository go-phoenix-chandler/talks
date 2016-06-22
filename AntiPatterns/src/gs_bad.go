// +build OMIT

package sample

//START OMIT
type name struct {
	first string
	last  string
}

func (n *name) GetFirstName() string {
	return n.first
}

func (n *name) GetLastName() string {
	return n.last
}

func (n *name) SetFirstName(value string) {
	n.first = value
}

func (n *name) SetLastName(value string) {
	n.last = value
}

//END OMIT
