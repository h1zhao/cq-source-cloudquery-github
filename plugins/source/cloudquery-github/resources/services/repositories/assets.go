package repositories

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/google/go-github/v48/github"
	"github.com/h1zhao/cq-source-cloudquery-github/plugins/source/cloudquery-github/client"
)

func assets() *schema.Table {
	return &schema.Table{
		Name:      "github_release_assets",
		Resolver:  fetchAssets,
		Transform: client.TransformWithStruct(&github.ReleaseAsset{}, transformers.WithPrimaryKeys("ID")),
		Columns:   []schema.Column{client.OrgColumn, client.RepositoryIDColumn},
	}
}

func fetchAssets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	release := parent.Item.(*github.RepositoryRelease)
	repo := parent.Parent.Item.(*github.Repository)
	opts := &github.ListOptions{PerPage: 100}
	for {
		releases, resp, err := c.Github.Repositories.ListReleaseAssets(ctx, c.Org, *repo.Name, *release.ID, opts)
		if err != nil {
			return err
		}
		res <- releases
		opts.Page = resp.NextPage
		if opts.Page == resp.LastPage {
			break
		}
	}
	return nil
}
