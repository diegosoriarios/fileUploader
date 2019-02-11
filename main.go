package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func uploadFile(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Uploading, File\n")

	r.ParseMultipartForm(10 << 20)
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving file from form-data")
		fmt.Println(err)
		return
	}

	defer file.Close()
	
	fmt.Printf("Upload File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	tempFile, err := ioutil.TempFile("temp-images", "upload-*.png")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	tempFile.Write(fileBytes)
	fmt.Fprintf(w, "Successfully Uploaded File\n")
}

func setupRoutes(){
	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":8000", nil)
}

func main(){
	fmt.Println("Go File Upload")
	setupRoutes() 
}