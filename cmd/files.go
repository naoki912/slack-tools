package cmd

import (
	"fmt"
	"github.com/antonholmquist/jason"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func init() {
	RootCmd.AddCommand(newFilesCmd())
}

func newFilesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "file",
		Short: "files api",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	cmd.AddCommand(
		newFilesUploadCmd(),
	)

	return cmd
}

func newFilesUploadCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upload",
		Short: "files.upload api",
		Run:   runFilesUploadCmd,
	}

	flags := cmd.Flags()
	flags.StringP("channels", "c", "", "channels")
	flags.StringP("filename", "f", "", "filename")
	flags.StringP("type", "t", "", "file type")

	viper.BindPFlag("slack.files.upload.channels", flags.Lookup("channels"))
	viper.BindPFlag("slack.files.upload.filename", flags.Lookup("filename"))
	viper.BindPFlag("slack.files.upload.type", flags.Lookup("type"))

	return cmd
}

func runFilesUploadCmd(cmd *cobra.Command, args []string) {
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "stdin read error: %s", err)
		os.Exit(1)
	}
	Content := string(b)

	param := url.Values{}
	param.Add("content", Content)
	param.Add("token", viper.GetString("slack.token"))
	param.Add("channels", viper.GetString("slack.files.upload.channels"))
	param.Add("filename", viper.GetString("slack.files.upload.type"))

	resp, err := http.PostForm("https://slack.com/api/files.upload", param)
	if err != nil {
		fmt.Fprintf(os.Stderr, "HTTP Error: %s\n", err)
		os.Exit(1)
	}
	jsonBody, _ := jason.NewObjectFromReader(resp.Body)

	ok, _ := jsonBody.GetBoolean("ok")
	if ok == false {
		e, _ := jsonBody.GetString("error")
		fmt.Fprintf(os.Stderr, "API Error: %s\n", e)
		os.Exit(1)
	}
}
