package datumgroup

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/utils/cli/tables"
)

var groupGetCmd = &cobra.Command{
	Use:   "get",
	Short: "get an existing new datum group",
	RunE: func(cmd *cobra.Command, args []string) error {
		return getGroup(cmd.Context())
	},
}

func init() {
	groupCmd.AddCommand(groupGetCmd)

	groupGetCmd.Flags().StringP("id", "i", "", "group id to query")
}

func getGroup(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	// filter options
	gID := datum.Config.String("id")

	writer := tables.NewTableWriter(groupCmd.OutOrStdout(), "ID", "Name", "Description", "Visibility", "Organization", "Members")

	// if an group ID is provided, filter on that group, otherwise get all
	if gID != "" {
		group, err := client.GetGroupByID(ctx, gID)
		cobra.CheckErr(err)

		if datum.OutputFormat == datum.JSONOutput {
			s, err := json.Marshal(group)
			if err != nil {
				return err
			}

			return datum.JSONPrint(s)
		}

		writer.AddRow(group.Group.ID, group.Group.DisplayName, *group.Group.Description, group.Group.Setting.Visibility, group.Group.Owner.DisplayName, len(group.Group.Members))
		writer.Render()

		return nil
	}

	// get all groups, will be filtered for the authorized organization(s)
	groups, err := client.GetAllGroups(ctx)
	cobra.CheckErr(err)

	if datum.OutputFormat == datum.JSONOutput {
		s, err := json.Marshal(groups)
		cobra.CheckErr(err)

		return datum.JSONPrint(s)
	}

	for _, g := range groups.Groups.Edges {
		writer.AddRow(g.Node.ID, g.Node.DisplayName, *g.Node.Description, g.Node.Setting.Visibility, g.Node.Owner.DisplayName, len(g.Node.Members))
	}

	writer.Render()

	return nil
}
