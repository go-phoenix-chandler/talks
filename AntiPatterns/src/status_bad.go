// +build OMIT

package sample

import "time"

//START OMIT
func (worker *Worker) Duration() time.Duration {
	if notStart {
		return -1
	}
	return time.Since(worker.start)
}

//END OMIT
