package utils

import (
	"HackChrome/model"
	"fmt"
	"github.com/alexeyco/simpletable"
	"io/ioutil"
	"os"
	"strconv"
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

// FormatOutput used to pretty query result
func FormatOutput(results []model.LoginInfo) {
	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "URL"},
			{Align: simpletable.AlignCenter, Text: "User Name"},
			{Align: simpletable.AlignCenter, Text: "Password"},
		},
	}
	for _, row := range results {
		r := []*simpletable.Cell{
			{Text: row.URL},
			{Text: row.UserName},
			{Text: row.Password},
		}
		table.Body.Cells = append(table.Body.Cells, r)
	}

	table.Footer = &simpletable.Footer{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignRight, Span: 3, Text: "Total Records: " + strconv.Itoa(len(results))},
		},
	}
	table.SetStyle(simpletable.StyleUnicode)
	fmt.Println(table.String())

}

// FileExists used to check whether file exists or not
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// RemoveFile used to remove a  file
func RemoveFile(file *os.File) bool {
	err := os.Remove(file.Name())
	if err != nil {
		return false
	}
	return true
}
