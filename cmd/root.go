package cmd

import (
	"os"

	"github.com/jacobmeredith/swarm/internal/collections"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "swarm",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		cd := cmd.Flag("collection-directory").Value.String()
		fn := cmd.Flag("file-name").Value.String()
		rn := cmd.Flag("request-name").Value.String()

		err := collections.RunCollectionRequest(cd, fn, rn)
		if err != nil {
			cmd.PrintErr(err)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("collection-directory", "c", "collections", "The directory where collections are stored")
	rootCmd.Flags().StringP("file-name", "f", "", "The file name of the collection")
	rootCmd.Flags().StringP("request-name", "n", "", "The name of the request to run")

	rootCmd.MarkFlagRequired("file-name")
	rootCmd.MarkFlagRequired("request-name")
}
