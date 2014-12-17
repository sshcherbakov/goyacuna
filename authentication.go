package goyacuna

import (
	"time"
	"fmt"
	"crypto/sha512"
	"encoding/hex"
)

const (
	H_ApiTokenId 	string = "Api-Token-Id"
	H_ApiToken 		string = "Api-Token"
	H_ApiTokenOTP 	string = "Api-Token-OTP"

)

type apiTokenInput struct {

	secret 	string
	method 	string
	path 	string
	query 	string
	body	string

}

func ApiToken(input *apiTokenInput) string {

	salt := fmt.Sprintf("%d", time.Now().UnixNano() / 1000000 )
	hashInput := salt + "@" + input.secret + "@" + input.method + "@" + input.path

	if input.query != "" {
		hashInput += "?" + input.query
	}

	if input.body != "" {
		hashInput += "@" + input.body
	}

	fmt.Println(hashInput)

	hashValue := sha512.Sum512([]byte(hashInput))
	return salt + "T" + hex.EncodeToString( hashValue[:] )

}
