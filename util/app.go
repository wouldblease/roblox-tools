package util

import (
	"fmt"

	"github.com/wouldblease/roblox-tools/update"
)

func StartApp() {
	Patchnotes, Error := update.GetPatchNotes()
	if Error != nil {
		panic(Error)
	}
	LocalVer, Erra := update.GetLocalVersion()
	if Erra != nil {
		panic(Erra)
	}

	CurrentVer, _ := update.GetLatestVersion()
	UpdateStatus := update.CompareVersions(LocalVer, CurrentVer)
	println(LocalVer)
	println(CurrentVer)
	println(UpdateStatus)
	switch UpdateStatus {
	case true:
		fmt.Println("An update is available.")
	case false:
		fmt.Println("You're updated!")
	}
	fmt.Println(CurrentVer)
	fmt.Println(Patchnotes)
	// util.GetProxyScrapeProxies()
}
