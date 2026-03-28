package usecase

import (
	"context"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"workshop-storage-api-docs/internal/model"
	"workshop-storage-api-docs/pkg/apierror"

	"github.com/google/uuid"
	storage_go "github.com/supabase-community/storage-go"
)

type IMediaUsecase interface {
	UploadImage(ctx context.Context, fileHeader *multipart.FileHeader, uploadType string) (*model.UploadResponse, error)
}

type MediaUsecase struct {
	storageClient *storage_go.Client
	bucketName    string
}

func NewMediaUsecase(storageClient *storage_go.Client) *MediaUsecase {
	return &MediaUsecase{
		storageClient: storageClient,
		bucketName:    os.Getenv("SUPABASE_BUCKET"),
	}
}

func (u *MediaUsecase) UploadImage(ctx context.Context, fileHeader *multipart.FileHeader, uploadType string) (*model.UploadResponse, error) {
	// validate file size
	const maxFileSize = 5 * 1024 * 1024
	if fileHeader.Size > maxFileSize {
		return nil, apierror.New(400, "file size exceeds the maximum limit of 5MB")
	}

	// validate file type
	file, err := fileHeader.Open()
	if err != nil {
		return nil, apierror.New(500, "failed to open file")
	}
	defer file.Close()

	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		return nil, apierror.New(500, "failed to read file")
	}

	file.Seek(0, 0)

	contentType := http.DetectContentType(buffer)

	var ext string
	switch contentType {
	case "image/jpeg":
		ext = ".jpg"
	case "image/png":
		ext = ".png"
	default:
		return nil, apierror.New(400, "unsupported file type, only JPEG and PNG are allowed")
	}

	// determine folder
	var folder string
	switch strings.ToLower(uploadType) {
	case "restaurant":
		folder = "restaurants/"
	case "item":
		folder = "items/"
	default:
		folder = "others/"
	}

	// generate unique file name, e.g. 7f437a6f-eb2b-4c5c-879d-50149b502bad.png
	fileName := uuid.New().String() + ext
	filePath := folder + fileName

	// upload file to supabase
	fileOptions := &storage_go.FileOptions{
		ContentType: &contentType,
	}

	_, err = u.storageClient.UploadFile(u.bucketName, filePath, file, *fileOptions)
	if err != nil {
		return nil, apierror.New(500, "failed to upload file to storage")
	}

	// get public url
	urlResponse := u.storageClient.GetPublicUrl(u.bucketName, filePath)

	return &model.UploadResponse{
		ImageURL: urlResponse.SignedURL,
	}, nil
}
