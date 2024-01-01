package cmd

import (
	"errors"
	"os"

	"github.com/jacobmeredith/swarm/internal/requests"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "swarm",
	Short: "A command line utitlity to make HTTP requests",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		collection_directory := cmd.Flag("collection-directory").Value.String()
		file_name := cmd.Flag("file-name").Value.String()
		request_name := cmd.Flag("request-name").Value.String()

		if collection_directory != "" && file_name != "" && request_name != "" {
			return
		}

		method := cmd.Flag("method").Value.String()
		url := cmd.Flag("url").Value.String()

		if method == "" {
			cmd.PrintErr(errors.New("No method provided"))
			return
		}

		if url == "" {
			cmd.PrintErr(errors.New("No URL provided"))
			return
		}

		content_type := cmd.Flag("content-type").Value.String()
		body := cmd.Flag("body").Value.String()
		headers := cmd.Flag("headers").Value.String()
		cookies := cmd.Flag("cookies").Value.String()

		request, err := requests.NewRequest(requests.RequestConfig{
			Url:         url,
			Method:      method,
			ContentType: content_type,
			Body:        body,
			Headers:     headers,
			Cookies:     cookies,
		})

		if err != nil {
			cmd.PrintErr(err)
		}

		_, err = request.Build()
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
	// collection flags
	rootCmd.Flags().StringP("collection-directory", "c", "collections", "The directory where collections are stored")
	rootCmd.Flags().StringP("file-name", "f", "", "The file name of the collection")
	rootCmd.Flags().StringP("request-name", "n", "", "The name of the request to run")

	// Individual request flags
	rootCmd.Flags().StringP("method", "m", "", "Method for request")
	rootCmd.Flags().StringP("url", "u", "", "URL for request")
	rootCmd.Flags().String("content-type", "", "Content type of the request body")
	rootCmd.Flags().StringP("body", "b", "", "Body in string format")
	rootCmd.Flags().String("headers", "", "Headers in following format \"key:value,key2:value2\"")
	rootCmd.Flags().String("cookies", "", "Cookies in following format \"key:value,key2:value2\"")
}
