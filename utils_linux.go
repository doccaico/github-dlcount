package utils

import (
	"os/exec"
	"strings"
)

func GetNames(url string) []string {
	out, err := exec.Command("sh", "-c", "cat latest | jq -r '.assets[].name'").Output()
	if err != nil {
		panic(err)
	}
	names := strings.Split(string(out), "\n")
	return names[:len(names)-1]
}

func GetCounts(url string) []string {
	out, err := exec.Command("sh", "-c", "cat latest | jq -r '.assets[] | select(.name).download_count'").Output()
	if err != nil {
		panic(err)
	}
	counts := strings.Split(string(out), "\n")
	return counts[:len(counts)-1]
}

func RemoveLatestFile() {
	cmd := exec.Command("rm", "latest")
    err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
