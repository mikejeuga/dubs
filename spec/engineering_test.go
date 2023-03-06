package spec_test

import (
	"context"
	"github.com/adamluzsi/testcase"
	"github.com/mikejeuga/dubs/team"
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
	for _, engineer := range s.Members {
		if engineer.ID == engineerID {
			return engineer, nil
		}
	}
	return team.Engineer{}, nil
}

type SpyTeam struct {
	Members    team.Squad
	findCalled bool
	findCalls  int
}

func (sp *SpyTeam) Find(ctx context.Context, engineerID int) (team.Engineer, error) {
	for _, engineer := range sp.Members {
		if engineer.ID == engineerID {
			sp.findCalledTrue()
			sp.Called()
			return engineer, nil
		}
	}
	return team.Engineer{}, nil
}

func (sp *SpyTeam) findCalledTrue() bool {
	sp.findCalled = true
	return sp.findCalled
}

func (sp *SpyTeam) Called() int {
	sp.findCalls += 1
	return sp.findCalls
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
		teamMembers := make(team.Squad, 0)
		theTeam := team.Team{Members: teamMembers}
		return &theTeam
	})

	SaltPay := testcase.Let(s, func(t *testcase.T) *team.Company {
		teams := make([]team.Team, 0)
		squadron := *aTeam.Get(t)
		squadron.Members = append(squadron.Members, *anEngineer.Get(t))
		teams = append(teams, squadron)
		return team.NewCompany(teams)
	})

	dummyTeam := testcase.Let(s, func(t *testcase.T) *DummyTeam {
		return &DummyTeam{}
	})

	StubbyTeam := testcase.Let(s, func(t *testcase.T) *StubTeam {
		squad := make(team.Squad, 0)
		squad = append(squad, *anEngineer.Get(t))
		return &StubTeam{Members: squad}
	})

	theSpyTeam := testcase.Let(s, func(t *testcase.T) *SpyTeam {
		squad := make(team.Squad, 0)
		squad = append(squad, *anEngineer.Get(t))
		return &SpyTeam{
			Members: squad,
		}
	})

	s.Describe("Dummy", func(s *testcase.Spec) {
		act := func(t *testcase.T) (team.Engineer, error) {
			return SaltPay.Get(t).FindEngineer(ctx, dummyTeam.Get(t))
		}
		s.Test("When the organisation is looking for an Engineer within a Dummy Team", func(t *testcase.T) {
			//The dependency I inject here is irrelevant, it is almost noop. Here the dummy team does nothing.
			_, err := act(t)
			t.Must.NoError(err)
		})
	})

	s.Describe("Stub", func(s *testcase.Spec) {
		act2 := func(t *testcase.T) (team.Engineer, error) {
			return SaltPay.Get(t).FindEngineer(ctx, StubbyTeam.Get(t))
		}
		s.Test("When the organisation is looking for an Engineer within a Stubby Team", func(t *testcase.T) {
			engineer, err := act2(t)
			t.Must.NoError(err)
			t.Must.NotNil(engineer)
			t.Must.Equal("Salter", engineer.LastName)
		})
	})

	s.Describe("Spy", func(s *testcase.Spec) {
		act2 := func(t *testcase.T) (team.Engineer, error) {
			return SaltPay.Get(t).FindEngineer(ctx, theSpyTeam.Get(t))
		}
		s.Test("When the organisation is looking for an Engineer within a Spy Team", func(t *testcase.T) {
			engineer, err := act2(t)
			t.Must.NoError(err)
			t.Must.NotNil(engineer)
			t.Must.True(theSpyTeam.Get(t).findCalledTrue())
			t.Must.Equal(1, theSpyTeam.Get(t).findCalls)
		})
	})

}
