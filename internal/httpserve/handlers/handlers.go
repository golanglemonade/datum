package handlers

import (
	echo "github.com/datumforge/echox"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"

	ent "github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/httpserve/authmanager"
	"github.com/datumforge/datum/pkg/analytics"
	"github.com/datumforge/datum/pkg/events/kafka/publisher"
	"github.com/datumforge/datum/pkg/sessions"
	"github.com/datumforge/datum/pkg/tokens"
	"github.com/datumforge/datum/pkg/utils/emails"
	"github.com/datumforge/datum/pkg/utils/marionette"
	"github.com/datumforge/datum/pkg/utils/totp"
)

// Handler contains configuration options for handlers
type Handler struct {
	// IsTest is a flag to determine if the application is running in test mode and will mock external calls
	IsTest bool
	// DBClient to interact with the generated ent schema
	DBClient *ent.Client
	// RedisClient to interact with redis
	RedisClient *redis.Client
	// AuthManager contains the required configuration for the auth session creation
	AuthManager *authmanager.Config
	// TokenManager contains the token manager in order to validate auth requests
	TokenManager *tokens.TokenManager
	// Logger provides the zap logger to do logging things from the handlers
	Logger *zap.SugaredLogger
	// ReadyChecks is a set of checkFuncs to determine if the application is "ready" upon startup
	ReadyChecks Checks
	// JWTKeys contains the set of valid JWT authentication key
	JWTKeys jwk.Set
	// SessionConfig to handle sessions
	SessionConfig *sessions.SessionConfig
	// EmailManager to handle sending emails
	EmailManager *emails.EmailManager
	// TaskMan manages tasks in a separate goroutine to allow for non blocking operations
	TaskMan *marionette.TaskManager
	// AnalyticsClient is the client used to send analytics events
	AnalyticsClient *analytics.EventManager
	// OauthProvider contains the configuration settings for all supported Oauth2 providers
	OauthProvider OauthProviderConfig
	// AuthMiddleware contains the middleware to be used for authenticated endpoints
	AuthMiddleware []echo.MiddlewareFunc
	// WebAuthn contains the configuration settings for the webauthn provider
	WebAuthn *webauthn.WebAuthn
	// OTPManager contains the configuration settings for the OTP provider
	OTPManager *totp.Manager
	// EventManager contains the configuration settings for the event publisher
	EventManager *publisher.KafkaPublisher
}
