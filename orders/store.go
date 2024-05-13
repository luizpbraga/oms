package main

import "context"

type store struct {
	// todo: add mongo db dependencie
}

func NewStore() *store {
	return &store{}
}

func (sto *store) Create(ctx context.Context) error {
	return nil
}
