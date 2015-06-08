package assert

var panicOnFailure bool

func init() {
	panicOnFailure = false
}

func SetPanic(b bool) {
	panicOnFailure = b
}
