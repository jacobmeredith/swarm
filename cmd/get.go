package cmd

import (
	"github.com/jacobmeredith/swarm/internal/requests"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Make a get request to a URL",
	Long: `This command allows you to make a GET request to a specified URL. For example:
swarm get -u https://google.com`,
	Run: func(cmd *cobra.Command, args []string) {
		err := requests.Get(cmd.Flag("url").Value.String())
		if err != nil {
			cmd.PrintErr(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.Flags().StringP("url", "u", "", "URL to get")

	getCmd.MarkFlagRequired("url")
}
