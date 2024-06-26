package datumapitokens

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
)

var apiTokenDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an existing datum api token token",
	RunE: func(cmd *cobra.Command, args []string) error {
		return deleteAPIToken(cmd.Context())
	},
}

func init() {
	apiTokenCmd.AddCommand(apiTokenDeleteCmd)

	apiTokenDeleteCmd.Flags().StringP("id", "i", "", "api token id to delete")
}

func deleteAPIToken(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	var s []byte

	tokenID := datum.Config.String("id")
	if tokenID == "" {
		return datum.NewRequiredFieldMissingError("token id")
	}

	o, err := client.DeleteAPIToken(ctx, tokenID)
	cobra.CheckErr(err)

	s, err = json.Marshal(o)
	cobra.CheckErr(err)

	return datum.JSONPrint(s)
}
