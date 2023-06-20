package update

import (
	"io/ioutil"
	"strings"

	version "github.com/mcuadros/go-version"
)

func GetLocalVersion() (string, error) {
	content, err := ioutil.ReadFile("version.txt")
	return strings.TrimSpace(string(content)), err
}

func GetPatchNotes() (string, error) {
	content, err := ioutil.ReadFile("patchnotes.txt")
	return strings.TrimSpace(string(content)), err
}

func CompareVersions(version1 string, version2 string) bool {
	return version.Compare(version1, version2, "<")
}
