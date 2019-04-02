package must

import "log"

func Do(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func DoF(f func() error) {
	Do(f())
}
