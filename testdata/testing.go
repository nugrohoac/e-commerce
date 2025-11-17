package testdata

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"testing"
)

// basepath is the root directory of this package
var basepath string

// FuncCaller form of function's call which includes expected input and output
type FuncCaller struct {
	IsCalled bool
	Input    []interface{}
	Output   []interface{}
}

func init() {
	_, currentFile, _, _ := runtime.Caller(0)
	basepath = filepath.Dir(currentFile)
}

func path(relPath string) string {
	if filepath.IsAbs(relPath) {
		return relPath
	}

	return filepath.Join(basepath, relPath)
}

// GetGolden is a function to get golden file
func GetGolden(t *testing.T, filename string) []byte {
	t.Helper()

	b, err := ioutil.ReadFile(path(filename + ".golden"))
	if err != nil {
		t.Fatal(t)
	}

	return b
}

// GoldenJSONUnmarshal read golden file and calling json.Unmarshal
func GoldenJSONUnmarshal(t *testing.T, filename string, input interface{}) {
	_bytes := GetGolden(t, filename)

	err := json.Unmarshal(_bytes, &input)

	if err != nil {
		t.Fatal(t)
	}
}
