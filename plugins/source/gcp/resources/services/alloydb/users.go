package sql

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	alloydb "google.golang.org/api/alloydb/v1"
	pb "google.golang.org/api/alloydb/v1"
)

func Users() *schema.Table {
	return &schema.Table{
		Name:        "gcp_alloydb_users",
		Description: `https://cloud.google.com/alloydb/docs/reference/rest/v1/projects.locations.clusters.users`,
		Resolver:    fetchUsers,
		Multiplex:   client.ProjectMultiplexEnabledServices("alloydb.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.Backup{}, transformers.WithPrimaryKeys("name")),
		Columns: []schema.Column{
			client.ProjectIDColumn(false),
		},
	}
}
func fetchUsers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	alloydbClient, err := alloydb.NewService(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	alloydbService := alloydbClient.Projects.Locations.Clusters
	if err != nil {
		return err
	}
	for {
		output, err := alloydbService.Users.List(c.ProjectId).PageSize(1000).PageToken(nextPageToken).Do()
		if err != nil {
			return err
		}
		res <- output.Users
		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}