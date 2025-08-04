package utils

import (
	"os/exec"
	"strings"
)

func GetNames(url string) []string {
	out, err := exec.Command("cmd", "/c", "type latest | jq -r .assets[].name").Output()
	if err != nil {
		panic(err)
	}
	names := strings.Split(string(out), "\r\n")
	return names[:len(names)-1]
}

func GetCounts(url string) []string {
	out, err := exec.Command("cmd", "/c", "type latest | jq -r .assets[]^|select(.name).download_count").Output()
	if err != nil {
		panic(err)
	}
	counts := strings.Split(string(out), "\r\n")
	return counts[:len(counts)-1]
}

func RemoveLatestFile() {
	cmd := exec.Command("cmd", "/c", "del latest")
    err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
