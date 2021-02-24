package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/cray-hpe/csmoci/pkg/version"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile, repositoryUrl, versionOutput string
var versionSimple, versionGit bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "csmoci",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version",
	Run: func(cmd *cobra.Command, args []string) {
		clientVersion := version.Get()
		if versionSimple {
			fmt.Printf("%v.%v\n", clientVersion.Major, clientVersion.Minor)
			os.Exit(0)
		}
		if versionGit {
			fmt.Println(clientVersion.GitCommit)
			os.Exit(0)
		}
		switch output := versionOutput; output {
		case "pretty":
			fmt.Printf("%-15s: %s\n", "Build Commit", clientVersion.GitCommit)
			fmt.Printf("%-15s: %s\n", "Build Time", clientVersion.BuildDate)
			fmt.Printf("%-15s: %s\n", "Go Version", clientVersion.GoVersion)
			fmt.Printf("%-15s: %s\n", "Git Version", clientVersion.GitVersion)
			fmt.Printf("%-15s: %s\n", "Platform", clientVersion.Platform)
			fmt.Printf("%-15s: %v.%v.%v\n", "App. Version", clientVersion.Major, clientVersion.Minor, clientVersion.FixVr)
		case "json":
			b, err := json.Marshal(clientVersion)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(string(b))
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.csmoci.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.AddCommand(versionCmd)
	rootCmd.PersistentFlags().StringVarP(&repositoryUrl, "repository", "r", "localhost:5000", "OCI Repository URL")

	versionCmd.Flags().StringVarP(&versionOutput, "output", "o", "pretty", "output format pretty,json")
	versionCmd.Flags().BoolVarP(&versionSimple, "simple", "s", false, "Simple version on a single line")
	versionCmd.Flags().BoolVarP(&versionGit, "git", "g", false, "Simple commit sha of the source tree on a single line. \"-dirty\" added to the end if uncommitted changes present")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".csmoci" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".csmoci")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
