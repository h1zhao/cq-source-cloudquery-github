package repositories

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/google/go-github/v48/github"
	"github.com/h1zhao/cq-source-cloudquery-github/client"
)

func Repositories() *schema.Table {
	return &schema.Table{
		Name:      "github_repositories",
		Resolver:  fetchRepositories,
		Multiplex: client.OrgRepositoryMultiplex,
		Transform: client.TransformWithStruct(&github.Repository{}, transformers.WithPrimaryKeys("ID")),
		Columns:   []schema.Column{client.OrgColumn},
		Relations: []*schema.Table{alerts(), releases(), secrets()},
	}
}

func fetchRepositories(_ context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	// Repositories are synced during init and multiplexed
	res <- c.Repository
	return nil
}
