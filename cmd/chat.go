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
	"strconv"
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
		Run:   runChatPostMessageCmd,
	}

	flags := cmd.Flags()
	flags.BoolP("asuser", "", false, "as_user")

	viper.BindPFlag("slack.chat.postMessage.asUser", flags.Lookup("asuser"))

	return cmd
}

func runChatPostMessageCmd(cmd *cobra.Command, args []string) {
	Channel := "general"
	Text := ""

	if len(args) > 0 {
		Channel = args[0]
	}

	if len(args) == 1 {
		b, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "stdin read error: %s", err)
			os.Exit(1)
		}
		Text = string(b)
	} else {
		Text = args[1]
	}

	param := url.Values{}
	param.Add("token", viper.GetString("slack.token"))
	param.Add("text", Text)
	param.Add("channel", Channel)
	param.Add("as_user", strconv.FormatBool(viper.GetBool("slack.chat.postMessage.asUser")))

	resp, err := http.PostForm("https://slack.com/api/chat.postMessage", param)
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
