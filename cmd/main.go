package main

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"github.com/mmycin/npminfo/pkg/api"
	"github.com/mmycin/npminfo/pkg/ui"
)

func main() {
	// Define the flag variable
	var packageName string

	// Root command
	var rootCmd = &cobra.Command{
		Use:   "npminfo",
		Short: "Fetch and display detailed information about an NPM package. Created by Mycin.",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Fetch the NPM package info using the -p flag
			if packageName == "" {
				return fmt.Errorf("package name is required, use -p flag to specify")
			}

			// Fetch the NPM package info
			npmInfo, err := api.FetchNpmPackageInfo(packageName)
			if err != nil {
				return err
			}

			// Display the package info in a table
			ui.DisplayNpmPackageInfo(npmInfo)

			return nil
		},
	}

	// Define the -p flag for specifying the package name
	rootCmd.Flags().StringVarP(&packageName, "package", "p", "", "The name of the NPM package")

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
