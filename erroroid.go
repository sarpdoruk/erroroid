// It's basically `error` type on steroids that helps you to easily find the error in your code by printing the filename, line, function name and the error message in a given format to stdout.
package erroroid

import (
	"fmt"
	"log"
	"runtime"
	"sort"
	"strings"
)

const (
	errParam      = "#err"
	fileParam     = "#file"
	lineParam     = "#line"
	funcNameParam = "#func"
	defFmt        = "ERROR: [#file:#line @#func] -> #err"
)

var fmtParamTypes = map[string]string{
	errParam:      "%v",
	fileParam:     "%s",
	lineParam:     "%d",
	funcNameParam: "%s",
}

// The main type that is the core to this package.
type Erroroid struct {
	err         interface{}
	errPos      int
	file        string
	filePos     int
	funcName    string
	funcNamePos int
	line        int
	linePos     int
	paramPos    []int
	fmt         string
	fmtParams   []interface{}
	printLog    bool
}

// The type that is used to modify Erroroid configuration.
type Option func(*Erroroid)

// Method used to set the format of the resulting string of the error.
//
// Default:
//     "ERROR: [#file:#line @#func] -> #err"
//
// You can set a custom format for your errors by including the following reserved placeholders in your format string;
//
//     #file
//
// Prints the full path of the file that error occurred in.
//
//     #line
//
// Prints the line number of the file that error occurred.
//
//     #func
//
// Prints the function name that error occurred in.
//
//     #err
//
// Prints the error as a string.
func Format(f string) Option {
	return func(eoid *Erroroid) {
		eoid.setFormat(f)
	}
}

// Method used to set whether the Error method should print the error to log or not.
//
// Default:
//     true
func PrintLog(p bool) Option {
	return func(eoid *Erroroid) {
		eoid.setPrintLog(p)
	}
}

// Returns a new Erroroid struct with a sensible default configuration.
func NewErroroid(opts ...Option) *Erroroid {
	eoid := &Erroroid{}

	eoid.setFormat(defFmt)
	eoid.setPrintLog(true)

	for _, opt := range opts {
		opt(eoid)
	}

	return eoid
}

// The main method that is used to create and represent errors in the provided format.
//
// It takes either a "string" or an "error" type and returns it in the "error" type so that it can be used in rest of the program natively if desired.
func (e *Erroroid) Error(err interface{}) error {
	if err != nil {
		pc, file, line, _ := runtime.Caller(1)
		f := runtime.FuncForPC(pc)
		funcName := f.Name()
		i := strings.LastIndex(funcName, "/")
		funcName = funcName[i+1:]
		i = strings.Index(funcName, ".")
		funcNameDtls := strings.Split(funcName[i+1:], ".")

		e.err = err
		e.file = file
		e.funcName = funcNameDtls[0]
		e.line = line
		e.fmtParams = make([]interface{}, 4)

		for i, pos := range e.paramPos {
			switch pos {
			case e.errPos:
				e.fmtParams[i] = e.err
			case e.filePos:
				e.fmtParams[i] = e.file
			case e.funcNamePos:
				e.fmtParams[i] = e.funcName
			case e.linePos:
				e.fmtParams[i] = e.line
			}
		}

		if e.printLog {
			log.Printf(e.fmt, e.fmtParams...)
		}

		return fmt.Errorf(e.fmt, e.fmtParams...)
	}

	return nil
}

func (e *Erroroid) setFormat(f string) {
	e.errPos = strings.Index(f, errParam)
	e.filePos = strings.Index(f, fileParam)
	e.funcNamePos = strings.Index(f, funcNameParam)
	e.linePos = strings.Index(f, lineParam)
	e.paramPos = []int{e.errPos, e.filePos, e.funcNamePos, e.linePos}

	sort.Ints(e.paramPos)

	for _, param := range [4]string{errParam, fileParam, funcNameParam, lineParam} {
		f = strings.Replace(f, param, fmtParamTypes[param], 1)
	}

	e.fmt = f
}

func (e *Erroroid) setPrintLog(p bool) {
	e.printLog = p
}
