package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKeyFromHeader(header http.Header) (string ,error){
	val:=header.Get("Authorization")
	if val== ""{
		return "" , errors.New("no authentication info found")
	}
	vals := strings.Split(val, " ")
	if len(vals) !=2  || vals[0] != "ApiKey"{
		return "" , errors.New("wrong authentication info")

	}
	return vals[1] , nil
}