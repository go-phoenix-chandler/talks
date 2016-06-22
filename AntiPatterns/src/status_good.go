// +build OMIT

package sample

import (
	"errors"
	"time"
)

//START OMIT
var ErrNotStart = errors.New("Worker not started")

func (worker *Worker) Duration() (time.Duration, error) {
	if notStart {
		return 0, ErrNotStart
	}
	return time.Since(worker.start), nil
}

//END OMIT
