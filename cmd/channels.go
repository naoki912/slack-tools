package cmd

import (
	"log"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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

	flags := cmd.Flags()
	flags.BoolP("exclude-archived", "", false, "excludeArchived")

	viper.BindPFlag("slack.channels.list.exclude-archived", flags.Lookup("exclude-archived"))

	return cmd
}

func runChannelsListCmd(cmd *cobra.Command, args []string) {

	api := newSlackApi()

	channels, err := api.GetChannels(viper.GetBool("slack.channels.list.exclude-archived"))

	if err != nil {
		log.Fatal(err)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"id",
		"name",
		"is_archived",
		"is_member",
		"num_members",
	})

	for _, channel := range channels {
		id := channel.ID
		name := channel.Name
		isArchived := channel.IsArchived
		isMember := channel.IsMember
		numMembers := len(channel.Members)

		table.Append([]string{
			id,
			name,
			strconv.FormatBool(isArchived),
			strconv.FormatBool(isMember),
			strconv.Itoa(numMembers),
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

	var channelName string
	var channelId string

	if len(args) > 0 {
		channelName = args[0]
	}

	api := newSlackApi()

	channels, err := api.GetChannels(true)
	if err != nil {
		log.Fatal(err)
	}

	for i := range channels {
		if channels[i].Name == channelName {
			channelId = channels[i].ID
			break
		}
	}

	channel, err := api.GetChannelInfo(channelId)
	if err != nil {
		log.Fatal(err)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"property",
		"value",
	})

	table.Append([]string{
		"ID",
		channel.ID,
	})
	table.Append([]string{
		"Name",
		channel.Name,
	})
	table.Append([]string{
		"Created",
		channel.Created.String(),
	})
	table.Append([]string{
		"Creator",
		channel.Creator,
	})
	table.Append([]string{
		"IsArchived",
		strconv.FormatBool(channel.IsArchived),
	})
	table.Append([]string{
		"IsMember",
		strconv.FormatBool(channel.IsMember),
	})
	table.Append([]string{
		"NumMembers",
		strconv.Itoa(len(channel.Members)),
	})
	table.Append([]string{
		"Topic:Value",
		channel.Topic.Value,
	})
	table.Append([]string{
		"Topic:Creator",
		channel.Topic.Creator,
	})
	table.Append([]string{
		"Topic:LastSet",
		channel.Topic.LastSet.String(),
	})
	table.Append([]string{
		"Purpose:Value",
		channel.Purpose.Value,
	})
	table.Append([]string{
		"Purpose:Creator",
		channel.Purpose.Creator,
	})
	table.Append([]string{
		"Purpose:LastSet",
		channel.Purpose.LastSet.String(),
	})

	table.Render()
}
