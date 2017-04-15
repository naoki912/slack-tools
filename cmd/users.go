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
)

func init() {
	RootCmd.AddCommand(newUsersCmd())
}

func newUsersCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user",
		Short: "users api",
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
		Short: "users.list api",
		Run:   runUsersListCmd,
	}

	return cmd
}

func runUsersListCmd(cmd *cobra.Command, args []string) {
	param := url.Values{}
	param.Add("token", viper.GetString("slack.token"))

	res, err := http.PostForm("https://slack.com/api/users.list", param)
	if err != nil {
		fmt.Fprintf(os.Stderr, "HTTP Error: %s\n", err)
		os.Exit(-1)
	}
	jsonBody, err := jason.NewObjectFromReader(res.Body)

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
		"real name",
		"email",
		"skype",
		"is admin",
		"is owner",
	})

	members, _ := jsonBody.GetObjectArray("members")

	for _, value := range members {
		id, _ := value.GetString("id")
		name, _ := value.GetString("name")
		real_name, _ := value.GetString("profile", "real_name")
		email, _ := value.GetString("profile", "email")
		skype, _ := value.GetString("profile", "skype")
		is_admin, _ := value.GetBoolean("is_admin")
		is_owner, _ := value.GetBoolean("is_admin")

		table.Append([]string{
			id,
			name,
			real_name,
			email,
			skype,
			strconv.FormatBool(is_admin),
			strconv.FormatBool(is_owner),
		})
	}

	table.Render()
}
