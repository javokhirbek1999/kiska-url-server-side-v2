package utils

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"strconv"
)

func HashURL(url string, user_id uint) (string, error) {

	if len(url) == 0 {
		return "", errors.New("URL cannot be blank")
	}

	if url[:4] != "http" {
		if url[:3] != "www" {
			return "", errors.New("Invalid url, url should start either with http/https or www")
		}
	}

	temp := strconv.Itoa(int(user_id))

	url = temp[:len(temp)/2] + url + temp[len(temp)/2:]

	hashedString := md5.Sum([]byte(url))

	return hex.EncodeToString(hashedString[:][:6]), nil

}
