package utils

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
	"strings"

	"github.com/tidwall/gjson"
)

func MakeGetRequest(url, token string) (*bytes.Buffer, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	if token != "" {
		req.Header.Set("Authorization", "token "+token)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var buf bytes.Buffer
	io.Copy(&buf, resp.Body)
	return &buf, nil
}

func BuildReleaseUrl(host, author, repo string, isLatest bool) string {
	if host == "" {
		host = "https://api.github.com"
	}
	url := fmt.Sprintf("%s/repos/%s/%s/releases", host, author, repo)
	if isLatest {
		url += "/latest"
	}
	return url
}

func GetAssetList(buf *bytes.Buffer) []string {
	var list []string
	gjson.GetBytes(buf.Bytes(), "assets").ForEach(func(key, value gjson.Result) bool {
		list = append(list, value.Get("browser_download_url").String())
		return true
	})
	return list
}

func Download(url string) {
	cmd := exec.Command("curl", "-sLO", url)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Downloading...")
	err = cmd.Wait()
	if err != nil {
		log.Printf("Command finished with error: %v", err)
	}
	log.Println("Done!")
}

func ParseRepoUrl(url string) (string, string) {
	arr := strings.Split(url, "/")
	if len(arr) < 2 {
		log.Fatal("Invalid url!")
	}
	return strings.Replace(arr[0], "/", "", -1), strings.Replace(arr[1], "/", "", -1)
}
