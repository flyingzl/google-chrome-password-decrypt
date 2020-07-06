package model

// LoginInfo Model
type LoginInfo struct {
	URL      string
	UserName string
	Password string
}

// LoginInfoQuery Model
type LoginInfoQuery struct {
	Key     []byte
	KeyWord string
}
