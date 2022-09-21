package team

import (
	"context"
)

type Engineer struct {
	ID                  string
	FirstName, LastName string
}

type Company struct {
	Engineering []Team
}

func NewCompany(engineering []Team) *Company {
	return &Company{Engineering: engineering}
}

type Squad []Engineer

func (c *Company) FindEngineer(ctx context.Context, team EngineerFinder) (Engineer, error) {
	return Engineer{}, nil
}
