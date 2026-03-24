package supabase

import (
	"os"

	storage_go "github.com/supabase-community/storage-go"
)

func NewStorageClient() *storage_go.Client {
	storageURL := os.Getenv("SUPABASE_API_URL") + "/storage/v1"
	storageKey := os.Getenv("SUPABASE_API_KEY")

	return storage_go.NewClient(storageURL, storageKey, nil)
}
