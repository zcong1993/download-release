package main

import (
	"github.com/zcong1993/download-release/utils"
	"gopkg.in/AlecAivazis/survey.v1"
	"log"
	"os"
)

func main() {
	host := os.Getenv("GITHUB_HOST")
	token := os.Getenv("GITHUB_TOKEN")

	var repoUrl string
	repoUrlPrompt := &survey.Input{
		Message: "Type the repo author and name: ",
	}
	err := survey.AskOne(repoUrlPrompt, &repoUrl, survey.Required)
	if err != nil {
		log.Fatal(err)
	}

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
