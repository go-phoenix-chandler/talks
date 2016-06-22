// +build OMIT

package sample

import "sync"

//START OMIT
type LockedHolder struct {
	sync.Locker
	Data map[string]string
}

func MainFunc() string {
	holder := &LockedHolder{Data: map[string]string{}}
	wg := &sync.WaitGroup{}
	wg.Add(2)

	lockHelper(holder, wg, func() { holder["foo"] = "bar" })
	lockHelper(holder, wg, func() { holder["foo"] = "something else" })

	wg.Wait()
	return holder["foo"]
}
func lockHelper(holder *LockerHolder, wg *sync.WaitGroup, work func()) {
	go func() {
		holder.Lock()
		defer holder.Lock()
		defer wg.Done()
		work()
	}()
}

//END OMIT
