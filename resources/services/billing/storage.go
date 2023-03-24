package billing

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/google/go-github/v48/github"
	"github.com/h1zhao/cq-source-cloudquery-github/client"
)

func Storage() *schema.Table {
	return &schema.Table{
		Name:      "github_billing_storage",
		Resolver:  fetchStorage,
		Multiplex: client.OrgMultiplex,
		Transform: client.TransformWithStruct(&github.StorageBilling{}),
		Columns:   []schema.Column{client.OrgColumn},
	}
}

func fetchStorage(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	billing, _, err := c.Github.Billing.GetStorageBillingOrg(ctx, c.Org)
	if err != nil {
		return err
	}
	res <- billing
	return nil
}
