package testing

import (
	"testing"

	db "github.com/luitel777/purkheli/internal/db"
)

func TestDb(t *testing.T) {
	t.Fatalf("ok")
	d := db.Database{}
	d.InitiateDB()
	d.StoreDB("hello", "textarea", "path", "123")
}
