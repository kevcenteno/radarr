package radarr

import "testing"

func TestError(t *testing.T) {
	e := Error{
		Code:    1234,
		Message: "foo",
	}
	expectedMessage := "Radarr error: code 1234, message 'foo'"

	if e.Error() != expectedMessage {
		t.Errorf("Message should be the same. Got '%s', want '%s'", e.Error(), expectedMessage)
	}
}
