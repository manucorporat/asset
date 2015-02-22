package asset

import (
	"fmt"
	"github.com/mitchellh/osext"
	"path"
	"runtime"
	"strings"
)

func basepath() string {
	absolutePath := executablePath()
	if strings.HasSuffix(absolutePath, "command-line-arguments/_obj/exe/") {
		return callerPath()
	} else {
		return absolutePath
	}
}

func executablePath() string {
	absP, err := osext.ExecutableFolder()
	if err != nil {
		panic(err.Error())
	}
	return absP
}

func callerPath() string {
	var index int = 2
	for {
		pc, filename, _, ok := runtime.Caller(index)
		if ok {
			function := runtime.FuncForPC(pc)
			fmt.Println(function.Name())
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
