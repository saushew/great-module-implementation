package main

import (
	"context"
	"fmt"
	"time"

	"github.com/saushew/great-app/core"
)

var ExecutorInstance Executor

type Executor struct {
	name string
}

func (m *Executor) Initialize(config *core.ModuleConfig) error {
	m.name = config.Name
	return nil
}

func (m *Executor) Execute(ctx context.Context) error {
	t := time.NewTicker(time.Second)
	defer t.Stop()
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-t.C:
			fmt.Println(m.name)
		}
	}
}

func (m *Executor) Stop() error {
	return nil
}
