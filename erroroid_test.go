package erroroid_test

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/sarpdoruk/erroroid"
	"log"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"testing"
)

const FileName = "erroroid_test.go"

func caller() (uintptr, string, int, bool) {
	return runtime.Caller(1)
}

func TestError(t *testing.T) {
	t.Run("error strings should be equal", func(t *testing.T) {
		filePath, _ := filepath.Abs(FileName)
		errStr := "default error"
		eoid := erroroid.NewErroroid()
		got := eoid.Error(errStr).Error()
		_, _, lnNum, _ := caller()
		want := "ERROR: [" + filePath + ":" + strconv.Itoa(lnNum-1) + " @TestError] -> " + errStr

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("should return nil when error is nil", func(t *testing.T) {
		eoid := erroroid.NewErroroid()
		got := eoid.Error(nil)

		if got != nil {
			t.Errorf("got %q want nil", got)
		}
	})
}

func ExampleErroroid_Error_string() {
	errFmt := erroroid.Format("encountered an error in [#file:#line @#func] #err")
	eoid := erroroid.NewErroroid(errFmt)
	fmt.Println(eoid.Error("my custom error"))
}

func ExampleErroroid_Error_error() {
	errFmt := erroroid.Format("encountered an error in [#file:#line @#func] #err")
	eoid := erroroid.NewErroroid(errFmt)
	fmt.Println(eoid.Error(errors.New("my custom error")))
}

func TestFormat(t *testing.T) {
	eoid := erroroid.NewErroroid(erroroid.Format("AN ERROR OCCURRED: [#file:#line @#func] #err \n"))

	filePath, _ := filepath.Abs(FileName)
	errStr := "this is my custom error"
	err := eoid.Error(errors.New(errStr))
	_, _, lnNum, _ := caller()
	errStrResult := "AN ERROR OCCURRED: [" + filePath + ":" + strconv.Itoa(lnNum-1) + " @TestFormat] " + errStr + " \n"

	if errStrResult != err.Error() {
		t.Errorf("got %q want %q", err, errStrResult)
	}
}

func ExampleFormat() {
	errFmt := erroroid.Format("encountered an error in [#file:#line @#func] #err")
	_ = erroroid.NewErroroid(errFmt)
}

func TestPrintLog(t *testing.T) {
	t.Run("it should print error to log", func(t *testing.T) {
		var str bytes.Buffer

		eoid := erroroid.NewErroroid(erroroid.PrintLog(true))

		filePath, _ := filepath.Abs(FileName)
		errStr := "this is my custom error"
		log.SetOutput(&str)
		_ = eoid.Error(errors.New(errStr))
		_, _, lnNum, _ := caller()
		errStrResult := "ERROR: [" + filePath + ":" + strconv.Itoa(lnNum-1) + " @TestPrintLog] -> " + errStr + "\n"

		if !strings.Contains(str.String(), errStrResult) {
			t.Errorf("\ngot  %q\nwant %q", str, errStrResult)
		}
	})

	t.Run("it should not print error to log", func(t *testing.T) {
		var str bytes.Buffer

		eoid := erroroid.NewErroroid(erroroid.PrintLog(false))

		filePath, _ := filepath.Abs(FileName)
		errStr := "this is my custom error"
		log.SetOutput(&str)
		_ = eoid.Error(errors.New(errStr))
		_, _, lnNum, _ := caller()
		errStrResult := "ERROR: [" + filePath + ":" + strconv.Itoa(lnNum-1) + " @TestPrintLog] -> " + errStr + "\n"

		if strings.Contains(str.String(), errStrResult) {
			t.Errorf("\ngot  %q\nwant %q", str, errStrResult)
		}
	})
}

func ExamplePrintLog() {
	errFmt := erroroid.PrintLog(true)
	_ = erroroid.NewErroroid(errFmt)
}
