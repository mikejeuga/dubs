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

type Squad []Engineer

func (c *Company) FindEngineer(ctx context.Context, team EngineerFinder) (Engineer, error) {
	//engineer, err := team.Find(ctx, id)
	//if err != nil {
	//	return Engineer{}, err
	//}
	//return engineer, nil
	var id int

	for _, engTeam := range c.Engineering {
		if len(engTeam.Members) > 0 {
			id = engTeam.Members[0].ID
		}
		break
	}

	engineer, err := team.Find(ctx, id)
	if err != nil {
		return Engineer{}, err
	}

	return engineer, err
}
