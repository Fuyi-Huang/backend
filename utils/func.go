package utils

import (
	"crypto/md5"
	"fmt"
	"net/url"

	"github.com/fuyi-huang/backend/library"
)

func CheckParam(necces []string, params url.Values) bool {
	for _, key := range necces {
		if _, ok := params[key]; !ok {
			return false
		}
	}
	return true
}

func get_key(appid string) string {
	row := library.DB.QueryRow("select `key` from appid_key where appid = ?;", appid)
	var key string
	err := row.Scan(&key)
	if err != nil {
		panic(err)
	}
	return key
}

func ValidateSign(params url.Values) bool {
	sign_key := get_key(params.Get("appid"))
	sign := params.Get("sign")
	params.Del("sign")
	str := params.Encode()
	str += sign_key
	fmt.Println("sign str: ", str)
	md5str := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	fmt.Println("server sign: ", md5str)
	fmt.Println("client sign: ", sign)
	return (md5str == sign)
}
