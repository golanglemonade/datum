package auth

import (
	"context"
	"time"

	echo "github.com/datumforge/echox"
	"github.com/golang-jwt/jwt/v5"

	"github.com/datumforge/datum/pkg/middleware/echocontext"
	"github.com/datumforge/datum/pkg/tokens"
)

// newValidClaims returns claims with a fake subject for testing purposes ONLY
func newValidClaims(subject string) *tokens.Claims {
	iat := time.Now()
	nbf := iat
	exp := time.Now().Add(time.Hour)

	claims := &tokens.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   subject,
			Issuer:    "test suite",
			IssuedAt:  jwt.NewNumericDate(iat),
			NotBefore: jwt.NewNumericDate(nbf),
			ExpiresAt: jwt.NewNumericDate(exp),
		},
		UserID:      subject,
		Email:       "rustys@datum.net",
		OrgID:       "nano_id_of_org",
		ParentOrgID: "nano_id_of_parent_org",
		Tier:        "premium",
	}

	return claims
}

// NewTestEchoContextWithValidUser creates an echo context with a fake subject for testing purposes ONLY
func NewTestEchoContextWithValidUser(subject string) (echo.Context, error) {
	ec := echocontext.NewTestEchoContext()

	claims := newValidClaims(subject)

	ec.Set(ContextUserClaims.name, claims)

	return ec, nil
}

func NewTestContextWithValidUser(subject string) (context.Context, error) {
	ec, err := NewTestEchoContextWithValidUser(subject)
	if err != nil {
		return nil, err
	}

	reqCtx := context.WithValue(ec.Request().Context(), echocontext.EchoContextKey, ec)

	ec.SetRequest(ec.Request().WithContext(reqCtx))

	return reqCtx, nil
}

// newValidClaims returns claims with a fake orgID for testing purposes ONLY
func newValidClaimsOrgID(sub, orgID string) *tokens.Claims {
	iat := time.Now()
	nbf := iat
	exp := time.Now().Add(time.Hour)

	claims := &tokens.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   sub,
			Issuer:    "test suite",
			IssuedAt:  jwt.NewNumericDate(iat),
			NotBefore: jwt.NewNumericDate(nbf),
			ExpiresAt: jwt.NewNumericDate(exp),
		},
		UserID:      sub,
		Email:       "rustys@datum.net",
		OrgID:       orgID,
		ParentOrgID: "nano_id_of_parent_org",
		Tier:        "premium",
	}

	return claims
}

// NewTestEchoContextWithOrgID creates an echo context with a fake orgID for testing purposes ONLY
func NewTestEchoContextWithOrgID(sub, orgID string) (echo.Context, error) {
	ec := echocontext.NewTestEchoContext()

	claims := newValidClaimsOrgID(sub, orgID)

	ec.Set(ContextUserClaims.name, claims)

	return ec, nil
}

func NewTestContextWithOrgID(sub, orgID string) (context.Context, error) {
	ec, err := NewTestEchoContextWithOrgID(sub, orgID)
	if err != nil {
		return nil, err
	}

	reqCtx := context.WithValue(ec.Request().Context(), echocontext.EchoContextKey, ec)

	ec.SetRequest(ec.Request().WithContext(reqCtx))

	return reqCtx, nil
}
