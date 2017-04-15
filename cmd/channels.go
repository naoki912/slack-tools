package cmd

import (
	"fmt"
	"github.com/antonholmquist/jason"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func init() {
	RootCmd.AddCommand(newChannelsCmd())
}

func newChannelsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "channel",
		Short: "channels api",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	cmd.AddCommand(
		newChannelsListCmd(),
		newChannelsShowCmd(),
	)

	return cmd
}

func newChannelsListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "channels.list api",
		Run:   runChannelsListCmd,
	}

	return cmd
}

func runChannelsListCmd(cmd *cobra.Command, args []string) {
	param := url.Values{}
	param.Add("token", viper.GetString("slack.token"))

	res, err := http.PostForm("https://slack.com/api/channels.list", param)
	if err != nil {
		fmt.Fprintf(os.Stderr, "HTTP Error: %s\n", err)
		os.Exit(1)
	}
	jsonBody, _ := jason.NewObjectFromReader(res.Body)

	ok, _ := jsonBody.GetBoolean("ok")
	if ok == false {
		e, _ := jsonBody.GetString("error")
		fmt.Fprintf(os.Stderr, "API Error: %s\n", e)
		os.Exit(1)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"id",
		"name",
		"is_archived",
		"is_member",
		"num_members",
	})

	channels, _ := jsonBody.GetObjectArray("channels")

	for _, value := range channels {
		id, _ := value.GetString("id")
		name, _ := value.GetString("name")
		is_archived, _ := value.GetBoolean("is_archived")
		is_member, _ := value.GetBoolean("is_member")
		num_members, _ := value.GetInt64("num_members")

		table.Append([]string{
			id,
			name,
			strconv.FormatBool(is_archived),
			strconv.FormatBool(is_member),
			strconv.FormatInt(num_members, 10),
		})
	}

	table.Render()
}

func newChannelsShowCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show",
		Short: "channels.info api",
		Run:   runChannelsShowCmd,
	}

	return cmd
}

func runChannelsShowCmd(cmd *cobra.Command, args []string) {
	param := url.Values{}
	param.Add("token", viper.GetString("slack.token"))

	if len(args) > 0 {
		param.Add("channel", args[0])
	}

	res, err := http.PostForm("https://slack.com/api/channels.info", param)
	if err != nil {
		fmt.Fprintf(os.Stderr, "HTTP Error: %s\n", err)
		os.Exit(1)
	}
	jsonBody, _ := jason.NewObjectFromReader(res.Body)

	ok, _ := jsonBody.GetBoolean("ok")
	if ok == false {
		e, _ := jsonBody.GetString("error")
		fmt.Fprintf(os.Stderr, "API Error: %s\n", e)
		os.Exit(1)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"property",
		"value",
	})

	channel, _ := jsonBody.GetObject("channel")

	for property, value := range channel.Map() {
		var stringValue string

		switch property {

		case "created", "unread_count", "unread_count_display":
			i, _ := value.Int64()
			stringValue = strconv.FormatInt(i, 10)
			fmt.Fprintf(os.Stdout, "%v", stringValue)

		case "is_archived", "is_general", "is_member", "is_starred":
			b, _ := value.Boolean()
			stringValue = strconv.FormatBool(b)

		case "members":
			var members []string
			a, _ := value.Array()
			for _, member := range a {
				m, _ := member.String()
				members = append(members, m)
			}
			stringValue = strings.Join(members[:], ", ")

		case "topic", "purpose", "latest":
			a, _ := value.Object()
			stringValue = a.String()

		default:
			stringValue, _ = value.String()
		}

		table.Append([]string{
			property,
			stringValue,
		})
	}

	table.Render()
}
