// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: items.sql

package db

import (
	"context"
)

const addItem = `-- name: AddItem :exec
insert into items (
  title,
  description,
  collection
) values (
  ?1,
  ?2,
  ?3
)
`

type AddItemParams struct {
	Title       string
	Description string
	Collection  string
}

func (q *Queries) AddItem(ctx context.Context, arg AddItemParams) error {
	_, err := q.db.ExecContext(ctx, addItem, arg.Title, arg.Description, arg.Collection)
	return err
}

const getCollections = `-- name: GetCollections :many
select distinct collection
from items
`

func (q *Queries) GetCollections(ctx context.Context) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, getCollections)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var collection string
		if err := rows.Scan(&collection); err != nil {
			return nil, err
		}
		items = append(items, collection)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getItems = `-- name: GetItems :many
select id, title, description, status, collection
from items
where
  collection = ?1
`

func (q *Queries) GetItems(ctx context.Context, collection string) ([]Item, error) {
	rows, err := q.db.QueryContext(ctx, getItems, collection)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Item
	for rows.Next() {
		var i Item
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.Status,
			&i.Collection,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateItem = `-- name: UpdateItem :exec
update items
set
  title = ?1,
  description = ?2,
  status = ?3,
  collection = ?4
where
  id = ?5
`

type UpdateItemParams struct {
	Title       string
	Description string
	Status      string
	Collection  string
	ID          int64
}

func (q *Queries) UpdateItem(ctx context.Context, arg UpdateItemParams) error {
	_, err := q.db.ExecContext(ctx, updateItem,
		arg.Title,
		arg.Description,
		arg.Status,
		arg.Collection,
		arg.ID,
	)
	return err
}
