package store

import (
	"database/sql"
	"errors"
	"path/filepath"
	"testing"
)

// TestBug11_TideGate verifies that a transaction surfaces the callback error
// instead of committing partial work. The injected defect ignores the
// callback result and always commits, so the caller loses the error identity.
func TestBug11_TideGate(t *testing.T) {
	s, e := Open(filepath.Join(t.TempDir(), "db.sqlite"))
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	want := errors.New("disk full")
	err := s.Transaction(func(tx *sql.Tx) error {
		return want
	})
	if !errors.Is(err, want) {
		t.Fatalf("transaction lost error identity: %v", err)
	}
}