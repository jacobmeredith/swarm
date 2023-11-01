package cmd

import (
	"github.com/jacobmeredith/swarm/internal/requests"
	"github.com/spf13/cobra"
)

// postCmd represents the post command
var postCmd = &cobra.Command{
	Use:   "post",
	Short: "Make a post request to a URL",
	Long: `This command allows you to make a POST request to a specified URL. For example:
	swarm get -u https://google.com --content-type application/json --body="{\"test\": true}"`,
	Run: func(cmd *cobra.Command, args []string) {
		err := requests.Post(cmd.Flag("url").Value.String(), cmd.Flag("content-type").Value.String(), cmd.Flag("body").Value.String())
		if err != nil {
			cmd.PrintErr(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(postCmd)

	postCmd.Flags().StringP("url", "u", "", "URL to get")
	postCmd.Flags().String("content-type", "", "Content type of the request body")
	postCmd.Flags().StringP("body", "b", "", "Body in string format")

	postCmd.MarkFlagRequired("url")
	postCmd.MarkFlagRequired("content-type")
	postCmd.MarkFlagRequired("body")
}
