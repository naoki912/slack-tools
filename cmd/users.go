package cmd

import (
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

func init() {
	RootCmd.AddCommand(newUsersCmd())
}

func newUsersCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user",
		Short: "users",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	cmd.AddCommand(
		newUsersListCmd(),
	)

	return cmd
}

func newUsersListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "users list",
		RunE:  runUsersListCmd,
	}

	return cmd
}

func runUsersListCmd(cmd *cobra.Command, args []string) error {
	api := newSlackApi()
	users, err := api.GetUsers()

	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"id",
		"name",
		"real name",
		"email",
		"skype",
		"is admin",
		"is owner",
	})

	for _, user := range users {
		table.Append([]string{
			user.ID,
			user.Name,
			user.RealName,
			user.Profile.Email,
			user.Profile.Skype,
			strconv.FormatBool(user.IsAdmin),
			strconv.FormatBool(user.IsOwner),
		})
	}

	table.Render()

	return err
}
