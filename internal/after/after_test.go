package after

import (
	"context"
	"testing"

	"github.com/google/go-github/v47/github"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type mockGHList struct {
	*github.OrganizationsService
}

func (mockGHList) List(context.Context, string, *github.ListOptions) ([]*github.Organization, *github.Response, error) {
	return []*github.Organization{{}, {}, {}}, nil, nil
}

func TestCountOrganizations(t *testing.T) {
	obj := obj{
		gh: mockGHList{},
	}

	c, err := obj.CountOrganizations(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if c != 3 {
		t.Fatalf("unexpected number of organizations: %d", c)
	}
}

type mockGHIsMember struct {
	mock.Mock
	*github.OrganizationsService
}

func (m *mockGHIsMember) IsMember(ctx context.Context, org, user string) (bool, *github.Response, error) {
	args := m.Called(ctx, org, user)
	return args.Bool(0), args.Get(1).(*github.Response), args.Error(2)
}

func TestIsMember(t *testing.T) {
	gh := mockGHIsMember{}
	obj := obj{
		user: "apexskier",
		gh:   &gh,
	}

	gh.On("IsMember", mock.MatchedBy(func(context.Context) bool { return true }), "org", obj.user).Return(true, (*github.Response)(nil), nil)

	result, err := obj.IsMemberOf(context.Background(), "org")
	require.NoError(t, err)
	assert.True(t, result)
}
