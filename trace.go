package tracer

import (
	"errors"
	// "fmt"
	"log"
	// "tools"
	"os"
	"path/filepath"
	"strings"
)

type Trace struct {
	tracePath    string
	traceName    string
	traceFile    bool
	traceInfo    *log.Logger
	traceError   *log.Logger
	traceWarning *log.Logger
}

func NewTrace(addTofile bool, filename string, path string) *Trace {
	t := new(Trace)
	t.traceFile = addTofile
	t.traceName = filename
	t.tracePath = path

	t.initTrace()
	return t
}

func (t *Trace) initTrace() {
	t.traceInfo = log.New(os.Stdout, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
	t.traceError = log.New(os.Stdout, "ERROR ", log.Ldate|log.Ltime|log.Lshortfile)
	t.traceWarning = log.New(os.Stdout, "WARNING ", log.Ldate|log.Ltime|log.Lshortfile)
}

func (t *Trace) Info(msg string) error {
	return t.createTrace(msg, "INFO")
}

func (t *Trace) Error(msg string) error {
	return t.createTrace(msg, "ERROR")
}

func (t *Trace) Warning(msg string) error {
	return t.createTrace(msg, "WARNING")
}

func (t *Trace) createTrace(msg string, traceType string) error {
	typeoftrace := strings.ToUpper(traceType) + " "
	if t.traceFile {
		if _, e := os.Stat(t.tracePath); os.IsNotExist(e) {
			// return errors.New("Trace error: " + e.Error())
			dest, e := os.Create(filepath.Join(t.tracePath, t.traceName))
			defer dest.Close()
			if e != nil {
				return errors.New("create log: " + e.Error())
			}
		}

		filename := filepath.Join(t.tracePath, t.traceName)
		f, e := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0)
		if e != nil {
			return errors.New("Add file error: " + e.Error())
		}
		defer f.Close()

		traceFile := log.New(f, typeoftrace, log.Ldate|log.Ltime|log.Lshortfile)
		traceFile.Println(msg)
	} else {
		if typeoftrace == "INFO " {
			t.traceInfo.Println(msg)
		} else if typeoftrace == "ERROR " {
			t.traceError.Println(msg)
		} else if typeoftrace == "WARNING " {
			t.traceWarning.Println(msg)
		}
	}

	return nil
}
