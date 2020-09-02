package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/zcong1993/download-release/utils"
)

func main() {
	host := os.Getenv("GITHUB_HOST")
	token := os.Getenv("GITHUB_TOKEN")

	flag.Parse()

	args := flag.Args()

	if len(args) != 1 {
		fmt.Println("use as download-release [owner/repo]")
		os.Exit(1)
	}

	repoUrl := args[0]

	author, repo := utils.ParseRepoUrl(repoUrl)

	apiUrl := utils.BuildReleaseUrl(host, author, repo, true)
	resp, err := utils.MakeGetRequest(apiUrl, token)
	if err != nil {
		log.Fatal(err)
	}
	assets := utils.GetAssetList(resp)
	prompt := &survey.Select{
		Message: "Choose an assets:",
		Options: assets,
	}
	var assetsUrl string
	err = survey.AskOne(prompt, &assetsUrl, nil)
	if err != nil {
		log.Fatal(err)
	}

	utils.Download(assetsUrl)
}
