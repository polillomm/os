package valueObject

import "testing"

func TestSslHostname(t *testing.T) {
	t.Run("ValidSslHostname", func(t *testing.T) {
		validSslHostname := []interface{}{
			"127.0.0.1", "8.8.8.8", "google.com", "goinfinite.net", "*.goinfinite.net",
			"*.mysite.com",
		}

		for _, sslHostname := range validSslHostname {
			_, err := NewSslHostname(sslHostname)
			if err != nil {
				t.Errorf(
					"Expected no error for '%v', got '%s'", sslHostname, err.Error(),
				)
			}
		}
	})

	t.Run("InvalidSslHostname", func(t *testing.T) {
		invalidSslHostname := []interface{}{
			"", "127.0.0.1.0", "8.8.8", "https://google.com", "http://goinfinite.net",
		}

		for _, sslHostname := range invalidSslHostname {
			_, err := NewSslHostname(sslHostname)
			if err == nil {
				t.Errorf("Expected error for '%v', got nil", sslHostname)
			}
		}
	})
}
