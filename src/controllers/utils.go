package controllers

import "golang.org/x/crypto/scrypt"

func getPasswordHash(password string) string{
	dk,err := scrypt.Key([]byte(password), []byte("MyRandomSalt"),16384,8,1,12);
	if err!=nil{
		panic(err)
	}
	return string(dk);
}

