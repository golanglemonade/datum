package interceptors

import (
	"context"

	"entgo.io/ent"

	"github.com/datumforge/datum/internal/ent/generated/intercept"
	"github.com/datumforge/datum/internal/ent/generated/orgmembership"
	"github.com/datumforge/datum/internal/ent/generated/privacy"
	"github.com/datumforge/datum/pkg/auth"
)

// InterceptorOrgMembers is middleware to change the Org Members query
func InterceptorOrgMembers() ent.Interceptor {
	return intercept.TraverseFunc(func(ctx context.Context, q intercept.Query) error {
		// bypass filter if the request is internal and already set to allowed
		if _, allow := privacy.DecisionFromContext(ctx); allow {
			return nil
		}

		// Organization list queries should not be filtered by organization id
		ctxQuery := ent.QueryFromContext(ctx)
		if ctxQuery.Type == "Organization" {
			return nil
		}

		orgIDs, err := auth.GetOrganizationIDsFromContext(ctx)
		if err != nil {
			return err
		}

		// get all parent orgs to ensure we get all users in the org tree
		allOrgsIDs, err := getAllParentOrgIDs(ctx, orgIDs)
		if err != nil {
			return err
		}

		// sets the organization id on the query for the current organization
		q.WhereP(orgmembership.OrganizationIDIn(allOrgsIDs...))

		return nil
	})
}
