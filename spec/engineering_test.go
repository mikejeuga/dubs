package spec_test

import (
	"context"
	"github.com/adamluzsi/testcase"
	"github.com/mikejeuga/dubs/team"
	"testing"
)

type DummyTeam struct {
}

func (d *DummyTeam) Find(ctx context.Context, engineerID string) (team.Engineer, error) {
	//TODO implement me
	panic("implement me")
}

func TestFind(t *testing.T) {
	s := testcase.NewSpec(t)
	ctx := context.Background()

	aTeam := testcase.Let(s, func(t *testcase.T) *team.Team {
		teamMembers := make([]team.Engineer, 0)
		theTeam := team.Team{Members: teamMembers}
		return &theTeam
	})

	SaltPay := testcase.Let(s, func(t *testcase.T) *team.Company {
		teams := make([]team.Team, 0)
		teams = append(teams, *aTeam.Get(t))
		return team.NewCompany(teams)
	})

	dummyTeam := testcase.Let(s, func(t *testcase.T) *DummyTeam {
		return &DummyTeam{}
	})

	s.Test("When the organisation is looking for an Engineer", func(t *testcase.T) {
		//The dependency I inject here is irrelevant, it is almost noop. Here the dummy team does nothing.
		act := func(t *testcase.T) (team.Engineer, error) {
			return SaltPay.Get(t).FindEngineer(ctx, dummyTeam.Get(t))
		}

		s.Then("It find the engineer", func(t *testcase.T) {
			engineer, err := act(t)
			t.Must.NoError(err)
			t.Must.Equal("", engineer.LastName)
		})
	})

}
