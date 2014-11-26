package presentation

import (
	"fmt"
	"time"
)

type Response struct {
	Id   int       `json:"some_file"` // This will be caught while linting for the name
	name string    `json:"name"`      // These comments will be formatted as well
	date time.Time `json:"date"`
}

func UnformattedExample() {
	fmt.Printf("%d", 8+9)
}
