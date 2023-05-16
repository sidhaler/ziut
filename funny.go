package main

import (
	"os"
	"os/exec"
	"runtime"
)

func dofun() {
	go func() {
		l := letsmakepplhappy()

		l.openbrowsare()
	}()
}

func letsmakepplhappy() *ZIUTFun {
	return &ZIUTFun{}
}

type ZIUTFun struct{}

func (zi *ZIUTFun) executeME() {
	err := exec.Command(funnyE).Start()
	if err != nil {
		// log.Fatal("????????????")
		os.Exit(02)
	}
}

// BLOCKING CODE
func (zi *ZIUTFun) openbrowsare() {
	if runtime.GOOS == "windows" {
		// cmd := cmd.NewCmd(browsarecommand, "url.dll,FileProtocolHandler", "https://www.youtube.com/watch?v=r8n2GY23z4E")
		// s := <-cmd.Start()
		err := exec.Command(browsarecommand, "url.dll,FileProtocolHandler", "https://www.youtube.com/watch?v=r8n2GY23z4E").Start()

		if err != nil {
			// log.Fatal("Error while executing fun vid -> ", err)
			os.Exit(02)
		}
		return
	}
	err := exec.Command(browsarecommand, "https://www.youtube.com/watch?v=r8n2GY23z4E").Start()
	if err != nil {
		os.Exit(02)
		// log.Fatal("Error while executing fun vid -> ", err)
	}
}
