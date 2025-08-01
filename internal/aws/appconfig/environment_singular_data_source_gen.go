// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package appconfig

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_appconfig_environment", environmentDataSource)
}

// environmentDataSource returns the Terraform awscc_appconfig_environment data source.
// This Terraform data source corresponds to the CloudFormation AWS::AppConfig::Environment resource.
func environmentDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ApplicationId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The application ID.",
		//	  "pattern": "[a-z0-9]{4,7}",
		//	  "type": "string"
		//	}
		"application_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The application ID.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DeletionProtectionCheck
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "On resource deletion this controls whether the Deletion Protection check should be applied, bypassed, or (the default) whether the behavior should be controlled by the account-level Deletion Protection setting. See https://docs.aws.amazon.com/appconfig/latest/userguide/deletion-protection.html",
		//	  "enum": [
		//	    "ACCOUNT_DEFAULT",
		//	    "APPLY",
		//	    "BYPASS"
		//	  ],
		//	  "type": "string"
		//	}
		"deletion_protection_check": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "On resource deletion this controls whether the Deletion Protection check should be applied, bypassed, or (the default) whether the behavior should be controlled by the account-level Deletion Protection setting. See https://docs.aws.amazon.com/appconfig/latest/userguide/deletion-protection.html",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A description of the environment.",
		//	  "maxLength": 1024,
		//	  "minLength": 0,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A description of the environment.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: EnvironmentId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The environment ID.",
		//	  "pattern": "[a-z0-9]{4,7}",
		//	  "type": "string"
		//	}
		"environment_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The environment ID.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Monitors
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Amazon CloudWatch alarms to monitor during the deployment process.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "Amazon CloudWatch alarm to monitor during the deployment process.",
		//	    "properties": {
		//	      "AlarmArn": {
		//	        "description": "Amazon Resource Name (ARN) of the Amazon CloudWatch alarm.",
		//	        "maxLength": 2048,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "AlarmRoleArn": {
		//	        "description": "ARN of an AWS Identity and Access Management (IAM) role for AWS AppConfig to monitor AlarmArn.",
		//	        "maxLength": 2048,
		//	        "minLength": 20,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "AlarmArn"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 5,
		//	  "minItems": 0,
		//	  "type": "array"
		//	}
		"monitors": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: AlarmArn
					"alarm_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Amazon Resource Name (ARN) of the Amazon CloudWatch alarm.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: AlarmRoleArn
					"alarm_role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "ARN of an AWS Identity and Access Management (IAM) role for AWS AppConfig to monitor AlarmArn.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "Amazon CloudWatch alarms to monitor during the deployment process.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A name for the environment.",
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A name for the environment.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Metadata to assign to the environment. Tags help organize and categorize your AWS AppConfig resources. Each tag consists of a key and an optional value, both of which you define.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "Metadata to assign to the environment. Tags help organize and categorize your AWS AppConfig resources. Each tag consists of a key and an optional value, both of which you define.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key-value string map. The valid character set is [a-zA-Z1-9+-=._:/]. The tag key can be up to 128 characters and must not start with aws:.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The tag value can be up to 256 characters.",
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key-value string map. The valid character set is [a-zA-Z1-9+-=._:/]. The tag key can be up to 128 characters and must not start with aws:.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The tag value can be up to 256 characters.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "Metadata to assign to the environment. Tags help organize and categorize your AWS AppConfig resources. Each tag consists of a key and an optional value, both of which you define.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::AppConfig::Environment",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::AppConfig::Environment").WithTerraformTypeName("awscc_appconfig_environment")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"alarm_arn":                 "AlarmArn",
		"alarm_role_arn":            "AlarmRoleArn",
		"application_id":            "ApplicationId",
		"deletion_protection_check": "DeletionProtectionCheck",
		"description":               "Description",
		"environment_id":            "EnvironmentId",
		"key":                       "Key",
		"monitors":                  "Monitors",
		"name":                      "Name",
		"tags":                      "Tags",
		"value":                     "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
