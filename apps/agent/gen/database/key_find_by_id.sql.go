// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: key_find_by_id.sql

package database

import (
	"context"
)

const findKeyById = `-- name: FindKeyById :one
SELECT
    id, hash, start, owner_id, meta, created_at, expires, ratelimit_type, ratelimit_limit, ratelimit_refill_rate, ratelimit_refill_interval, workspace_id, for_workspace_id, name, remaining_requests, key_auth_id, total_uses
FROM
    ` + "`" + `keys` + "`" + `
WHERE
    id = ?
`

func (q *Queries) FindKeyById(ctx context.Context, id string) (Key, error) {
	row := q.db.QueryRowContext(ctx, findKeyById, id)
	var i Key
	err := row.Scan(
		&i.ID,
		&i.Hash,
		&i.Start,
		&i.OwnerID,
		&i.Meta,
		&i.CreatedAt,
		&i.Expires,
		&i.RatelimitType,
		&i.RatelimitLimit,
		&i.RatelimitRefillRate,
		&i.RatelimitRefillInterval,
		&i.WorkspaceID,
		&i.ForWorkspaceID,
		&i.Name,
		&i.RemainingRequests,
		&i.KeyAuthID,
		&i.TotalUses,
	)
	return i, err
}
