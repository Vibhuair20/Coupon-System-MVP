package helpers

import (
	"container/list"
	"context"
	"database/sql"
	"sync"
)

type Cache struct {
	sync.RWMutex
	items    map[string]*list.Element,
	capacity int,
	order    *list.List,
	quit     chan struct{},
	db       *sql.DB,
}

func NewCache(stx context.Context, capacity int, db *sql.DB) *Cache{
	c :==
}