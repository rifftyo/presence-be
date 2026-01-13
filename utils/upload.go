package utils

import (
	"fmt"
	"mime/multipart"
	"os"

	storage_go "github.com/supabase-community/storage-go"
)

func SaveFileToSupabase(file multipart.File, fileName string, fileHeader *multipart.FileHeader) (string, error) {
	supabaseURL := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_SERVICE_ROLE_KEY")
	bucketName := os.Getenv("SUPABASE_BUCKET")

	client := storage_go.NewClient(supabaseURL, supabaseKey, nil)

	opts := storage_go.FileOptions{
		ContentType: ptr(fileHeader.Header.Get("Content-Type")),
	}

	_, err := client.UploadFile(bucketName, fileName, file, opts)
	if err != nil {
		return "", fmt.Errorf("failed to upload to supabase: %v", err)
	}

	publicURL := fmt.Sprintf("%s/storage/v1/object/public/%s/%s", supabaseURL, bucketName, fileName)

	return publicURL, nil
}

func ptr(s string) *string {
	return &s
}