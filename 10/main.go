package main

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
)

// Find retrieves an entity from the database via its identifier.
func (r Repository[E]) Find(ctx context.Context, id string, table string, transaction Transaction[sqlx.Tx, sqlx.DB]) (*E, error) {
	var row *sql.Row
	var entity E
	fields := strings.Join(r.GetFields(entity), ",")
	statement := fmt.Sprintf("SELECT %s FROM %s WHERE id = %s", fields, table, id)
	if transaction != nil {
		row = transaction.GetDriver().QueryRowxContext(ctx, statement)
	} else {
		conn := r.manager.FindByType(sql.ConnectionTypeRead)
		row = conn.GetDriver().QueryRowxContext(ctx, statement)
	}
	if err := row.StructScan(&entity); err != nil {
		if err == stdsql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}
	return &entity, nil
}
