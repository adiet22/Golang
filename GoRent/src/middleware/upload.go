package middleware

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/adiet95/Golang/GoRent/src/helpers"
)

func UploadFile(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const MAX_UPLOAD_SIZE = 1024 * 1024 // 1MB
		r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
		if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
			http.Error(w, "The uploaded file is too big. Please choose an file that's less than 1MB in size", http.StatusBadRequest)
			return
		}
		// The argument to FormFile must match the name attribute
		// of the file input on the frontend
		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			helpers.New("invalid attribute", 401, true).Send(w)
			return
		}
		defer file.Close()

		// Create the uploads folder if it doesn't
		// already exist
		err = os.MkdirAll("./uploads", os.ModePerm)
		if err != nil {
			helpers.New("error build file location", 401, true).Send(w)
			return
		}

		fileName := r.FormValue("file_name")
		// Create a new file in the uploads directory
		res := fmt.Sprintf("./uploads/%d-%s%s", time.Now().UnixNano(), fileName, filepath.Ext(fileHeader.Filename))
		dst, err := os.Create(res)
		if err != nil {
			helpers.New("error while upload file", 401, true).Send(w)
			return
		}
		defer dst.Close()

		// Copy the uploaded file to the filesystem
		// at the specified destination
		_, err = io.Copy(dst, file)
		if err != nil {
			helpers.New("error copy filesystem", 401, true).Send(w)
			return
		}
		helpers.New("File uploaded successful", 200, false).Send(w)
		next.ServeHTTP(w, r)
	}
}
