package token

import "time"

// Maker is an interface for creating and parsing tokens.
type Maker interface {
	CreateToken(username string, duration time.Duration) (string, error)
	ValidateToken(token string) (*Payload, error)
}
