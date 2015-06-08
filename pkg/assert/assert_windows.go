package assert

import "github.com/Sirupsen/logrus"

func Assert(expr bool, e string) {
	if expr == false {
		if panicOnFailure {
			panic(e)
		} else {
			logrus.Errorf("**ASSERTION FAILED**: %s", e)
		}
	}
}
