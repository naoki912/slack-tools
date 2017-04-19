package cmd

import (
	"github.com/nlopes/slack"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
)

func init() {
	RootCmd.AddCommand(newChatCmd())
}

func newChatCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "chat",
		Short: "chat api",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	cmd.AddCommand(
		newChatPostMessageCmd(),
	)

	return cmd
}

func newChatPostMessageCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "post",
		Short: "chat.postMessage",
		RunE:  runChatPostMessageCmd,
	}

	flags := cmd.Flags()
	flags.BoolP("asuser", "", false, "as_user")

	viper.BindPFlag("slack.chat.postMessage.asUser", flags.Lookup("asuser"))

	return cmd
}

func runChatPostMessageCmd(cmd *cobra.Command, args []string) error {
	Channel := "general"
	Text := ""

	if len(args) > 0 {
		Channel = args[0]
	}

	if len(args) == 1 {
		b, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return err
		}
		Text = string(b)
	} else {
		Text = args[1]
	}

	param := slack.NewPostMessageParameters()
	param.AsUser = viper.GetBool("slack.chat.postMessage.asUser")

	api := newSlackApi()
	_, _, err := api.PostMessage(Channel, Text, param)

	return err
}
