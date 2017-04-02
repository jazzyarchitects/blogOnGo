package controllers

import (
	"golang.org/x/crypto/scrypt"
	"math/rand"
	"time"
)

func getPasswordHash(password string) string{
	dk,err := scrypt.Key([]byte(password), []byte("MyRandomSalt"),16384,8,1,12);
	if err!=nil{
		panic(err)
	}
	return string(dk);
}

func GetRandomString(length int) string{
	var result string = ""
	allowed := "abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNOPQRSTUVWXYZ123456789$-";

	rand.Seed(time.Now().UTC().UnixNano())

	for i:=0;i<length;i++{
		result += string(allowed[rand.Intn(len(allowed))]);
	}
	return result;
}