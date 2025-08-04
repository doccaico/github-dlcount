package main

import (
	"fmt"
	"github.com/doccaico/github-dlcount"
	"os"
	"os/exec"
	"path/filepath"
)

func usage() {
	fmt.Println("Usage:\n  github_dlcount [URL]")
	fmt.Println("\tgithub_dlcount doccaico/hogehoge")
	os.Exit(0)
}

func checkRequirements() {
	for _, exe := range [...]string{"jq", "curl"} {
		_, err := exec.LookPath(exe)
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	if len(os.Args) != 2 {
		usage()
	}

	checkRequirements()

	home_path, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	// cd ~/Downloadsする
	download_path := filepath.Join(home_path, "Downloads")
	err = os.Chdir(download_path)
	if err != nil {
		panic(err)
	}

	// latestを取得する
	url := fmt.Sprintf("https://api.github.com/repos/%s/releases/latest", os.Args[1])
	cmd := exec.Command("curl", "-fsSOL", url)
	err = cmd.Run()
	if err != nil {
		panic(err)
	}

	names := utils.GetNames(url)
	counts := utils.GetCounts(url)
	if len(names) != len(counts) {
		panic("error: len(names) != len(counts)")
	}

	name_max_width := 0
	count_max_width := 0
	for i := 0; i < len(names); i++ {
		if name_max_width < len(names[i]) {
			name_max_width = len(names[i])
		}
		if count_max_width < len(counts[i]) {
			count_max_width = len(counts[i])
		}
	}

	name_max_width += 8

	for i := 0; i < len(names); i++ {
		fmt.Printf("%-*s%*s\n", name_max_width, names[i], count_max_width, counts[i])
	}

	// latestを削除する
    utils.RemoveLatestFile()
}
