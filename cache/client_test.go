package cache

import "testing"

func TestNewNoTLSByDefault(t *testing.T) {
	c := New(Config{Address: "localhost:6379"})
	defer c.Close()

	if c.Options().TLSConfig != nil {
		t.Fatal("expected no TLSConfig when TLS is disabled")
	}
}

func TestNewTLSEnabled(t *testing.T) {
	c := New(Config{
		Address: "example.db.ondigitalocean.com:25061",
		TLS:     TLSConfig{Enabled: true},
	})
	defer c.Close()

	tc := c.Options().TLSConfig
	if tc == nil {
		t.Fatal("expected TLSConfig when TLS is enabled")
	}

	if tc.ServerName != "example.db.ondigitalocean.com" {
		t.Fatalf("expected ServerName derived from Address, got %q", tc.ServerName)
	}

	if tc.InsecureSkipVerify {
		t.Fatal("expected certificate verification to stay on")
	}
}

func TestNewTLSServerNameOverride(t *testing.T) {
	c := New(Config{
		Address: "10.0.0.5:25061",
		TLS:     TLSConfig{Enabled: true, ServerName: "valkey.internal"},
	})
	defer c.Close()

	if got := c.Options().TLSConfig.ServerName; got != "valkey.internal" {
		t.Fatalf("expected explicit ServerName to win, got %q", got)
	}
}
