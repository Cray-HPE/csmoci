package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

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
	Use:   "push",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("squashfs push called with args: ", args)
	},
}

// pushCmd represents the push command
var squashfsPullCmd = &cobra.Command{
	Use:   "pull",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("squashfs pull called with args: ", args)
	},
}

func init() {
	rootCmd.AddCommand(squashfsCmd)
	squashfsCmd.AddCommand(squashfsPushCmd)
	squashfsCmd.AddCommand(squashfsPullCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// squashfsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// squashfsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
