package after

import (
	"context"
	"fmt"

	"github.com/google/go-github/v47/github"
)

type githubOrganizations interface {
	List(ctx context.Context, user string, opts *github.ListOptions) ([]*github.Organization, *github.Response, error)
	IsMember(ctx context.Context, org, user string) (bool, *github.Response, error)
}

type obj struct {
	user string
	gh   githubOrganizations
}

func New(user string) obj {
	return obj{
		user: user,
		gh:   github.NewClient(nil).Organizations,
	}
}

func (m *obj) CountOrganizations(ctx context.Context) (int, error) {
	orgs, _, err := m.gh.List(ctx, m.user, nil)
	if err != nil {
		return 0, fmt.Errorf("couldn't list organizations: %w", err)
	}
	return len(orgs), nil
}

func (m *obj) IsMemberOf(ctx context.Context, org string) (bool, error) {
	result, _, err := m.gh.IsMember(ctx, org, m.user)
	if err != nil {
		return false, fmt.Errorf("couldn't call IsMember: %w", err)
	}
	return result, nil
}
