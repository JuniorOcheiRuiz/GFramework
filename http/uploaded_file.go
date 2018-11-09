package http

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

type UploadedFile struct {
	FileHeader *multipart.FileHeader
}

func NewUploadFile(fileHeader *multipart.FileHeader) *UploadedFile {
	return &UploadedFile{FileHeader: fileHeader}
}

// Save the file into destination
func (u *UploadedFile) Save(destination string) error {
	fileSrc, err := u.FileHeader.Open()

	if err != nil {
		return err
	}

	fileDest, err := os.Create(destination)

	if err != nil {
		return err
	}

	_, err = io.Copy(fileDest, fileSrc)

	return err
}

func (u *UploadedFile) GetType() string {
	return u.FileHeader.Header.Get("Content-Type")
}

func (u *UploadedFile) GetExtension() string {
	return filepath.Ext(u.FileHeader.Filename)
}

func (u *UploadedFile) GetSize() int64 {
	return u.FileHeader.Size
}
