module github.com/cray-hpe/csmoci

go 1.15

replace (
	// WARNING! Do NOT replace these without also replacing their lines in the `require` stanza below.
	// These `replace` stanzas are IGNORED when this is imported as a library
	github.com/docker/distribution => github.com/docker/distribution v0.0.0-20191216044856-a8371794149d
	github.com/docker/docker => github.com/moby/moby v17.12.0-ce-rc1.0.20200618181300-9dc6525e6118+incompatible
)

require (
	github.com/containerd/containerd v1.4.11
	github.com/deislabs/oras v0.10.0
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/mitchellh/go-homedir v1.1.0
	github.com/opencontainers/image-spec v1.0.1
	github.com/spf13/cobra v1.1.3
	github.com/spf13/viper v1.7.0
	google.golang.org/grpc v1.41.0 // indirect
)
