package team

import (
	"context"
)

type Team struct {
	Members Squad
}

func (t Team) Find(ctx context.Context, engineerID string) (Engineer, error) {
	//TODO implement me
	panic("implement me")
}

func NewARS(members Squad) *Team {
	return &Team{Members: members}
}

type EngineerFinder interface {
	Find(ctx context.Context, engineerID string) (Engineer, error)
}
