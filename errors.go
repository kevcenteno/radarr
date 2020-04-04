package radarr

import "fmt"

// Error Radarr error response
type Error struct {
	Code    int
	Message string
}

func (e *Error) Error() string {
	return fmt.Sprintf("Radarr error: code %d, message '%s'", e.Code, e.Message)
}
