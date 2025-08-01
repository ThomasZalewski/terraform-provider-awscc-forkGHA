// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package bcmdataexports

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_bcmdataexports_export", exportDataSource)
}

// exportDataSource returns the Terraform awscc_bcmdataexports_export data source.
// This Terraform data source corresponds to the CloudFormation AWS::BCMDataExports::Export resource.
func exportDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Export
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "DataQuery": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "QueryStatement": {
		//	          "maxLength": 36000,
		//	          "minLength": 1,
		//	          "pattern": "^[\\S\\s]*$",
		//	          "type": "string"
		//	        },
		//	        "TableConfigurations": {
		//	          "additionalProperties": false,
		//	          "patternProperties": {
		//	            "": {
		//	              "additionalProperties": false,
		//	              "patternProperties": {
		//	                "": {
		//	                  "maxLength": 1024,
		//	                  "minLength": 0,
		//	                  "pattern": "^[\\S\\s]*$",
		//	                  "type": "string"
		//	                }
		//	              },
		//	              "type": "object"
		//	            }
		//	          },
		//	          "type": "object"
		//	        }
		//	      },
		//	      "required": [
		//	        "QueryStatement"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "Description": {
		//	      "maxLength": 1024,
		//	      "minLength": 0,
		//	      "pattern": "^[\\S\\s]*$",
		//	      "type": "string"
		//	    },
		//	    "DestinationConfigurations": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "S3Destination": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "S3Bucket": {
		//	              "maxLength": 1024,
		//	              "minLength": 0,
		//	              "pattern": "^[\\S\\s]*$",
		//	              "type": "string"
		//	            },
		//	            "S3OutputConfigurations": {
		//	              "additionalProperties": false,
		//	              "properties": {
		//	                "Compression": {
		//	                  "enum": [
		//	                    "GZIP",
		//	                    "PARQUET"
		//	                  ],
		//	                  "type": "string"
		//	                },
		//	                "Format": {
		//	                  "enum": [
		//	                    "TEXT_OR_CSV",
		//	                    "PARQUET"
		//	                  ],
		//	                  "type": "string"
		//	                },
		//	                "OutputType": {
		//	                  "enum": [
		//	                    "CUSTOM"
		//	                  ],
		//	                  "type": "string"
		//	                },
		//	                "Overwrite": {
		//	                  "enum": [
		//	                    "CREATE_NEW_REPORT",
		//	                    "OVERWRITE_REPORT"
		//	                  ],
		//	                  "type": "string"
		//	                }
		//	              },
		//	              "required": [
		//	                "Compression",
		//	                "Format",
		//	                "OutputType",
		//	                "Overwrite"
		//	              ],
		//	              "type": "object"
		//	            },
		//	            "S3Prefix": {
		//	              "maxLength": 1024,
		//	              "minLength": 0,
		//	              "pattern": "^[\\S\\s]*$",
		//	              "type": "string"
		//	            },
		//	            "S3Region": {
		//	              "maxLength": 1024,
		//	              "minLength": 0,
		//	              "pattern": "^[\\S\\s]*$",
		//	              "type": "string"
		//	            }
		//	          },
		//	          "required": [
		//	            "S3Bucket",
		//	            "S3OutputConfigurations",
		//	            "S3Prefix",
		//	            "S3Region"
		//	          ],
		//	          "type": "object"
		//	        }
		//	      },
		//	      "required": [
		//	        "S3Destination"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "ExportArn": {
		//	      "maxLength": 2048,
		//	      "minLength": 20,
		//	      "pattern": "^arn:aws[-a-z0-9]*:[-a-z0-9]+:[-a-z0-9]*:[0-9]{12}:[-a-zA-Z0-9/:_]+$",
		//	      "type": "string"
		//	    },
		//	    "Name": {
		//	      "maxLength": 128,
		//	      "minLength": 1,
		//	      "pattern": "^[0-9A-Za-z\\-_]+$",
		//	      "type": "string"
		//	    },
		//	    "RefreshCadence": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "Frequency": {
		//	          "enum": [
		//	            "SYNCHRONOUS"
		//	          ],
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "Frequency"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "required": [
		//	    "DataQuery",
		//	    "DestinationConfigurations",
		//	    "Name",
		//	    "RefreshCadence"
		//	  ],
		//	  "type": "object"
		//	}
		"export": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: DataQuery
				"data_query": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: QueryStatement
						"query_statement": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: TableConfigurations
						"table_configurations": // Pattern: ""
						schema.MapAttribute{    /*START ATTRIBUTE*/
							ElementType: types.MapType{ElemType: types.StringType},
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: Description
				"description": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: DestinationConfigurations
				"destination_configurations": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: S3Destination
						"s3_destination": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: S3Bucket
								"s3_bucket": schema.StringAttribute{ /*START ATTRIBUTE*/
									Computed: true,
								}, /*END ATTRIBUTE*/
								// Property: S3OutputConfigurations
								"s3_output_configurations": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: Compression
										"compression": schema.StringAttribute{ /*START ATTRIBUTE*/
											Computed: true,
										}, /*END ATTRIBUTE*/
										// Property: Format
										"format": schema.StringAttribute{ /*START ATTRIBUTE*/
											Computed: true,
										}, /*END ATTRIBUTE*/
										// Property: OutputType
										"output_type": schema.StringAttribute{ /*START ATTRIBUTE*/
											Computed: true,
										}, /*END ATTRIBUTE*/
										// Property: Overwrite
										"overwrite": schema.StringAttribute{ /*START ATTRIBUTE*/
											Computed: true,
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
									Computed: true,
								}, /*END ATTRIBUTE*/
								// Property: S3Prefix
								"s3_prefix": schema.StringAttribute{ /*START ATTRIBUTE*/
									Computed: true,
								}, /*END ATTRIBUTE*/
								// Property: S3Region
								"s3_region": schema.StringAttribute{ /*START ATTRIBUTE*/
									Computed: true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: ExportArn
				"export_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: Name
				"name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: RefreshCadence
				"refresh_cadence": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Frequency
						"frequency": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: ExportArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 2048,
		//	  "minLength": 20,
		//	  "pattern": "^arn:aws[-a-z0-9]*:[-a-z0-9]+:[-a-z0-9]*:[0-9]{12}:[-a-zA-Z0-9/:_]+$",
		//	  "type": "string"
		//	}
		"export_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
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
		//	  "maxItems": 200,
		//	  "minItems": 0,
		//	  "type": "array"
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Computed: true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::BCMDataExports::Export",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::BCMDataExports::Export").WithTerraformTypeName("awscc_bcmdataexports_export")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"compression":                "Compression",
		"data_query":                 "DataQuery",
		"description":                "Description",
		"destination_configurations": "DestinationConfigurations",
		"export":                     "Export",
		"export_arn":                 "ExportArn",
		"format":                     "Format",
		"frequency":                  "Frequency",
		"key":                        "Key",
		"name":                       "Name",
		"output_type":                "OutputType",
		"overwrite":                  "Overwrite",
		"query_statement":            "QueryStatement",
		"refresh_cadence":            "RefreshCadence",
		"s3_bucket":                  "S3Bucket",
		"s3_destination":             "S3Destination",
		"s3_output_configurations":   "S3OutputConfigurations",
		"s3_prefix":                  "S3Prefix",
		"s3_region":                  "S3Region",
		"table_configurations":       "TableConfigurations",
		"tags":                       "Tags",
		"value":                      "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
