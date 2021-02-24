package squashfs

import (
	"context"
	"io/ioutil"
	"log"
	"strings"

	"github.com/containerd/containerd/remotes"
	"github.com/deislabs/oras/pkg/content"
	"github.com/deislabs/oras/pkg/oras"
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
)

// SupportedCompression returns the list of supported compression algorithms
func SupportedCompression() []string {
	return []string{"gzip", "lzo", "lz4", "xz", "zstd"}
}

// getSquashfsMediaTypes is a workaround because go doesn't support maps as constants
func getSquashfsMediaTypes() map[string]string {
	const baseMediaType = "application/vnd.oci.image.layer.v1.squashfs"

	mediaTypes := make(map[string]string)
	mediaTypes[""] = baseMediaType
	for _, compression := range SupportedCompression() {
		mediaTypes[compression] = baseMediaType + "+" + compression
	}
	return mediaTypes
}

// getSquashfsMediaType uses the fake constant to find the right media type for the compression ratio
func getSquashfsMediaType(compression string) string {
	if compression == "none" {
		return getSquashfsMediaTypes()[""]
	}
	return getSquashfsMediaTypes()[strings.ToLower(compression)]
}

// AllowedMediaTypes returns the valid media types for squashfs
func AllowedMediaTypes() []string {
	var allowedMediaTypes []string
	for _, t := range getSquashfsMediaTypes() {
		allowedMediaTypes = append(allowedMediaTypes, t)
	}
	return allowedMediaTypes
}

// PushSquashFS pushes a sqaushfs layer to a remote repository
func PushSquashFS(ctx context.Context, resolver remotes.Resolver, fileName string, reference string, compression string) (ocispec.Descriptor, error) {
	fileContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	// Push file(s) w custom mediatype to registry
	memoryStore := content.NewMemoryStore()
	log.Printf("Adding %s with mediaType: %s\n", fileName, getSquashfsMediaType(compression))
	desc := memoryStore.Add(fileName, getSquashfsMediaType(compression), fileContent)
	pushContents := []ocispec.Descriptor{desc}
	return oras.Push(ctx, resolver, reference, memoryStore, pushContents)
}

// PullSquashFS pulls one or more layers from a remote repository
func PullSquashFS(ctx context.Context, resolver remotes.Resolver, reference string, targetDirectory string) (ocispec.Descriptor, []ocispec.Descriptor, error) {
	log.Printf("Pulling %s to store in %s\n", reference, targetDirectory)
	fileStore := content.NewFileStore(targetDirectory)
	defer fileStore.Close()
	return oras.Pull(ctx, resolver, reference, fileStore, oras.WithAllowedMediaTypes(AllowedMediaTypes()))
}
