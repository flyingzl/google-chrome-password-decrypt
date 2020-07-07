package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"github.com/havoc-io/go-keytar"
	"golang.org/x/crypto/pbkdf2"
	"io/ioutil"
	"os"
	"os/user"
)

var (
	// SALT for AES
	SALT = "saltysalt"
	// ITERATIONS For AES
	ITERATIONS = 1003
	// KEYLENGTH for AES
	KEYLENGTH = 16
)

// GetDerivedKey 获取密钥串
func GetDerivedKey() ([]byte, error) {
	user, _ := user.Current()
	keyDir := user.HomeDir + "/.hackChrome/"
	if !FileExists(keyDir) {
		os.Mkdir(keyDir, os.ModePerm)
	}
	keyPath := keyDir + "key"
	input, err := ioutil.ReadFile(keyPath)
	if err != nil {
		keychain, err := keytar.GetKeychain()
		if err != nil {
			return nil, err
		}
		chromePassword, err := keychain.GetPassword("Chrome Safe Storage", "Chrome")
		if err != nil {
			return nil, err
		}
		// save chromePassword in keyfile
		ioutil.WriteFile(keyPath, []byte(chromePassword), 0644)
		dk := pbkdf2.Key([]byte(chromePassword), []byte(SALT), ITERATIONS, KEYLENGTH, sha1.New)
		return dk, nil
	}

	dk := pbkdf2.Key(input, []byte(SALT), ITERATIONS, KEYLENGTH, sha1.New)
	return dk, nil

}

func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// ChromeDecrypt used for decrypting password
// Decryption based on http://n8henrie.com/2014/05/decrypt-chrome-cookies-with-python/
func ChromeDecrypt(key []byte, encrypted []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	iv := make([]byte, 16)
	for i := 0; i < 16; i++ {
		iv[i] = ' '
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(encrypted))
	blockMode.CryptBlocks(origData, encrypted)
	origData = pkcs5UnPadding(origData)
	return string(origData), nil
}
