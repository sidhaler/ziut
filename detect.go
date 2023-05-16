package main

import "runtime"

type OrganiseVars struct{}

func (o *OrganiseVars) org() {
	sM.Lock()
	run := runtime.GOOS
	switch run {
	case "windows":
		fExt = ".exe"
		browsarecommand = "rundll32"
	case "linux":
		fExt = ""
		browsarecommand = "xdg-open"
	case "darwin":
		fExt = "not yet done "
		browsarecommand = "open"
	}

}

func organiser() *OrganiseVars {
	return &OrganiseVars{}
}
