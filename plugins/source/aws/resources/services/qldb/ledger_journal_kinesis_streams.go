// Code generated by codegen; DO NOT EDIT.

package qldb

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func LedgerJournalKinesisStreams() *schema.Table {
	return &schema.Table{
		Name:        "aws_qldb_ledger_journal_kinesis_streams",
		Description: "https://docs.aws.amazon.com/qldb/latest/developerguide/API_JournalKinesisStreamDescription.html",
		Resolver:    fetchQldbLedgerJournalKinesisStreams,
		Multiplex:   client.ServiceAccountRegionMultiplexer("qldb"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "ledger_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "kinesis_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("KinesisConfiguration"),
			},
			{
				Name:     "ledger_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LedgerName"),
			},
			{
				Name:     "role_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RoleArn"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "stream_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StreamId"),
			},
			{
				Name:     "stream_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StreamName"),
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
			},
			{
				Name:     "creation_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreationTime"),
			},
			{
				Name:     "error_cause",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ErrorCause"),
			},
			{
				Name:     "exclusive_end_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("ExclusiveEndTime"),
			},
			{
				Name:     "inclusive_start_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("InclusiveStartTime"),
			},
		},
	}
}
