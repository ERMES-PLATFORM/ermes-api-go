package session

import (
	"context"

	"github.com/ERMES-PLATFORM/api-ermes-go/infrastructure"
)

// Commands to create a new session.
type CreateSessionCommands interface {
	// Creates a new session and returns the id of the session.
	CreateSession(
		ctx context.Context,
		opt CreateSessionOptions,
	) (string, error)
	// Returns the ids of the sessions.
	ScanSessions(
		ctx context.Context,
		cursor uint64,
		count int64,
	) (ids []string, newCursor uint64, err error)
}

// Options that defines how a session is created.
type CreateSessionOptions struct {
	// The geographic coordinates associated with the client that owns the session.
	// If nil, the client sessionLocation is initially approximated to the sessionLocation of
	// the node that creates the session. Default is nil.
	clientGeoCoordinates *infrastructure.GeoCoordinates
	// The expiration time is expressed as a Unix timestamp (UTC). If the
	// expiration time is nil, the session does not expire. Default is nil.
	expiresAt *int64
	// Optional session ID. If nil, a new session ID is generated. Default is nil.
	sessionId *string
}
