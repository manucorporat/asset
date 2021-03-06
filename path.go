package asset

import (
	"path"
)

const (
	Environment = 1 << iota
	Flags       = 1 << iota
	NoInput     = 0
)

var (
	EnvironmentVariable = "ASSETS_PATH"
)

var _basePath string = ""

func init() {
	computeBasePath("")
}

func Config(defaultPath string, flags int) {
	value, ok := readValue(flags)
	if !ok {
		value = defaultPath
	}
	computeBasePath(value)
}

func Path(filename string) string {
	return path.Join(_basePath, filename)
}

func computeBasePath(aPath string) {
	if path.IsAbs(aPath) {
		_basePath = aPath
	} else {
		_basePath = path.Join(basepath(), aPath)
	}
}
