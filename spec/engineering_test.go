package spec_test

import (
	"context"
	"github.com/adamluzsi/testcase"
	"github.com/mikejeuga/dubs/team"
	"math/rand"
	"testing"
)

type DummyTeam struct {
}

func (d *DummyTeam) Find(ctx context.Context, engineerID int) (team.Engineer, error) {
	return team.Engineer{}, nil
}

type StubTeam struct {
	Members team.Squad
}

func (s *StubTeam) Find(ctx context.Context, engineerID int) (team.Engineer, error) {
	return s.Members[engineerID], nil
}

func TestFind(t *testing.T) {
	s := testcase.NewSpec(t)
	ctx := context.Background()

	anEngineer := testcase.Let(s, func(t *testcase.T) *team.Engineer {
		engineer := &team.Engineer{
			ID:        t.Random.Int(),
			FirstName: "Engineer",
			LastName:  "Salter",
		}

		return engineer
	})

	aTeam := testcase.Let(s, func(t *testcase.T) *team.Team {
		teamMembers := make(map[int]team.Engineer)

		theTeam := team.Team{Members: teamMembers}
		return &theTeam
	})

	SaltPay := testcase.Let(s, func(t *testcase.T) *team.Company {
		teams := make([]team.Team, 0)
		squadron := *aTeam.Get(t)
		teams = append(teams, squadron)
		return team.NewCompany(teams)
	})

	dummyTeam := testcase.Let(s, func(t *testcase.T) *DummyTeam {
		return &DummyTeam{}
	})

	StubbyTeam := testcase.Let(s, func(t *testcase.T) *StubTeam {
		squad := make(team.Squad, 0)
		squad[anEngineer.Get(t).ID] = *anEngineer.Get(t)
		return &StubTeam{Members: squad}
	})

	s.Describe("Dummy", func(s *testcase.Spec) {
		s.Test("When the organisation is looking for an Engineer within a Dummy Team", func(t *testcase.T) {
			//The dependency I inject here is irrelevant, it is almost noop. Here the dummy team does nothing.
			act := func(t *testcase.T) (team.Engineer, error) {
				return SaltPay.Get(t).FindEngineer(ctx, dummyTeam.Get(t), rand.Int())
			}

			s.Then("It find the engineer", func(t *testcase.T) {
				_, err := act(t)
				t.Must.NoError(err)
			})
		})
	})

	s.Describe("Stub", func(s *testcase.Spec) {
		s.Test("When the organisation is looking for an Engineer within a Stubby Team", func(t *testcase.T) {
			act2 := func(t *testcase.T) (team.Engineer, error) {
				return SaltPay.Get(t).FindEngineer(ctx, StubbyTeam.Get(t), anEngineer.Get(t).ID)
			}

			s.Then("It find the engineer", func(t *testcase.T) {
				engineer, err := act2(t)
				t.Must.NoError(err)
				t.Must.Equal("Salter", engineer.LastName)
			})
		})
	})

}
