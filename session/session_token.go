package session

// Session token owned by the client holding the session
type SessionToken struct {
	// The id of the session.
	SessionId string `json:"sessionId"`
	// The node hosting the session.
	Host string `json:"host"`
}

func NewSessionToken(sessionId string, host string) SessionToken {
	return SessionToken{
		SessionId: sessionId,
		Host:      host,
	}
}
