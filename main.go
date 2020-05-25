package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	sal "github.com/salrashid123/oauth2/google"
)

const defaultProject = "terraform-admin-278110"
const defaultServiceAccount = "service-agent@terraform-admin-278110.iam.gserviceaccount.com"
const cloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"
const iamScope = "https://www.googleapis.com/auth/iam"

func getAuthenticatedClient(ctx context.Context) *storage.Client {
	ts := getServiceAccountTokenSource(ctx)

	ptr, err := storage.NewClient(ctx, option.WithTokenSource(ts))
	if err != nil {
		log.Fatalf("Unable to acquire an authenticated Storage Client: %v", err)
	}

	return ptr
}

func getServiceAccountTokenSource(ctx context.Context) oauth2.TokenSource {
	r, err := google.DefaultTokenSource(ctx, iamScope)
	if err != nil {
		log.Fatal(err)
	}

	ts, err := sal.ImpersonatedTokenSource(
		&sal.ImpersonatedTokenConfig{
			TargetPrincipal: defaultServiceAccount,
			Lifetime:        3600 * time.Second,
			Delegates:       []string{},
			TargetScopes:    []string{cloudPlatformScope},
			RootTokenSource: r,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	return ts
}

type storageClient struct {
	*storage.Client
}

func (sc storageClient) listBuckets(ctx context.Context) []storage.BucketAttrs {
	var buckets []storage.BucketAttrs

	it := sc.Buckets(ctx, defaultProject)
	for {
		bucketData, err := it.Next()

		if err == iterator.Done {
			break
		}

		if err != nil {
			log.Fatalf("Unable to fetch the next bucket: %v", err)
		}

		buckets = append(buckets, *bucketData)
	}

	return buckets
}

func main() {
	ctx := context.Background()

	sc := storageClient{getAuthenticatedClient(ctx)}
	slice := sc.listBuckets(ctx)

	r, _ := json.MarshalIndent(slice, "", "  ")
	fmt.Printf("%s\n", r)
}
