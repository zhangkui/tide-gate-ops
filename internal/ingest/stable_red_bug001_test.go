package ingest
import ("context"; "testing")
func TestBug01_TideGate(t *testing.T) {
	q := NewQueue(1); q.Close()
	for i:=0; i<100; i++ { if _, err := q.Pop(context.Background()); err == nil { t.Fatal("closed queue returned a zero reading instead of an error") } }
}