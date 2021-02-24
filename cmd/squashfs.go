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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("squashfs called with args:", args)
	},
}

// pushCmd represents the push command
var squashfsPushCmd = &cobra.Command{
	Use:   "push <filename> <reference>",
	Short: "A brief description of your command",
	Args:  cobra.MinimumNArgs(2),
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("squashfs push called with args: ", args)
		fmt.Println("repository url: ", repositoryUrl)
		fmt.Println("Push Compression is set to: ", pushCompression)
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
	Short: "A brief description of your command",
	Args:  cobra.MinimumNArgs(2),
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("squashfs pull called with args: ", args)
		ctx := context.Background()
		resolver := docker.NewResolver(docker.ResolverOptions{})
		desc, _, err := squashfs.PullSquashFS(ctx, resolver, args[1], args[0])
		if err != nil {
			panic(err)
		}
		fmt.Printf("Pulled from %s with digest %s\n", args[1], desc.Digest)
	},
}

func init() {
	rootCmd.AddCommand(squashfsCmd)
	squashfsCmd.AddCommand(squashfsPushCmd)
	squashfsCmd.AddCommand(squashfsPullCmd)
	squashfsPushCmd.Flags().StringVarP(&pushCompression, "compression", "c", "gzip", fmt.Sprintf("Compression Algorithm"))

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// squashfsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// squashfsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
