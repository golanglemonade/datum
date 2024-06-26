package hooks

import (
	"context"
	"time"

	"entgo.io/ent"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/hook"
	"github.com/datumforge/datum/pkg/auth"
)

const (
	redacted = "*****************************"
)

// HookCreatePersonalAccessToken runs on accesstoken mutations and sets expires and owner id
func HookCreatePersonalAccessToken() ent.Hook {
	return hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.PersonalAccessTokenFunc(func(ctx context.Context, mutation *generated.PersonalAccessTokenMutation) (generated.Value, error) {
			// default the expiration to 7 days
			expires, ok := mutation.ExpiresAt()
			if !ok || expires.IsZero() {
				mutation.SetExpiresAt(time.Now().Add(time.Hour * 24 * 7)) //nolint:mnd
			}

			userID, err := auth.GetUserIDFromContext(ctx)
			if err != nil {
				return nil, err
			}

			// set user on the token
			mutation.SetOwnerID(userID)

			return next.Mutate(ctx, mutation)
		})
	}, ent.OpCreate)
}

// HookUpdatePersonalAccessToken runs on accesstoken update and redacts the token
func HookUpdatePersonalAccessToken() ent.Hook {
	return hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.PersonalAccessTokenFunc(func(ctx context.Context, mutation *generated.PersonalAccessTokenMutation) (generated.Value, error) {
			// do not allow user to be changed
			_, ok := mutation.OwnerID()
			if ok {
				mutation.ClearOwner()
			}

			retVal, err := next.Mutate(ctx, mutation)
			if err != nil {
				return nil, err
			}

			// redact the token
			pat, ok := retVal.(*generated.PersonalAccessToken)
			if !ok {
				return retVal, nil
			}

			pat.Token = redacted

			return pat, nil
		})
	}, ent.OpUpdate|ent.OpUpdateOne)
}
