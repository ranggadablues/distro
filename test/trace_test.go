package test

import (
	"fmt"
	"github.com/ranggadablues/tracer"
	"os"
	"path/filepath"
	"testing"
)

var (
	path, _  = os.Getwd()
	filename = "LOG.log"
	pathfile = filepath.Join(path, filename)
)

func TestTracePath(t *testing.T) {
	log := tracer.NewTrace(true, filename, path)
	e := log.Info("test log trace")
	if e != nil {
		t.Error("Found " + e.Error())
		return
	}

	e = log.Error("ini error")
	if e != nil {
		t.Error("Found " + e.Error())
		return
	}

	e = log.Warning("ini warning")
	if e != nil {
		t.Error("Found " + e.Error())
		return
	}

	fmt.Printf("success test path:%v\n", pathfile)
	fmt.Println()
}

func TestTrace(t *testing.T) {
	log := tracer.NewTrace(false, "", "")
	log.Info("ini info")
	log.Error("ini error")
	log.Warning("ini warning")
}
