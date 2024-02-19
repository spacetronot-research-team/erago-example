package funcs

import "runtime"

// GetMyName return string func name.
func GetMyName() string {
	pc, _, _, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)

	if ok && details != nil {
		return details.Name()
	}

	return "err get my func name"
}
