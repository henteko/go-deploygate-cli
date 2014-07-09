package main

import (
	"flag"
	"fmt"
	"github.com/ddliu/go-httpclient"
)

const (
	USERAGENT       = "golang deploygate cl tool"
	TIMEOUT         = 30
	CONNECT_TIMEOUT = 5
)

func main() {
	var (
		userName    string
		apiKey      string
		appFilePath string
	)

	flag.StringVar(&userName, "username", "blank", "your DeployGate User Name")
	flag.StringVar(&userName, "u", "blank", "your DeployGate User Name")
	flag.StringVar(&apiKey, "apikey", "blank", "your DeployGate Api Key")
	flag.StringVar(&apiKey, "a", "blank", "your DeployGate Api Key")
	flag.StringVar(&appFilePath, "file", "blank", "upload App File Path")
	flag.StringVar(&appFilePath, "f", "blank", "upload App File Path")
	flag.Parse()

	if userName == "" || apiKey == "" || appFilePath == "" {
		fmt.Println("Error: must input")
		return
	}

	c := httpclient.NewHttpClient(map[int]interface{}{
		httpclient.OPT_USERAGENT: USERAGENT,
		httpclient.OPT_TIMEOUT:   TIMEOUT,
	})

	url := "https://deploygate.com/api/users/" + userName + "/apps"
	res, err := c.Post(url, map[string]string{
		"@file": appFilePath,
		"token": apiKey,
	})

	fmt.Println(res, err)
}
