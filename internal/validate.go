package validate

import (
	"fmt"
	"net/http"
)

func IsValidHttpRequest(incomingRequest *http.Request, expectedMethod string) error {
	if incomingRequest.Method != expectedMethod {
		return fmt.Errorf("error checking validity of request: got %s, expected %s", incomingRequest.Method, expectedMethod)
	}
	return nil
}
