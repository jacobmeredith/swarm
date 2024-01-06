package cmd

import (
	"bufio"
	"net/http"
	"os"

	"github.com/jacobmeredith/swarm/internal/requests"
	"github.com/jacobmeredith/swarm/internal/responses"
	"github.com/jacobmeredith/swarm/internal/runner"
	"github.com/spf13/cobra"
)

type CollectionDirectoryFlags struct {
	directory   string
	fileName    string
	requestName string
}

func (cdf *CollectionDirectoryFlags) isValid() bool {
	return cdf.fileName != "" && cdf.directory != "" && cdf.requestName != ""
}

type NormalFlags struct {
	method      string
	url         string
	contentType string
	body        string
	headers     string
	cookies     string
}

type Flags struct {
	collectionFlags CollectionDirectoryFlags
	normalFlags     NormalFlags
}

func getFlags(cmd *cobra.Command) *Flags {
	return &Flags{
		collectionFlags: CollectionDirectoryFlags{
			directory:   cmd.Flag("collection-directory").Value.String(),
			fileName:    cmd.Flag("file-name").Value.String(),
			requestName: cmd.Flag("request-name").Value.String(),
		},
		normalFlags: NormalFlags{
			method:      cmd.Flag("method").Value.String(),
			url:         cmd.Flag("url").Value.String(),
			contentType: cmd.Flag("content-type").Value.String(),
			body:        cmd.Flag("body").Value.String(),
			headers:     cmd.Flag("headers").Value.String(),
			cookies:     cmd.Flag("cookies").Value.String(),
		},
	}
}

func execute(request *requests.Request) error {
	runner := runner.NewRunner(&http.Client{})
	formatter := responses.NewDefaultResponseFormatter()

	runner.SetResponseFormatter(formatter)

	formatted_response, err := runner.Run(request)
	if err != nil {
		return err
	}

	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(formatted_response)
	writer.Flush()

	println(formatted_response)

	return nil
}

func command(cmd *cobra.Command, args []string) {
	flags := getFlags(cmd)

	if flags.collectionFlags.isValid() {
		// Run collection request
		collection, err := requests.NewCollection(flags.collectionFlags.directory, flags.collectionFlags.fileName)
		if err != nil {
			cmd.PrintErr(err)
			return
		}

		request, err := collection.TransformRequest(flags.collectionFlags.requestName)
		if err != nil {
			cmd.PrintErr(err)
			return
		}

		if err := execute(request); err != nil {
			cmd.PrintErr(err)
		}

		return
	}

	// Run adhoc request
	request, err := requests.NewRequest(requests.RequestConfig{
		Url:         flags.normalFlags.url,
		Method:      flags.normalFlags.method,
		ContentType: flags.normalFlags.contentType,
		Body:        flags.normalFlags.body,
		Headers:     flags.normalFlags.headers,
		Cookies:     flags.normalFlags.cookies,
	})
	if err != nil {
		cmd.PrintErr(err)
	}

	if err := execute(request); err != nil {
		cmd.PrintErr(err)
	}
}

var rootCmd = &cobra.Command{
	Use:   "swarm",
	Short: "A command line utitlity to make HTTP requests",
	Long:  ``,
	Run:   command,
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
