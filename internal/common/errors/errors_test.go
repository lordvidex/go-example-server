package errors

import "testing"

func TestGeneralHTTPError_StatusCode(t *testing.T) {
	t.Run("All undefined errors are Internal Server Error by default", func(t *testing.T) {

	})
}

func TestNotFound_StatusMessage(t *testing.T) {
	t.Run("Method impl is correctly generated from embedded generalHTTPError", func(t *testing.T) {
		e := NotFound{}

		if e.StatusMessage() != "Not Found" {
			t.Errorf("StatusMessage() should return Not Found. Expected %s got %s", "Not Found", e.StatusMessage())
		}
	})
}
