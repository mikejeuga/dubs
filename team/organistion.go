package team

import (
	"context"
)

type Engineer struct {
	ID                  int
	FirstName, LastName string
}

type Company struct {
	Engineering []Team
}

func NewCompany(engineering []Team) *Company {
	return &Company{Engineering: engineering}
}

type Squad map[int]Engineer

func (c *Company) FindEngineer(ctx context.Context, team EngineerFinder, id int) (Engineer, error) {
	engineer, err := team.Find(ctx, id)
	if err != nil {
		return Engineer{}, err
	}
	return engineer, nil
}
