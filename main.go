package main

import (
	"fmt"
	"github.com/ddliu/go-httpclient"
	flags "github.com/jessevdk/go-flags"
	"os"
)

const (
	USERAGENT       = "golang deploygate cl tool"
	TIMEOUT         = 30
	CONNECT_TIMEOUT = 5
)

type Options struct {
	ApiKey      string `short:"a" long:"apikey" description:"Your DeployGate Api Key"`
	UserName    string `short:"u" long:"username" description:"Your DeployGate User Name"`
	AppFilePath string `short:"f" long:"AppFilePath" description:"Upload App File Path"`
}

var opts Options

func main() {
	parser := flags.NewParser(&opts, flags.Default)
	parser.Name = "godeploy"
	parser.Usage = "[OPTIONS] PATTERN [PATH]"

	args, _ := parser.Parse()

	if (opts.ApiKey == "" || opts.UserName == "" || opts.AppFilePath == "") && len(args) == 0 {
		parser.WriteHelp(os.Stdout)
		return
	}

	c := httpclient.NewHttpClient(map[int]interface{}{
		httpclient.OPT_USERAGENT: USERAGENT,
		httpclient.OPT_TIMEOUT:   TIMEOUT,
	})

	url := "https://deploygate.com/api/users/" + opts.UserName + "/apps"
	res, err := c.Post(url, map[string]string{
		"@file": opts.AppFilePath,
		"token": opts.ApiKey,
	})

	fmt.Println(res, err)
}
