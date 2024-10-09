package utils

import (
	"chapter1/config"
	"context"
	"io"
	"log"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type CloudinaryUtilityInterface interface {
	FileCheck(file *multipart.FileHeader) (multipart.File, error)
	UploadToCloudinary(file io.Reader, filename string) (string, error)
	FileOpener(file *multipart.FileHeader) (multipart.File, error)
}

type CloudinaryUtility struct{}

func NewCloudinaryUtility() CloudinaryUtilityInterface {
	return &CloudinaryUtility{}
}

func (c *CloudinaryUtility) UploadToCloudinary(file io.Reader, filename string) (string, error) {
	cld, err := cloudinary.NewFromURL(config.ImportSetting().CldKey)

	if err != nil {
		return "", err
	}

	uploadParams := uploader.UploadParams{
		Folder:   "images",
		PublicID: filename,
	}

	uploadResult, err := cld.Upload.Upload(context.Background(), file, uploadParams)
	if err != nil {
		return "", err
	}

	publicURL := uploadResult.SecureURL
	return publicURL, nil
}

func (c *CloudinaryUtility) FileCheck(file *multipart.FileHeader) (multipart.File, error) {
	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()
	return src, nil
}

func (c *CloudinaryUtility) FileOpener(file *multipart.FileHeader) (multipart.File, error) {

	src, err := file.Open()
	if err != nil {
		log.Println("error when open image", err)
		return src, err
	}
	defer src.Close()

	return src, nil
}
