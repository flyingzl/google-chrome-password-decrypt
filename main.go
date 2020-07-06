package main

import (
	"HackChrome/core"
	"HackChrome/model"
	"HackChrome/utils"
	"flag"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os/user"
)

var keyword string
var profile string

func init() {
	flag.StringVar(&keyword, "k", "", "网址关键字")
	flag.StringVar(&profile, "p", "", "Google Chrome个人资料路径, 输入chrome://version 可以查看")
}

func checkError(err error) bool {
	if nil != err {
		fmt.Println(err)
		return true
	}
	return false
}

func main() {

	flag.Parse()

	user, err := user.Current()

	if checkError(err) {
		return
	}

	dbFile := fmt.Sprintf("%s/Library/ApplicationSupport/Google/Chrome/Default/Login Data", user.HomeDir)

	if profile != "" {
		dbFile = profile + "/Login Data"
	}

	if !utils.FileExists(dbFile) {
		fmt.Printf("file %s doesn't exists, please check it!\n", dbFile)
		return
	}

	file, err := utils.BuildTempFile(dbFile)

	if err != nil {
		fmt.Println(err)
		return
	}

	key, err := utils.GetDerivedKey()

	if checkError(err) {
		return
	}

	queryModel := model.LoginInfoQuery{
		KeyWord: keyword,
		Key:     key,
	}

	result, err := core.QueryPasswordInfo(file.Name(), &queryModel)
	if checkError(err) {
		return
	}

	utils.FormatOutput(result)

}
