/*
Copyright © 2023 @GeekOpsUA
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Add app version declaration
var appVersion = "Version"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Long: `Print the version number of the application.
	You can use it to check if you have the latest version of the application.
	For example:
	$ geekbot version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(appVersion)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
