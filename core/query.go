package core

import (
	"HackChrome/model"
	"HackChrome/utils"
	"database/sql"
	"fmt"
)

// QueryPasswordInfo Used to descrpt password
func QueryPasswordInfo(file string, query *model.LoginInfoQuery) ([]model.LoginInfo, error) {
	result := []model.LoginInfo{}
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// You can use %% for literal %
	sql := fmt.Sprintf(`SELECT action_url, username_value, password_value FROM logins where action_url like '%%%s%%'`, query.KeyWord)
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var url string
		var username string
		var encryptedPwd []byte
		err = rows.Scan(&url, &username, &encryptedPwd)
		if err != nil {
			return nil, err
		}
		if len(encryptedPwd) > 0 {
			password, err := utils.ChromeDecrypt(query.Key, encryptedPwd[3:])
			if err != nil {
				return nil, err
			}
			if len(url) > 0 {
				loginInfo := model.LoginInfo{}
				loginInfo.URL = url
				loginInfo.UserName = username
				loginInfo.Password = password
				result = append(result, loginInfo)
			}
		}
	}

	return result, nil
}
