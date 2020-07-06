package utils

import (
	"HackChrome/model"
	"fmt"
	"io/ioutil"
	"os"
)

// BuildTempFile used copy a file to tempdir and return the temp file
func BuildTempFile(source string) (*os.File, error) {
	input, err := ioutil.ReadFile(source)
	if err != nil {
		return nil, err
	}
	file, _ := ioutil.TempFile(os.TempDir(), "sql.db")
	if err != nil {
		return nil, err
	}
	err = ioutil.WriteFile(file.Name(), input, 0644)
	return file, nil
}

// FormatOutput used to ouput result
func FormatOutput(results []model.LoginInfo) {
	for _, v := range results {
		fmt.Printf("====================\n")
		fmt.Printf("Url: %s\nUsername: %s\nPassword:%s\n\n", v.URL, v.UserName, v.Password)
	}

	fmt.Printf("\nTotal Auth: %d\n", len(results))
}

// FileExists used to check whether file exists or not
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
