package organizations

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/google/go-github/v48/github"
	"github.com/h1zhao/cq-source-cloudquery-github/plugins/source/cloudquery-github/client"
)

func alerts() *schema.Table {
	return &schema.Table{
		Name:      "github_organization_dependabot_alerts",
		Resolver:  fetchAlerts,
		Transform: client.TransformWithStruct(&github.DependabotAlert{}, transformers.WithPrimaryKeys("HTMLURL")),
		Columns:   []schema.Column{client.OrgColumn},
	}
}

func fetchAlerts(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	opts := &github.ListAlertsOptions{ListCursorOptions: github.ListCursorOptions{PerPage: 99}}

	for {
		alerts, resp, err := c.Github.Dependabot.ListOrgAlerts(ctx, c.Org, opts)
		if err != nil {
			return err
		}
		res <- alerts
		opts.After = resp.After
		if resp.After == "" {
			break
		}
	}

	return nil
}
