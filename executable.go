package asset

import (
	"path"
	"runtime"
)

func executablePath() string {
	var index int = 2
	for {
		pc, filename, _, ok := runtime.Caller(index)
		if ok {
			function := runtime.FuncForPC(pc)
			if function.Name() == "main.init" {
				return path.Dir(filename)
			}
		} else {
			break
		}
		index++
	}
	return ""
}
