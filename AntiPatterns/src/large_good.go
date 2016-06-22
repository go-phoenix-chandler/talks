// +build OMIT

package sample

//START OMIT
type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}

type Marshaler interface {
	MarshalJSON() ([]byte, error)
}

//END OMIT
