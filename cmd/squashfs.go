/*
Copyright 2021 Hewlett Packard Enterprise Development LP
*/

package cmd

import (
	"context"
	"fmt"

	"github.com/containerd/containerd/remotes/docker"
	"github.com/cray-hpe/csmoci/pkg/squashfs"
	"github.com/spf13/cobra"
)

var pushCompression string

// squashfsCmd represents the squashfs command
var squashfsCmd = &cobra.Command{
	Use:   "squashfs",
	Short: "push and pull squashfs filesystems to OCI repositories",
}

// pushCmd represents the push command
var squashfsPushCmd = &cobra.Command{
	Use:   "push <filename> <reference>",
	Short: "push squashfs filesystem to an OCI registry",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		resolver := docker.NewResolver(docker.ResolverOptions{})
		desc, err := squashfs.PushSquashFS(ctx, resolver, args[0], args[1], pushCompression)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Pushed to %s with digest %s\n", args[1], desc.Digest)
	},
}

// pushCmd represents the push command
var squashfsPullCmd = &cobra.Command{
	Use:   "pull <reference> <directory>",
	Short: "Pull squashfs filesystems from OCI registry",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("squashfs pull called with args: ", args)
		ctx := context.Background()
		resolver := docker.NewResolver(docker.ResolverOptions{})
		desc, _, err := squashfs.PullSquashFS(ctx, resolver, args[0], args[1])
		if err != nil {
			panic(err)
		}
		fmt.Printf("Pulled from %s with digest %s\n", args[0], desc.Digest)
	},
}

func init() {
	rootCmd.AddCommand(squashfsCmd)
	squashfsCmd.AddCommand(squashfsPushCmd)
	squashfsCmd.AddCommand(squashfsPullCmd)
	squashfsPushCmd.Flags().StringVarP(&pushCompression, "compression", "c", "gzip", fmt.Sprintf("Compression Algorithm"))
}
