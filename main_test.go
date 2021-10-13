package bug

import (
	"context"
	"entgo.io/bug/ent/enttest"
	"entgo.io/ent/dialect"
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

func TestBug(t *testing.T) {
	ctx := context.Background()
	client := enttest.Open(t, dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	// Run schema migration.
	if err := client.Schema.Create(ctx); err != nil {
		t.Fatal(err)
	}
	client.User.Create().ExecX(ctx)
	if n := client.User.Query().CountX(ctx); n != 1 {
		t.Errorf("unexpected number of users: %d", n)
	}
}