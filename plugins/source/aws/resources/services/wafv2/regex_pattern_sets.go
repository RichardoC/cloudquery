// Code generated by codegen; DO NOT EDIT.

package wafv2

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func RegexPatternSets() *schema.Table {
	return &schema.Table{
		Name:                "aws_wafv2_regex_pattern_sets",
		Description:         "https://docs.aws.amazon.com/waf/latest/APIReference/API_RegexPatternSet.html",
		Resolver:            fetchWafv2RegexPatternSets,
		PreResourceResolver: getRegexPatternSet,
		Multiplex:           client.ServiceAccountRegionScopeMultiplexer("waf-regional"),
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
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveRegexPatternSetTags,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ARN"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Id"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "regular_expression_list",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("RegularExpressionList"),
			},
		},
	}
}
