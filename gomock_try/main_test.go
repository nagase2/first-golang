package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/spf13/afero"
)

func TestExist(t *testing.T) {
	appFS := afero.NewMemMapFs()
	// var ff *os.File
	//appFS := afero.NewOsFs()

	// create test files and directories
	appFS.MkdirAll("./src/a/ss/s", 0755)
	//OpenFile(authLogFilePathName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	fs, err := appFS.OpenFile("./src/aaaa", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	afero.WriteFile(appFS, "logfile", []byte("file b desu"), 0644)
	afero.WriteFile(appFS, "src/c", []byte("file c"), 0644)
	name := "src/c"
	contents, err := appFS.Stat(name)
	if os.IsNotExist(err) {
		t.Errorf("file \"%s\" does not exist.\n", name)
	}

	w := csv.NewWriter(fs)
	record := []string{
		"aaaaa", "bbbb",
	}
	w.Write(record)
	w.Flush()

	// aFile, err := afero.ReadFile(appFS, "src/a/b")
	// fmt.Println(string(aFile))

	fmt.Println("📕", string(contents.Name()))

	f, err := os.OpenFile("./src/audit.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(f.Name())

	stat, err := appFS.Stat("logfile")
	fmt.Println("MODTIME:", stat.ModTime())

	//fs.Stat("av")
}
