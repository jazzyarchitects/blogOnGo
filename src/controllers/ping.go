package controllers

import (
	"net/http"
	"fmt"
	"os"
	"io"
	"strconv"
)

type PingController struct {
	
}


func NewPingController() *PingController {
	return &PingController{}
}

/*
This will retrieve a list of all the blogs stored in the database\
 */
func (pc PingController)RespondPing(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello...");
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, "{\"success\": true}")
}

func (pc PingController)SendSwagger(w http.ResponseWriter, r *http.Request){
	Openfile, err := os.Open("/Users/rentomojo/Documents/go/src/blogOnGo/swagger.json")
	defer Openfile.Close()

	if err != nil{
		fmt.Fprint(w, "{\"error\": true}")
	}

	FileHeader := make([]byte, 512)
	//Copy the headers into the FileHeader buffer
	Openfile.Read(FileHeader)
	//Get content type of file
	FileContentType := http.DetectContentType(FileHeader)

	//Get the file size
	FileStat, _ := Openfile.Stat()                     //Get info from file
	FileSize := strconv.FormatInt(FileStat.Size(), 10) //Get file size as a string

	//Send the headers
	w.Header().Set("Content-Disposition", "attachment; filename=swagger.json")
	w.Header().Set("Content-Type", FileContentType)
	w.Header().Set("Content-Length", FileSize)

	//Send the file
	//We read 512 bytes from the file already so we reset the offset back to 0
	Openfile.Seek(0, 0)
	io.Copy(w, Openfile) //'Copy' the file to the client
	return
}
