package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	echo "github.com/datumforge/echox"
	"golang.org/x/oauth2"

	"github.com/datumforge/datum/internal/ent/enums"
	"github.com/datumforge/datum/internal/ent/privacy/token"
	"github.com/datumforge/datum/internal/httpserve/middleware/auth"
	"github.com/datumforge/datum/internal/providers/github"
	"github.com/datumforge/datum/internal/providers/google"
	"github.com/datumforge/datum/internal/sessions"
	"github.com/datumforge/datum/internal/tokens"
)

// OauthTokenRequest to authenticate an oauth user with the Datum Server
type OauthTokenRequest struct {
	Name             string `json:"name"`
	Email            string `json:"email"`
	AuthProvider     string `json:"authProvider"`
	ExternalUserID   string `json:"externalUserId"`
	ExternalUserName string `json:"externalUserName"`
	ClientToken      string `json:"clientToken"`
}

// OauthRegister returns the TokenResponse for a verified authenticated external oauth user
func (h *Handler) OauthRegister(ctx echo.Context) error {
	var r OauthTokenRequest

	// parse request body
	if err := json.NewDecoder(ctx.Request().Body).Decode(&r); err != nil {
		return ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
	}

	ctxWithToken := token.NewContextWithOauthTooToken(ctx.Request().Context(), r.Email)

	// create oauth2 token from rquest input
	tok := &oauth2.Token{
		AccessToken: r.ClientToken,
	}

	// verify the token provided to ensure the user is valid
	if err := h.verifyClientToken(ctxWithToken, r.AuthProvider, tok, r.Email); err != nil {
		return ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
	}

	// check if users exists and create if not, updates last seen of existing user
	user, err := h.CheckAndCreateUser(ctxWithToken, r.Name, r.Email, enums.AuthProvider(strings.ToUpper(r.AuthProvider)))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
	}

	// create claims for verified user
	claims := createClaims(user)

	access, refresh, err := h.TM.CreateTokenPair(claims)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
	}

	// set cokies for the user
	auth.SetAuthCookies(ctx.Response().Writer, access, refresh)

	setSessionMap := map[string]string{}
	setSessionMap[sessions.ExternalUserIDKey] = fmt.Sprintf("%v", r.ExternalUserID)
	setSessionMap[sessions.UsernameKey] = r.ExternalUserName
	setSessionMap[sessions.UserTypeKey] = r.AuthProvider
	setSessionMap[sessions.EmailKey] = r.Email
	setSessionMap[sessions.UserIDKey] = user.ID

	if err := h.SessionConfig.SaveAndStoreSession(ctx.Request().Context(), ctx.Response().Writer, setSessionMap, user.ID); err != nil {
		return ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
	}

	// Return the access token
	return ctx.JSON(http.StatusOK, tokens.TokenResponse{
		AccessToken:  access,
		RefreshToken: refresh,
		TokenType:    "Bearer",
		ExpiresIn:    claims.ExpiresAt.Unix(),
	})
}

// verifyClientToken verifies the provided access token from an external oauth2 provider is valid and matches the user's email
// supported providers are Github and Google
func (h *Handler) verifyClientToken(ctx context.Context, provider string, token *oauth2.Token, email string) error {
	switch strings.ToLower(provider) {
	case githubProvider:
		config := h.getGithubOauth2Config()
		cc := github.Config{IsEnterprise: false, IsMock: h.IsTest}

		return github.VerifyClientToken(ctx, token, config, email, &cc)
	case googleProvider:
		config := h.getGoogleOauth2Config()
		return google.VerifyClientToken(ctx, token, config, email)
	default:
		return ErrInvalidProvider
	}
}