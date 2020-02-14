package storage

import (
	"database/sql"
)

type Base struct {
	ID         string
	InsertedAt sql.NullString
	UpdatedAt  sql.NullString
}

func base(tableName string) string {
	return `
	CREATE INDEX IF NOT EXISTS idx_` + tableName + `_created ON ` + tableName + `(created_at);
	CREATE INDEX IF NOT EXISTS idx_` + tableName + `_update  ON ` + tableName + `(updated_at);

	CREATE TRIGGER IF NOT EXISTS trg_` + tableName + `_created
		AFTER INSERT ON ` + tableName + `
	BEGIN
		UPDATE ` + tableName + ` SET created_at = DATETIME('now', 'localtime') WHERE id=new.id;
	END;

	CREATE TRIGGER IF NOT EXISTS trg_` + tableName + `_updated
		AFTER INSERT ON ` + tableName + `
	BEGIN
		UPDATE ` + tableName + ` SET updated_at = DATETIME('now', 'localtime') WHERE id=new.id;
	END;`
}
