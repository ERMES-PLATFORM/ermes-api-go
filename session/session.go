package session

import (
	"github.com/ERMES-PLATFORM/ermes-api-go/infrastructure"
)

// Session contained in a node and consuming resources.
type Session struct {
	// The id of the session.
	SessionId string `json:"sessionId"`
	// The node hosting the session.
	Host string `json:"host"`
	// Resources consumed by the session
	Resources infrastructure.Resources
}

func NewSessionInstance(sessionId string, host string, resources infrastructure.Resources) Session {
	return Session{
		Host:      host,
		SessionId: sessionId,
		Resources: resources,
	}
}
