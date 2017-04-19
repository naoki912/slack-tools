package cmd

import (
	"github.com/nlopes/slack"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
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
		RunE:  runFilesUploadCmd,
	}

	flags := cmd.Flags()
	flags.StringSliceP("channels", "c", []string{""}, "channels")
	flags.StringP("filename", "f", "", "filename")
	flags.StringP("type", "t", "", "file type")

	viper.BindPFlag("slack.files.upload.channels", flags.Lookup("channels"))
	viper.BindPFlag("slack.files.upload.filename", flags.Lookup("filename"))
	viper.BindPFlag("slack.files.upload.type", flags.Lookup("type"))

	return cmd
}

func runFilesUploadCmd(cmd *cobra.Command, args []string) error {
	api := newSlackApi()
	param := slack.FileUploadParameters{}

	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return err
	}

	param.Content = string(b)
	param.Channels = viper.GetStringSlice("slack.files.upload.channels")
	param.Filetype = viper.GetString("slack.files.upload.type")

	_, err = api.UploadFile(param)

	return err
}
