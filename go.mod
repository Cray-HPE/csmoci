module github.com/cray-hpe/csmoci

go 1.15

replace (
	// WARNING! Do NOT replace these without also replacing their lines in the `require` stanza below.
	// These `replace` stanzas are IGNORED when this is imported as a library
	github.com/docker/distribution => github.com/docker/distribution v0.0.0-20191216044856-a8371794149d
	github.com/docker/docker => github.com/moby/moby v17.12.0-ce-rc1.0.20200618181300-9dc6525e6118+incompatible
)

require (
	github.com/AlekSi/gocov-xml v0.0.0-20190121064608-3a14fb1c4737 // indirect
	github.com/axw/gocov v1.0.0 // indirect
	github.com/containerd/containerd v1.4.3
	github.com/deislabs/oras v0.10.0
	github.com/jstemmer/go-junit-report v0.9.1 // indirect
	github.com/mitchellh/go-homedir v1.1.0
	github.com/opencontainers/image-spec v1.0.1
	github.com/spf13/cobra v1.1.3
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.7.0
	github.com/t-yuki/gocover-cobertura v0.0.0-20180217150009-aaee18c8195c // indirect
	golang.org/x/lint v0.0.0-20201208152925-83fdc39ff7b5 // indirect
	golang.org/x/tools v0.1.0 // indirect
)
