package store

import (
	"context"
	"database/sql"
	"encoding/json"
	"os"
	"path/filepath"
	"time"

	_ "modernc.org/sqlite"
)

type Store struct{ db *sql.DB }

func Open(path string) (*Store, error) {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return nil, err
	}
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(1)
	s := &Store{db: db}
	if err := s.schema(); err != nil {
		db.Close()
		return nil, err
	}
	return s, nil
}

func (s *Store) schema() error {
	_, err := s.db.Exec(`CREATE TABLE IF NOT EXISTS records (kind TEXT NOT NULL, id TEXT NOT NULL, payload BLOB NOT NULL, updated_at TEXT NOT NULL, PRIMARY KEY(kind,id)); CREATE TABLE IF NOT EXISTS events (id INTEGER PRIMARY KEY AUTOINCREMENT, subject TEXT NOT NULL, action TEXT NOT NULL, created_at TEXT NOT NULL);`)
	return err
}
func (s *Store) Close() error { return s.db.Close() }
func (s *Store) Save(kind, id string, value any) error {
	raw, err := json.Marshal(value)
	if err != nil {
		return err
	}
	_, err = s.db.Exec(`INSERT INTO records(kind,id,payload,updated_at) VALUES(?,?,?,?) ON CONFLICT(kind,id) DO UPDATE SET payload=excluded.payload,updated_at=excluded.updated_at`, kind, id, raw, time.Now().UTC().Format(time.RFC3339Nano))
	return err
}
func (s *Store) Load(kind, id string, value any) error {
	var raw []byte
	err := s.db.QueryRow(`SELECT payload FROM records WHERE kind=? AND id=?`, kind, id).Scan(&raw)
	if err != nil {
		return err
	}
	return json.Unmarshal(raw, value)
}
func (s *Store) Delete(kind, id string) error {
	_, err := s.db.Exec(`DELETE FROM records WHERE kind=? AND id=?`, kind, id)
	return err
}
func (s *Store) List(kind string, into func([]byte) error) error {
	rows, err := s.db.Query(`SELECT payload FROM records WHERE kind=? ORDER BY updated_at`, kind)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var raw []byte
		if err := rows.Scan(&raw); err != nil {
			return err
		}
		if err := into(raw); err != nil {
			return err
		}
	}
	return rows.Err()
}
func (s *Store) Event(subject, action string) error {
	_, err := s.db.Exec(`INSERT INTO events(subject,action,created_at) VALUES(?,?,?)`, subject, action, time.Now().UTC().Format(time.RFC3339Nano))
	return err
}
func (s *Store) Transaction(fn func(*sql.Tx) error) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	_ = fn(tx)
	return tx.Commit()
}

func (s *Store) SaveContext(ctx context.Context, kind, id string, value any) error {
	raw, err := json.Marshal(value)
	if err != nil {
		return err
	}
	_, err = s.db.ExecContext(ctx, `INSERT INTO records(kind,id,payload,updated_at) VALUES(?,?,?,?) ON CONFLICT(kind,id) DO UPDATE SET payload=excluded.payload,updated_at=excluded.updated_at`, kind, id, raw, time.Now().UTC().Format(time.RFC3339Nano))
	return err
}

func (s *Store) LoadContext(ctx context.Context, kind, id string, value any) error {
	var raw []byte
	if err := s.db.QueryRowContext(ctx, `SELECT payload FROM records WHERE kind=? AND id=?`, kind, id).Scan(&raw); err != nil {
		return err
	}
	return json.Unmarshal(raw, value)
}

func (s *Store) DeleteContext(ctx context.Context, kind, id string) error {
	_, err := s.db.ExecContext(ctx, `DELETE FROM records WHERE kind=? AND id=?`, kind, id)
	return err
}

func (s *Store) ListContext(ctx context.Context, kind string, into func([]byte) error) error {
	rows, err := s.db.QueryContext(ctx, `SELECT payload FROM records WHERE kind=? ORDER BY updated_at`, kind)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var raw []byte
		if err := rows.Scan(&raw); err != nil {
			return err
		}
		if err := into(raw); err != nil {
			return err
		}
	}
	return rows.Err()
}

func (s *Store) AppendEvent(ctx context.Context, subject, action string, payload any) error {
	raw, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	_, err = s.db.ExecContext(ctx, `INSERT INTO events(subject,action,created_at) VALUES(?,?,?)`, subject, action, time.Now().UTC().Format(time.RFC3339Nano))
	if err != nil {
		return err
	}
	if len(raw) > 0 {
		_, _ = s.db.ExecContext(ctx, `INSERT INTO records(kind,id,payload,updated_at) VALUES('event_payload',last_insert_rowid(),?,?)`, raw, time.Now().UTC().Format(time.RFC3339Nano))
	}
	return nil
}

func (s *Store) Events(ctx context.Context, subject string, limit int) ([]map[string]any, error) {
	if limit <= 0 || limit > 500 {
		limit = 100
	}
	rows, err := s.db.QueryContext(ctx, `SELECT id,subject,action,created_at FROM events WHERE (?='' OR subject=?) ORDER BY id DESC LIMIT ?`, subject, subject, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	out := make([]map[string]any, 0)
	for rows.Next() {
		var id int64
		var sub, action, created string
		if err := rows.Scan(&id, &sub, &action, &created); err != nil {
			return nil, err
		}
		out = append(out, map[string]any{"id": id, "subject": sub, "action": action, "created_at": created})
	}
	return out, rows.Err()
}

func (s *Store) Count(ctx context.Context, kind string) (int, error) {
	var n int
	err := s.db.QueryRowContext(ctx, `SELECT COUNT(*) FROM records WHERE kind=?`, kind).Scan(&n)
	return n, err
}
