package controllers

import (
	"context"

	"github.com/somatom98/todoist/domain"
)

type ItemsRepo interface {
	Get(ctx context.Context, collection domain.Collection, status domain.Status) ([]domain.Item, error)
	Collections(ctx context.Context) ([]domain.Collection, error)
	Add(ctx context.Context, item domain.Item) error
	Update(ctx context.Context, id int64, item domain.Item) error
}

type PaneSelector interface {
	CurrentFocus() domain.Pane
	FocusNext()
	FocusPrev()
	SetFocus(view domain.Pane)
}
