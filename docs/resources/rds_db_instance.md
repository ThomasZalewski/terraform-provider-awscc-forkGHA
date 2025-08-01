---
page_title: "awscc_rds_db_instance Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  The AWS::RDS::DBInstance resource creates an Amazon DB instance. The new DB instance can be an RDS DB instance, or it can be a DB instance in an Aurora DB cluster.
  For more information about creating an RDS DB instance, see Creating an Amazon RDS DB instance https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_CreateDBInstance.html in the Amazon RDS User Guide.
  For more information about creating a DB instance in an Aurora DB cluster, see Creating an Amazon Aurora DB cluster https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Aurora.CreateInstance.html in the Amazon Aurora User Guide.
  If you import an existing DB instance, and the template configuration doesn't match the actual configuration of the DB instance, AWS CloudFormation applies the changes in the template during the import operation.
  If a DB instance is deleted or replaced during an update, AWS CloudFormation deletes all automated snapshots. However, it retains manual DB snapshots. During an update that requires replacement, you can apply a stack policy to prevent DB instances from being replaced. For more information, see Prevent Updates to Stack Resources https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/protect-stack-resources.html.
  Updating DB instances
  When properties labeled "Update requires:Replacement https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement" are updated, AWS CloudFormation first creates a replacement DB instance, then changes references from other dependent resources to point to the replacement DB instance, and finally deletes the old DB instance.
  We highly recommend that you take a snapshot of the database before updating the stack. If you don't, you lose the data when AWS CloudFormation replaces your DB instance. To preserve your data, perform the following procedure:
  Deactivate any applications that are using the DB instance so that there's no activity on the DB instance.Create a snapshot of the DB instance. For more information, see Creating a DB Snapshot https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_CreateSnapshot.html.If you want to restore your instance using a DB snapshot, modify the updated template with your DB instance changes and add the DBSnapshotIdentifier property with the ID of the DB snapshot that you want to use.
  After you restore a DB instance with a DBSnapshotIdentifier property, you can delete the DBSnapshotIdentifier property. When you specify this property for an update, the DB instance is not restored from the DB snapshot again, and the data in the database is not changed. However, if you don't specify the DBSnapshotIdentifier property, an empty DB instance is created, and the original DB instance is deleted. If you specify a property that is different from the previous snapshot restore property, a new DB instance is restored from the specified DBSnapshotIdentifier property, and the original DB instance is deleted.Update the stack.
  For more information about updating other properties of this resource, see ModifyDBInstance. For more information about updating stacks, see CloudFormation Stacks Updates https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks.html.
  Deleting DB instances
  For DB instances that are part of an Aurora DB cluster, you can set a deletion policy for your DB instance to control how AWS CloudFormation handles the DB instance when the stack is deleted. For Amazon RDS DB instances, you can choose to retain the DB instance, to delete the DB instance, or to create a snapshot of the DB instance. The default AWS CloudFormation behavior depends on the DBClusterIdentifier property:
  For AWS::RDS::DBInstance resources that don't specify the DBClusterIdentifier property, AWS CloudFormation saves a snapshot of the DB instance.For AWS::RDS::DBInstance resources that do specify the DBClusterIdentifier property, AWS CloudFormation deletes the DB instance.
  For more information, see DeletionPolicy Attribute https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html.
---

# awscc_rds_db_instance (Resource)

The ``AWS::RDS::DBInstance`` resource creates an Amazon DB instance. The new DB instance can be an RDS DB instance, or it can be a DB instance in an Aurora DB cluster.
 For more information about creating an RDS DB instance, see [Creating an Amazon RDS DB instance](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_CreateDBInstance.html) in the *Amazon RDS User Guide*.
 For more information about creating a DB instance in an Aurora DB cluster, see [Creating an Amazon Aurora DB cluster](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Aurora.CreateInstance.html) in the *Amazon Aurora User Guide*.
 If you import an existing DB instance, and the template configuration doesn't match the actual configuration of the DB instance, AWS CloudFormation applies the changes in the template during the import operation.
  If a DB instance is deleted or replaced during an update, AWS CloudFormation deletes all automated snapshots. However, it retains manual DB snapshots. During an update that requires replacement, you can apply a stack policy to prevent DB instances from being replaced. For more information, see [Prevent Updates to Stack Resources](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/protect-stack-resources.html).
   *Updating DB instances* 
 When properties labeled "*Update requires:*[Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)" are updated, AWS CloudFormation first creates a replacement DB instance, then changes references from other dependent resources to point to the replacement DB instance, and finally deletes the old DB instance.
  We highly recommend that you take a snapshot of the database before updating the stack. If you don't, you lose the data when AWS CloudFormation replaces your DB instance. To preserve your data, perform the following procedure:
  1.  Deactivate any applications that are using the DB instance so that there's no activity on the DB instance.
  1.  Create a snapshot of the DB instance. For more information, see [Creating a DB Snapshot](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_CreateSnapshot.html).
  1.  If you want to restore your instance using a DB snapshot, modify the updated template with your DB instance changes and add the ``DBSnapshotIdentifier`` property with the ID of the DB snapshot that you want to use.
 After you restore a DB instance with a ``DBSnapshotIdentifier`` property, you can delete the ``DBSnapshotIdentifier`` property. When you specify this property for an update, the DB instance is not restored from the DB snapshot again, and the data in the database is not changed. However, if you don't specify the ``DBSnapshotIdentifier`` property, an empty DB instance is created, and the original DB instance is deleted. If you specify a property that is different from the previous snapshot restore property, a new DB instance is restored from the specified ``DBSnapshotIdentifier`` property, and the original DB instance is deleted.
  1.  Update the stack.
  
  For more information about updating other properties of this resource, see ``ModifyDBInstance``. For more information about updating stacks, see [CloudFormation Stacks Updates](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks.html).
  *Deleting DB instances* 
 For DB instances that are part of an Aurora DB cluster, you can set a deletion policy for your DB instance to control how AWS CloudFormation handles the DB instance when the stack is deleted. For Amazon RDS DB instances, you can choose to *retain* the DB instance, to *delete* the DB instance, or to *create a snapshot* of the DB instance. The default AWS CloudFormation behavior depends on the ``DBClusterIdentifier`` property:
  1.  For ``AWS::RDS::DBInstance`` resources that don't specify the ``DBClusterIdentifier`` property, AWS CloudFormation saves a snapshot of the DB instance.
  1.   For ``AWS::RDS::DBInstance`` resources that do specify the ``DBClusterIdentifier`` property, AWS CloudFormation deletes the DB instance.
  
  For more information, see [DeletionPolicy Attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html).

## Example Usage

### Basic example
To create a AWS RDS DB instance with basic details
```terraform
resource "awscc_rds_db_instance" "this" {
  allocated_storage       = 10
  db_name                 = "mydb"
  engine                  = "mysql"
  engine_version          = "5.7"
  db_instance_class       = "db.t3.micro"
  master_username         = "foo"
  master_user_password    = "foobarbaz"
  db_parameter_group_name = "default.mysql5.7"
}
```

### Storage Autoscaling example
To enable Storage Autoscaling with instances that support the feature, define the max_allocated_storage argument higher than the allocated_storage argument. Terraform will automatically hide differences with the allocated_storage argument value if autoscaling occurs.
```terraform
resource "awscc_rds_db_instance" "this" {
  # ... other configuration ...
  db_name                 = "mydb"
  engine                  = "mysql"
  engine_version          = "5.7"
  db_instance_class       = "db.t3.micro"
  master_username         = "foo"
  master_user_password    = "foobarbaz"
  db_parameter_group_name = "default.mysql5.7"
  allocated_storage       = 50
  max_allocated_storage   = 100
}
```

### Managed Master Passwords via Secrets Manager, default KMS Key example
You can specify the manage_master_user_password attribute to enable managing the master password with Secrets Manager. You can also update an existing cluster to use Secrets Manager by specify the manage_master_user_password attribute and removing the password attribute (removal is required).
```terraform
resource "awscc_rds_db_instance" "this" {
  allocated_storage           = 10
  db_name                     = "mydb"
  engine                      = "mysql"
  engine_version              = "5.7"
  db_instance_class           = "db.t3.micro"
  manage_master_user_password = true
  master_username             = "foo"
  db_parameter_group_name     = "default.mysql5.7"
}
```

### Managed Master Passwords via Secrets Manager, specific KMS Key example
You can specify the kms_key_id attribute under nested block master_user_secret to specify a specific KMS Key.
```terraform
resource "aws_kms_key" "this" {
  description = "Example KMS Key"
}

resource "awscc_rds_db_instance" "this" {
  allocated_storage           = 10
  db_name                     = "mydb"
  engine                      = "mysql"
  engine_version              = "5.7"
  db_instance_class           = "db.t3.micro"
  manage_master_user_password = true
  master_username             = "foo"
  master_user_secret = {
    kms_key_id = aws_kms_key.this.key_id
  }
  db_parameter_group_name = "default.mysql5.7"
}
```

### DB Instance creation with custom subnet group example
You can create RDS DB instance by using custom db subnet group
```terraform
resource "awscc_rds_db_subnet_group" "this" {
  db_subnet_group_name        = "example"
  db_subnet_group_description = "example subnet group"
  subnet_ids                  = ["subnet-006182af0254ccbc4", "subnet-0c40688dd8ca51435"]
}

resource "awscc_rds_db_instance" "this" {
  allocated_storage       = 10
  db_name                 = "mydb"
  engine                  = "mysql"
  engine_version          = "5.7"
  db_instance_class       = "db.t3.micro"
  master_username         = "foo"
  master_user_password    = "foobarbaz"
  db_parameter_group_name = "default.mysql5.7"
  db_subnet_group_name    = awscc_rds_db_subnet_group.this.id
  tags = [{
    key   = "Name"
    value = "this"
  }]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `allocated_storage` (String) The amount of storage in gibibytes (GiB) to be initially allocated for the database instance.
  If any value is set in the ``Iops`` parameter, ``AllocatedStorage`` must be at least 100 GiB, which corresponds to the minimum Iops value of 1,000. If you increase the ``Iops`` value (in 1,000 IOPS increments), then you must also increase the ``AllocatedStorage`` value (in 100-GiB increments). 
   *Amazon Aurora* 
 Not applicable. Aurora cluster volumes automatically grow as the amount of data in your database increases, though you are only charged for the space that you use in an Aurora cluster volume.
  *Db2* 
 Constraints to the amount of storage for each storage type are the following:
  +  General Purpose (SSD) storage (gp3): Must be an integer from 20 to 64000.
  +  Provisioned IOPS storage (io1): Must be an integer from 100 to 64000.
  
  *MySQL* 
 Constraints to the amount of storage for each storage type are the following: 
  +  General Purpose (SSD) storage (gp2): Must be an integer from 20 to 65536.
  +  Provisioned IOPS storage (io1): Must be an integer from 100 to 65536.
  +  Magnetic storage (standard): Must be an integer from 5 to 3072.
  
  *MariaDB* 
 Constraints to the amount of storage for each storage type are the following: 
  +  General Purpose (SSD) storage (gp2): Must be an integer from 20 to 65536.
  +  Provisioned IOPS storage (io1): Must be an integer from 100 to 65536.
  +  Magnetic storage (standard): Must be an integer from 5 to 3072.
  
  *PostgreSQL* 
 Constraints to the amount of storage for each storage type are the following: 
  +  General Purpose (SSD) storage (gp2): Must be an integer from 20 to 65536.
  +  Provisioned IOPS storage (io1): Must be an integer from 100 to 65536.
  +  Magnetic storage (standard): Must be an integer from 5 to 3072.
  
  *Oracle* 
 Constraints to the amount of storage for each storage type are the following: 
  +  General Purpose (SSD) storage (gp2): Must be an integer from 20 to 65536.
  +  Provisioned IOPS storage (io1): Must be an integer from 100 to 65536.
  +  Magnetic storage (standard): Must be an integer from 10 to 3072.
  
  *SQL Server* 
 Constraints to the amount of storage for each storage type are the following: 
  +  General Purpose (SSD) storage (gp2):
  +  Enterprise and Standard editions: Must be an integer from 20 to 16384.
  +  Web and Express editions: Must be an integer from 20 to 16384.
  
  +  Provisioned IOPS storage (io1):
  +  Enterprise and Standard editions: Must be an integer from 20 to 16384.
  +  Web and Express editions: Must be an integer from 20 to 16384.
  
  +  Magnetic storage (standard):
  +  Enterprise and Standard editions: Must be an integer from 20 to 1024.
  +  Web and Express editions: Must be an integer from 20 to 1024.
- `allow_major_version_upgrade` (Boolean) A value that indicates whether major version upgrades are allowed. Changing this parameter doesn't result in an outage and the change is asynchronously applied as soon as possible.
 Constraints: Major version upgrades must be allowed when specifying a value for the ``EngineVersion`` parameter that is a different major version than the DB instance's current version.
- `apply_immediately` (Boolean) Specifies whether changes to the DB instance and any pending modifications are applied immediately, regardless of the ``PreferredMaintenanceWindow`` setting. If set to ``false``, changes are applied during the next maintenance window. Until RDS applies the changes, the DB instance remains in a drift state. As a result, the configuration doesn't fully reflect the requested modifications and temporarily diverges from the intended state.
 In addition to the settings described in [Modifying a DB instance](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Overview.DBInstance.Modifying.html), this property also determines whether the DB instance reboots when a static parameter is modified in the associated DB parameter group.
 Default: ``true``
- `associated_roles` (Attributes List) The IAMlong (IAM) roles associated with the DB instance. 
  *Amazon Aurora* 
 Not applicable. The associated roles are managed by the DB cluster. (see [below for nested schema](#nestedatt--associated_roles))
- `auto_minor_version_upgrade` (Boolean) A value that indicates whether minor engine upgrades are applied automatically to the DB instance during the maintenance window. By default, minor engine upgrades are applied automatically.
- `automatic_backup_replication_kms_key_id` (String) The AWS KMS key identifier for encryption of the replicated automated backups. The KMS key ID is the Amazon Resource Name (ARN) for the KMS encryption key in the destination AWS-Region, for example, ``arn:aws:kms:us-east-1:123456789012:key/AKIAIOSFODNN7EXAMPLE``.
- `automatic_backup_replication_region` (String) The AWS-Region associated with the automated backup.
- `automatic_backup_replication_retention_period` (Number) The retention period for automated backups in a different AWS Region. Use this parameter to set a unique retention period that only applies to cross-Region automated backups. To enable automated backups in a different Region, specify a positive value for the ``AutomaticBackupReplicationRegion`` parameter. 
 If not specified, this parameter defaults to the value of the ``BackupRetentionPeriod`` parameter. The maximum allowed value is 35.
- `availability_zone` (String) The Availability Zone (AZ) where the database will be created. For information on AWS-Regions and Availability Zones, see [Regions and Availability Zones](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.RegionsAndAvailabilityZones.html).
 For Amazon Aurora, each Aurora DB cluster hosts copies of its storage in three separate Availability Zones. Specify one of these Availability Zones. Aurora automatically chooses an appropriate Availability Zone if you don't specify one.
 Default: A random, system-chosen Availability Zone in the endpoint's AWS-Region.
 Constraints:
  +  The ``AvailabilityZone`` parameter can't be specified if the DB instance is a Multi-AZ deployment.
  +  The specified Availability Zone must be in the same AWS-Region as the current endpoint.
  
 Example: ``us-east-1d``
- `backup_retention_period` (Number) The number of days for which automated backups are retained. Setting this parameter to a positive number enables backups. Setting this parameter to 0 disables automated backups.
  *Amazon Aurora* 
 Not applicable. The retention period for automated backups is managed by the DB cluster.
 Default: 1
 Constraints:
  +  Must be a value from 0 to 35
  +  Can't be set to 0 if the DB instance is a source to read replicas
- `backup_target` (String) The location for storing automated backups and manual snapshots.
 Valid Values:
  +  ``local`` (Dedicated Local Zone)
  +  ``outposts`` (AWS Outposts)
  +  ``region`` (AWS-Region)
  
 Default: ``region``
 For more information, see [Working with Amazon RDS on Outposts](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-on-outposts.html) in the *Amazon RDS User Guide*.
- `ca_certificate_identifier` (String) The identifier of the CA certificate for this DB instance.
 For more information, see [Using SSL/TLS to encrypt a connection to a DB instance](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.SSL.html) in the *Amazon RDS User Guide* and [Using SSL/TLS to encrypt a connection to a DB cluster](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/UsingWithRDS.SSL.html) in the *Amazon Aurora User Guide*.
- `certificate_rotation_restart` (Boolean) Specifies whether the DB instance is restarted when you rotate your SSL/TLS certificate.
 By default, the DB instance is restarted when you rotate your SSL/TLS certificate. The certificate is not updated until the DB instance is restarted.
  Set this parameter only if you are *not* using SSL/TLS to connect to the DB instance.
  If you are using SSL/TLS to connect to the DB instance, follow the appropriate instructions for your DB engine to rotate your SSL/TLS certificate:
  +  For more information about rotating your SSL/TLS certificate for RDS DB engines, see [Rotating Your SSL/TLS Certificate.](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.SSL-certificate-rotation.html) in the *Amazon RDS User Guide.*
  +  For more information about rotating your SSL/TLS certificate for Aurora DB engines, see [Rotating Your SSL/TLS Certificate](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/UsingWithRDS.SSL-certificate-rotation.html) in the *Amazon Aurora User Guide*.
  
 This setting doesn't apply to RDS Custom DB instances.
- `character_set_name` (String) For supported engines, indicates that the DB instance should be associated with the specified character set.
  *Amazon Aurora* 
 Not applicable. The character set is managed by the DB cluster. For more information, see [AWS::RDS::DBCluster](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html).
- `copy_tags_to_snapshot` (Boolean) Specifies whether to copy tags from the DB instance to snapshots of the DB instance. By default, tags are not copied.
 This setting doesn't apply to Amazon Aurora DB instances. Copying tags to snapshots is managed by the DB cluster. Setting this value for an Aurora DB instance has no effect on the DB cluster setting.
- `custom_iam_instance_profile` (String) The instance profile associated with the underlying Amazon EC2 instance of an RDS Custom DB instance.
 This setting is required for RDS Custom.
 Constraints:
  +  The profile must exist in your account.
  +  The profile must have an IAM role that Amazon EC2 has permissions to assume.
  +  The instance profile name and the associated IAM role name must start with the prefix ``AWSRDSCustom``.
  
 For the list of permissions required for the IAM role, see [Configure IAM and your VPC](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-setup-orcl.html#custom-setup-orcl.iam-vpc) in the *Amazon RDS User Guide*.
- `database_insights_mode` (String) The mode of Database Insights to enable for the DB instance.
  Aurora DB instances inherit this value from the DB cluster, so you can't change this value.
- `db_cluster_identifier` (String) The identifier of the DB cluster that this DB instance will belong to.
 This setting doesn't apply to RDS Custom DB instances.
- `db_cluster_snapshot_identifier` (String) The identifier for the Multi-AZ DB cluster snapshot to restore from.
 For more information on Multi-AZ DB clusters, see [Multi-AZ DB cluster deployments](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html) in the *Amazon RDS User Guide*.
 Constraints:
  +  Must match the identifier of an existing Multi-AZ DB cluster snapshot.
  +  Can't be specified when ``DBSnapshotIdentifier`` is specified.
  +  Must be specified when ``DBSnapshotIdentifier`` isn't specified.
  +  If you are restoring from a shared manual Multi-AZ DB cluster snapshot, the ``DBClusterSnapshotIdentifier`` must be the ARN of the shared snapshot.
  +  Can't be the identifier of an Aurora DB cluster snapshot.
- `db_instance_class` (String) The compute and memory capacity of the DB instance, for example ``db.m5.large``. Not all DB instance classes are available in all AWS-Regions, or for all database engines. For the full list of DB instance classes, and availability for your engine, see [DB instance classes](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.DBInstanceClass.html) in the *Amazon RDS User Guide* or [Aurora DB instance classes](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Concepts.DBInstanceClass.html) in the *Amazon Aurora User Guide*.
- `db_instance_identifier` (String) A name for the DB instance. If you specify a name, AWS CloudFormation converts it to lowercase. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the DB instance. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html).
 For information about constraints that apply to DB instance identifiers, see [Naming constraints in Amazon RDS](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Limits.html#RDS_Limits.Constraints) in the *Amazon RDS User Guide*.
  If you specify a name, you can't perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
- `db_name` (String) The meaning of this parameter differs according to the database engine you use.
  If you specify the ``DBSnapshotIdentifier`` property, this property only applies to RDS for Oracle.
   *Amazon Aurora* 
 Not applicable. The database name is managed by the DB cluster.
  *Db2* 
 The name of the database to create when the DB instance is created. If this parameter isn't specified, no database is created in the DB instance.
 Constraints:
  +  Must contain 1 to 64 letters or numbers.
  +  Must begin with a letter. Subsequent characters can be letters, underscores, or digits (0-9).
  +  Can't be a word reserved by the specified database engine.
  
  *MySQL* 
 The name of the database to create when the DB instance is created. If this parameter is not specified, no database is created in the DB instance.
 Constraints:
  +  Must contain 1 to 64 letters or numbers.
  +  Can't be a word reserved by the specified database engine
  
  *MariaDB* 
 The name of the database to create when the DB instance is created. If this parameter is not specified, no database is created in the DB instance.
 Constraints:
  +  Must contain 1 to 64 letters or numbers.
  +  Can't be a word reserved by the specified database engine
  
  *PostgreSQL* 
 The name of the database to create when the DB instance is created. If this parameter is not specified, the default ``postgres`` database is created in the DB instance.
 Constraints:
  +  Must begin with a letter. Subsequent characters can be letters, underscores, or digits (0-9).
  +  Must contain 1 to 63 characters.
  +  Can't be a word reserved by the specified database engine
  
  *Oracle* 
 The Oracle System ID (SID) of the created DB instance. If you specify ``null``, the default value ``ORCL`` is used. You can't specify the string NULL, or any other reserved word, for ``DBName``. 
 Default: ``ORCL``
 Constraints:
  +  Can't be longer than 8 characters
  
  *SQL Server* 
 Not applicable. Must be null.
- `db_parameter_group_name` (String) The name of an existing DB parameter group or a reference to an [AWS::RDS::DBParameterGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbparametergroup.html) resource created in the template.
 To list all of the available DB parameter group names, use the following command:
  ``aws rds describe-db-parameter-groups --query "DBParameterGroups[].DBParameterGroupName" --output text`` 
  If any of the data members of the referenced parameter group are changed during an update, the DB instance might need to be restarted, which causes some interruption. If the parameter group contains static parameters, whether they were changed or not, an update triggers a reboot.
  If you don't specify a value for ``DBParameterGroupName`` property, the default DB parameter group for the specified engine and engine version is used.
- `db_security_groups` (List of String) A list of the DB security groups to assign to the DB instance. The list can include both the name of existing DB security groups or references to AWS::RDS::DBSecurityGroup resources created in the template.
  If you set DBSecurityGroups, you must not set VPCSecurityGroups, and vice versa. Also, note that the DBSecurityGroups property exists only for backwards compatibility with older regions and is no longer recommended for providing security information to an RDS DB instance. Instead, use VPCSecurityGroups.
  If you specify this property, AWS CloudFormation sends only the following properties (if specified) to Amazon RDS during create operations:
  +   ``AllocatedStorage`` 
  +   ``AutoMinorVersionUpgrade`` 
  +   ``AvailabilityZone`` 
  +   ``BackupRetentionPeriod`` 
  +   ``CharacterSetName`` 
  +   ``DBInstanceClass`` 
  +   ``DBName`` 
  +   ``DBParameterGroupName`` 
  +   ``DBSecurityGroups`` 
  +   ``DBSubnetGroupName`` 
  +   ``Engine`` 
  +   ``EngineVersion`` 
  +   ``Iops`` 
  +   ``LicenseModel`` 
  +   ``MasterUsername`` 
  +   ``MasterUserPassword`` 
  +   ``MultiAZ`` 
  +   ``OptionGroupName`` 
  +   ``PreferredBackupWindow`` 
  +   ``PreferredMaintenanceWindow`` 
  
 All other properties are ignored. Specify a virtual private cloud (VPC) security group if you want to submit other properties, such as ``StorageType``, ``StorageEncrypted``, or ``KmsKeyId``. If you're already using the ``DBSecurityGroups`` property, you can't use these other properties by updating your DB instance to use a VPC security group. You must recreate the DB instance.
- `db_snapshot_identifier` (String) The name or Amazon Resource Name (ARN) of the DB snapshot that's used to restore the DB instance. If you're restoring from a shared manual DB snapshot, you must specify the ARN of the snapshot.
 By specifying this property, you can create a DB instance from the specified DB snapshot. If the ``DBSnapshotIdentifier`` property is an empty string or the ``AWS::RDS::DBInstance`` declaration has no ``DBSnapshotIdentifier`` property, AWS CloudFormation creates a new database. If the property contains a value (other than an empty string), AWS CloudFormation creates a database from the specified snapshot. If a snapshot with the specified name doesn't exist, AWS CloudFormation can't create the database and it rolls back the stack.
 Some DB instance properties aren't valid when you restore from a snapshot, such as the ``MasterUsername`` and ``MasterUserPassword`` properties, and the point-in-time recovery properties ``RestoreTime`` and ``UseLatestRestorableTime``. For information about the properties that you can specify, see the [RestoreDBInstanceFromDBSnapshot](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_RestoreDBInstanceFromDBSnapshot.html) action in the *Amazon RDS API Reference*.
 After you restore a DB instance with a ``DBSnapshotIdentifier`` property, you must specify the same ``DBSnapshotIdentifier`` property for any future updates to the DB instance. When you specify this property for an update, the DB instance is not restored from the DB snapshot again, and the data in the database is not changed. However, if you don't specify the ``DBSnapshotIdentifier`` property, an empty DB instance is created, and the original DB instance is deleted. If you specify a property that is different from the previous snapshot restore property, a new DB instance is restored from the specified ``DBSnapshotIdentifier`` property, and the original DB instance is deleted.
 If you specify the ``DBSnapshotIdentifier`` property to restore a DB instance (as opposed to specifying it for DB instance updates), then don't specify the following properties:
  +   ``CharacterSetName`` 
  +   ``DBClusterIdentifier`` 
  +   ``DBName`` 
  +   ``KmsKeyId`` 
  +   ``MasterUsername`` 
  +   ``MasterUserPassword`` 
  +   ``PromotionTier`` 
  +   ``SourceDBInstanceIdentifier`` 
  +   ``SourceRegion`` 
  +  ``StorageEncrypted`` (for an unencrypted snapshot)
  +   ``Timezone`` 
  
  *Amazon Aurora* 
 Not applicable. Snapshot restore is managed by the DB cluster.
- `db_subnet_group_name` (String) A DB subnet group to associate with the DB instance. If you update this value, the new subnet group must be a subnet group in a new VPC. 
 If you don't specify a DB subnet group, RDS uses the default DB subnet group if one exists. If a default DB subnet group does not exist, and you don't specify a ``DBSubnetGroupName``, the DB instance fails to launch. 
 For more information about using Amazon RDS in a VPC, see [Amazon VPC and Amazon RDS](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_VPC.html) in the *Amazon RDS User Guide*. 
 This setting doesn't apply to Amazon Aurora DB instances. The DB subnet group is managed by the DB cluster. If specified, the setting must match the DB cluster setting.
- `db_system_id` (String) The Oracle system identifier (SID), which is the name of the Oracle database instance that manages your database files. In this context, the term "Oracle database instance" refers exclusively to the system global area (SGA) and Oracle background processes. If you don't specify a SID, the value defaults to ``RDSCDB``. The Oracle SID is also the name of your CDB.
- `dedicated_log_volume` (Boolean) Indicates whether the DB instance has a dedicated log volume (DLV) enabled.
- `delete_automated_backups` (Boolean) A value that indicates whether to remove automated backups immediately after the DB instance is deleted. This parameter isn't case-sensitive. The default is to remove automated backups immediately after the DB instance is deleted.
  *Amazon Aurora* 
 Not applicable. When you delete a DB cluster, all automated backups for that DB cluster are deleted and can't be recovered. Manual DB cluster snapshots of the DB cluster are not deleted.
- `deletion_protection` (Boolean) Specifies whether the DB instance has deletion protection enabled. The database can't be deleted when deletion protection is enabled. By default, deletion protection isn't enabled. For more information, see [Deleting a DB Instance](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_DeleteInstance.html).
 This setting doesn't apply to Amazon Aurora DB instances. You can enable or disable deletion protection for the DB cluster. For more information, see ``CreateDBCluster``. DB instances in a DB cluster can be deleted even when deletion protection is enabled for the DB cluster.
- `domain` (String) The Active Directory directory ID to create the DB instance in. Currently, only Db2, MySQL, Microsoft SQL Server, Oracle, and PostgreSQL DB instances can be created in an Active Directory Domain.
 For more information, see [Kerberos Authentication](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/kerberos-authentication.html) in the *Amazon RDS User Guide*.
- `domain_auth_secret_arn` (String) The ARN for the Secrets Manager secret with the credentials for the user joining the domain.
 Example: ``arn:aws:secretsmanager:region:account-number:secret:myselfmanagedADtestsecret-123456``
- `domain_dns_ips` (List of String) The IPv4 DNS IP addresses of your primary and secondary Active Directory domain controllers.
 Constraints:
  +  Two IP addresses must be provided. If there isn't a secondary domain controller, use the IP address of the primary domain controller for both entries in the list.
  
 Example: ``123.124.125.126,234.235.236.237``
- `domain_fqdn` (String) The fully qualified domain name (FQDN) of an Active Directory domain.
 Constraints:
  +  Can't be longer than 64 characters.
  
 Example: ``mymanagedADtest.mymanagedAD.mydomain``
- `domain_iam_role_name` (String) The name of the IAM role to use when making API calls to the Directory Service.
 This setting doesn't apply to the following DB instances:
  +  Amazon Aurora (The domain is managed by the DB cluster.)
  +  RDS Custom
- `domain_ou` (String) The Active Directory organizational unit for your DB instance to join.
 Constraints:
  +  Must be in the distinguished name format.
  +  Can't be longer than 64 characters.
  
 Example: ``OU=mymanagedADtestOU,DC=mymanagedADtest,DC=mymanagedAD,DC=mydomain``
- `enable_cloudwatch_logs_exports` (List of String) The list of log types that need to be enabled for exporting to CloudWatch Logs. The values in the list depend on the DB engine being used. For more information, see [Publishing Database Logs to Amazon CloudWatch Logs](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_LogAccess.html#USER_LogAccess.Procedural.UploadtoCloudWatch) in the *Amazon Relational Database Service User Guide*.
  *Amazon Aurora* 
 Not applicable. CloudWatch Logs exports are managed by the DB cluster. 
  *Db2* 
 Valid values: ``diag.log``, ``notify.log``
  *MariaDB* 
 Valid values: ``audit``, ``error``, ``general``, ``slowquery``
  *Microsoft SQL Server* 
 Valid values: ``agent``, ``error``
  *MySQL* 
 Valid values: ``audit``, ``error``, ``general``, ``slowquery``
  *Oracle* 
 Valid values: ``alert``, ``audit``, ``listener``, ``trace``, ``oemagent``
  *PostgreSQL* 
 Valid values: ``postgresql``, ``upgrade``
- `enable_iam_database_authentication` (Boolean) A value that indicates whether to enable mapping of AWS Identity and Access Management (IAM) accounts to database accounts. By default, mapping is disabled.
 This property is supported for RDS for MariaDB, RDS for MySQL, and RDS for PostgreSQL. For more information, see [IAM Database Authentication for MariaDB, MySQL, and PostgreSQL](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.IAMDBAuth.html) in the *Amazon RDS User Guide.*
  *Amazon Aurora* 
 Not applicable. Mapping AWS IAM accounts to database accounts is managed by the DB cluster.
- `enable_performance_insights` (Boolean) Specifies whether to enable Performance Insights for the DB instance. For more information, see [Using Amazon Performance Insights](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PerfInsights.html) in the *Amazon RDS User Guide*.
 This setting doesn't apply to RDS Custom DB instances.
- `engine` (String) The name of the database engine to use for this DB instance. Not every database engine is available in every AWS Region.
 This property is required when creating a DB instance.
  You can convert an Oracle database from the non-CDB architecture to the container database (CDB) architecture by updating the ``Engine`` value in your templates from ``oracle-ee`` to ``oracle-ee-cdb`` or from ``oracle-se2`` to ``oracle-se2-cdb``. Converting to the CDB architecture requires an interruption.
  Valid Values:
  +  ``aurora-mysql`` (for Aurora MySQL DB instances)
  +  ``aurora-postgresql`` (for Aurora PostgreSQL DB instances)
  +  ``custom-oracle-ee`` (for RDS Custom for Oracle DB instances)
  +  ``custom-oracle-ee-cdb`` (for RDS Custom for Oracle DB instances)
  +  ``custom-sqlserver-ee`` (for RDS Custom for SQL Server DB instances)
  +  ``custom-sqlserver-se`` (for RDS Custom for SQL Server DB instances)
  +  ``custom-sqlserver-web`` (for RDS Custom for SQL Server DB instances)
  +   ``db2-ae`` 
  +   ``db2-se`` 
  +   ``mariadb`` 
  +   ``mysql`` 
  +   ``oracle-ee`` 
  +   ``oracle-ee-cdb`` 
  +   ``oracle-se2`` 
  +   ``oracle-se2-cdb`` 
  +   ``postgres`` 
  +   ``sqlserver-ee`` 
  +   ``sqlserver-se`` 
  +   ``sqlserver-ex`` 
  +   ``sqlserver-web``
- `engine_lifecycle_support` (String) The life cycle type for this DB instance.
  By default, this value is set to ``open-source-rds-extended-support``, which enrolls your DB instance into Amazon RDS Extended Support. At the end of standard support, you can avoid charges for Extended Support by setting the value to ``open-source-rds-extended-support-disabled``. In this case, creating the DB instance will fail if the DB major version is past its end of standard support date.
  This setting applies only to RDS for MySQL and RDS for PostgreSQL. For Amazon Aurora DB instances, the life cycle type is managed by the DB cluster.
 You can use this setting to enroll your DB instance into Amazon RDS Extended Support. With RDS Extended Support, you can run the selected major engine version on your DB instance past the end of standard support for that engine version. For more information, see [Amazon RDS Extended Support with Amazon RDS](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/extended-support.html) in the *Amazon RDS User Guide*.
 Valid Values: ``open-source-rds-extended-support | open-source-rds-extended-support-disabled``
 Default: ``open-source-rds-extended-support``
- `engine_version` (String) The version number of the database engine to use.
 For a list of valid engine versions, use the ``DescribeDBEngineVersions`` action.
 The following are the database engines and links to information about the major and minor versions that are available with Amazon RDS. Not every database engine is available for every AWS Region.
  *Amazon Aurora* 
 Not applicable. The version number of the database engine to be used by the DB instance is managed by the DB cluster.
  *Db2* 
 See [Amazon RDS for Db2](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Db2.html#Db2.Concepts.VersionMgmt) in the *Amazon RDS User Guide.*
  *MariaDB* 
 See [MariaDB on Amazon RDS Versions](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_MariaDB.html#MariaDB.Concepts.VersionMgmt) in the *Amazon RDS User Guide.*
  *Microsoft SQL Server* 
 See [Microsoft SQL Server Versions on Amazon RDS](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_SQLServer.html#SQLServer.Concepts.General.VersionSupport) in the *Amazon RDS User Guide.*
  *MySQL* 
 See [MySQL on Amazon RDS Versions](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_MySQL.html#MySQL.Concepts.VersionMgmt) in the *Amazon RDS User Guide.*
  *Oracle* 
 See [Oracle Database Engine Release Notes](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.Oracle.PatchComposition.html) in the *Amazon RDS User Guide.*
  *PostgreSQL* 
 See [Supported PostgreSQL Database Versions](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_PostgreSQL.html#PostgreSQL.Concepts.General.DBVersions) in the *Amazon RDS User Guide.*
- `iops` (Number) The number of I/O operations per second (IOPS) that the database provisions. The value must be equal to or greater than 1000. 
 If you specify this property, you must follow the range of allowed ratios of your requested IOPS rate to the amount of storage that you allocate (IOPS to allocated storage). For example, you can provision an Oracle database instance with 1000 IOPS and 200 GiB of storage (a ratio of 5:1), or specify 2000 IOPS with 200 GiB of storage (a ratio of 10:1). For more information, see [Amazon RDS Provisioned IOPS Storage to Improve Performance](https://docs.aws.amazon.com/AmazonRDS/latest/DeveloperGuide/CHAP_Storage.html#USER_PIOPS) in the *Amazon RDS User Guide*.
  If you specify ``io1`` for the ``StorageType`` property, then you must also specify the ``Iops`` property.
  Constraints:
  +  For RDS for Db2, MariaDB, MySQL, Oracle, and PostgreSQL - Must be a multiple between .5 and 50 of the storage amount for the DB instance.
  +  For RDS for SQL Server - Must be a multiple between 1 and 50 of the storage amount for the DB instance.
- `kms_key_id` (String) The ARN of the AWS KMS key that's used to encrypt the DB instance, such as ``arn:aws:kms:us-east-1:012345678910:key/abcd1234-a123-456a-a12b-a123b4cd56ef``. If you enable the StorageEncrypted property but don't specify this property, AWS CloudFormation uses the default KMS key. If you specify this property, you must set the StorageEncrypted property to true. 
 If you specify the ``SourceDBInstanceIdentifier`` or ``SourceDbiResourceId`` property, don't specify this property. The value is inherited from the source DB instance, and if the DB instance is encrypted, the specified ``KmsKeyId`` property is used. However, if the source DB instance is in a different AWS Region, you must specify a KMS key ID.
 If you specify the ``SourceDBInstanceAutomatedBackupsArn`` property, don't specify this property. The value is inherited from the source DB instance automated backup, and if the automated backup is encrypted, the specified ``KmsKeyId`` property is used.
 If you create an encrypted read replica in a different AWS Region, then you must specify a KMS key for the destination AWS Region. KMS encryption keys are specific to the region that they're created in, and you can't use encryption keys from one region in another region.
 If you specify the ``DBSnapshotIdentifier`` property, don't specify this property. The ``StorageEncrypted`` property value is inherited from the snapshot. If the DB instance is encrypted, the specified ``KmsKeyId`` property is also inherited from the snapshot.
 If you specify ``DBSecurityGroups``, AWS CloudFormation ignores this property. To specify both a security group and this property, you must use a VPC security group. For more information about Amazon RDS and VPC, see [Using Amazon RDS with Amazon VPC](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_VPC.html) in the *Amazon RDS User Guide*.
  *Amazon Aurora* 
 Not applicable. The KMS key identifier is managed by the DB cluster.
- `license_model` (String) License model information for this DB instance.
  Valid Values:
  +  Aurora MySQL - ``general-public-license``
  +  Aurora PostgreSQL - ``postgresql-license``
  +  RDS for Db2 - ``bring-your-own-license``. For more information about RDS for Db2 licensing, see [](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/db2-licensing.html) in the *Amazon RDS User Guide.*
  +  RDS for MariaDB - ``general-public-license``
  +  RDS for Microsoft SQL Server - ``license-included``
  +  RDS for MySQL - ``general-public-license``
  +  RDS for Oracle - ``bring-your-own-license`` or ``license-included``
  +  RDS for PostgreSQL - ``postgresql-license``
  
  If you've specified ``DBSecurityGroups`` and then you update the license model, AWS CloudFormation replaces the underlying DB instance. This will incur some interruptions to database availability.
- `manage_master_user_password` (Boolean) Specifies whether to manage the master user password with AWS Secrets Manager.
 For more information, see [Password management with Secrets Manager](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-secrets-manager.html) in the *Amazon RDS User Guide.*
 Constraints:
  +  Can't manage the master user password with AWS Secrets Manager if ``MasterUserPassword`` is specified.
- `master_user_password` (String) The password for the master user. The password can include any printable ASCII character except "/", """, or "@".
  *Amazon Aurora* 
 Not applicable. The password for the master user is managed by the DB cluster.
  *RDS for Db2* 
 Must contain from 8 to 255 characters.
  *RDS for MariaDB* 
 Constraints: Must contain from 8 to 41 characters.
  *RDS for Microsoft SQL Server* 
 Constraints: Must contain from 8 to 128 characters.
  *RDS for MySQL* 
 Constraints: Must contain from 8 to 41 characters.
  *RDS for Oracle* 
 Constraints: Must contain from 8 to 30 characters.
  *RDS for PostgreSQL* 
 Constraints: Must contain from 8 to 128 characters.
- `master_user_secret` (Attributes) The secret managed by RDS in AWS Secrets Manager for the master user password.
 For more information, see [Password management with Secrets Manager](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-secrets-manager.html) in the *Amazon RDS User Guide.* (see [below for nested schema](#nestedatt--master_user_secret))
- `master_username` (String) The master user name for the DB instance.
  If you specify the ``SourceDBInstanceIdentifier`` or ``DBSnapshotIdentifier`` property, don't specify this property. The value is inherited from the source DB instance or snapshot.
 When migrating a self-managed Db2 database, we recommend that you use the same master username as your self-managed Db2 instance name.
   *Amazon Aurora* 
 Not applicable. The name for the master user is managed by the DB cluster. 
  *RDS for Db2* 
 Constraints:
  +  Must be 1 to 16 letters or numbers.
  +  First character must be a letter.
  +  Can't be a reserved word for the chosen database engine.
  
  *RDS for MariaDB* 
 Constraints:
  +  Must be 1 to 16 letters or numbers.
  +  Can't be a reserved word for the chosen database engine.
  
  *RDS for Microsoft SQL Server* 
 Constraints:
  +  Must be 1 to 128 letters or numbers.
  +  First character must be a letter.
  +  Can't be a reserved word for the chosen database engine.
  
  *RDS for MySQL* 
 Constraints:
  +  Must be 1 to 16 letters or numbers.
  +  First character must be a letter.
  +  Can't be a reserved word for the chosen database engine.
  
  *RDS for Oracle* 
 Constraints:
  +  Must be 1 to 30 letters or numbers.
  +  First character must be a letter.
  +  Can't be a reserved word for the chosen database engine.
  
  *RDS for PostgreSQL* 
 Constraints:
  +  Must be 1 to 63 letters or numbers.
  +  First character must be a letter.
  +  Can't be a reserved word for the chosen database engine.
- `max_allocated_storage` (Number) The upper limit in gibibytes (GiB) to which Amazon RDS can automatically scale the storage of the DB instance.
 For more information about this setting, including limitations that apply to it, see [Managing capacity automatically with Amazon RDS storage autoscaling](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PIOPS.StorageTypes.html#USER_PIOPS.Autoscaling) in the *Amazon RDS User Guide*.
 This setting doesn't apply to the following DB instances:
  +  Amazon Aurora (Storage is managed by the DB cluster.)
  +  RDS Custom
- `monitoring_interval` (Number) The interval, in seconds, between points when Enhanced Monitoring metrics are collected for the DB instance. To disable collection of Enhanced Monitoring metrics, specify ``0``.
 If ``MonitoringRoleArn`` is specified, then you must set ``MonitoringInterval`` to a value other than ``0``.
 This setting doesn't apply to RDS Custom DB instances.
 Valid Values: ``0 | 1 | 5 | 10 | 15 | 30 | 60``
 Default: ``0``
- `monitoring_role_arn` (String) The ARN for the IAM role that permits RDS to send enhanced monitoring metrics to Amazon CloudWatch Logs. For example, ``arn:aws:iam:123456789012:role/emaccess``. For information on creating a monitoring role, see [Setting Up and Enabling Enhanced Monitoring](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Monitoring.OS.html#USER_Monitoring.OS.Enabling) in the *Amazon RDS User Guide*.
 If ``MonitoringInterval`` is set to a value other than ``0``, then you must supply a ``MonitoringRoleArn`` value.
 This setting doesn't apply to RDS Custom DB instances.
- `multi_az` (Boolean) Specifies whether the DB instance is a Multi-AZ deployment. You can't set the ``AvailabilityZone`` parameter if the DB instance is a Multi-AZ deployment.
 This setting doesn't apply to Amazon Aurora because the DB instance Availability Zones (AZs) are managed by the DB cluster.
- `nchar_character_set_name` (String) The name of the NCHAR character set for the Oracle DB instance.
 This setting doesn't apply to RDS Custom DB instances.
- `network_type` (String) The network type of the DB instance.
 Valid values:
  +   ``IPV4`` 
  +   ``DUAL`` 
  
 The network type is determined by the ``DBSubnetGroup`` specified for the DB instance. A ``DBSubnetGroup`` can support only the IPv4 protocol or the IPv4 and IPv6 protocols (``DUAL``).
 For more information, see [Working with a DB instance in a VPC](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_VPC.WorkingWithRDSInstanceinaVPC.html) in the *Amazon RDS User Guide.*
- `option_group_name` (String) Indicates that the DB instance should be associated with the specified option group.
 Permanent options, such as the TDE option for Oracle Advanced Security TDE, can't be removed from an option group. Also, that option group can't be removed from a DB instance once it is associated with a DB instance.
- `performance_insights_kms_key_id` (String) The AWS KMS key identifier for encryption of Performance Insights data.
 The KMS key identifier is the key ARN, key ID, alias ARN, or alias name for the KMS key.
 If you do not specify a value for ``PerformanceInsightsKMSKeyId``, then Amazon RDS uses your default KMS key. There is a default KMS key for your AWS account. Your AWS account has a different default KMS key for each AWS Region.
 For information about enabling Performance Insights, see [EnablePerformanceInsights](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html#cfn-rds-dbinstance-enableperformanceinsights).
- `performance_insights_retention_period` (Number) The number of days to retain Performance Insights data. When creating a DB instance without enabling Performance Insights, you can't specify the parameter ``PerformanceInsightsRetentionPeriod``.
 This setting doesn't apply to RDS Custom DB instances.
 Valid Values:
  +   ``7`` 
  +  *month* * 31, where *month* is a number of months from 1-23. Examples: ``93`` (3 months * 31), ``341`` (11 months * 31), ``589`` (19 months * 31)
  +   ``731`` 
  
 Default: ``7`` days
 If you specify a retention period that isn't valid, such as ``94``, Amazon RDS returns an error.
- `port` (String) The port number on which the database accepts connections.
 This setting doesn't apply to Aurora DB instances. The port number is managed by the cluster.
 Valid Values: ``1150-65535``
 Default:
  +  RDS for Db2 - ``50000``
  +  RDS for MariaDB - ``3306``
  +  RDS for Microsoft SQL Server - ``1433``
  +  RDS for MySQL - ``3306``
  +  RDS for Oracle - ``1521``
  +  RDS for PostgreSQL - ``5432``
  
 Constraints:
  +  For RDS for Microsoft SQL Server, the value can't be ``1234``, ``1434``, ``3260``, ``3343``, ``3389``, ``47001``, or ``49152-49156``.
- `preferred_backup_window` (String) The daily time range during which automated backups are created if automated backups are enabled, using the ``BackupRetentionPeriod`` parameter. For more information, see [Backup Window](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_WorkingWithAutomatedBackups.html#USER_WorkingWithAutomatedBackups.BackupWindow) in the *Amazon RDS User Guide.*
 Constraints:
  +  Must be in the format ``hh24:mi-hh24:mi``.
  +  Must be in Universal Coordinated Time (UTC).
  +  Must not conflict with the preferred maintenance window.
  +  Must be at least 30 minutes.
  
  *Amazon Aurora* 
 Not applicable. The daily time range for creating automated backups is managed by the DB cluster.
- `preferred_maintenance_window` (String) The weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).
 Format: ``ddd:hh24:mi-ddd:hh24:mi``
 The default is a 30-minute window selected at random from an 8-hour block of time for each AWS Region, occurring on a random day of the week. To see the time blocks available, see [Maintaining a DB instance](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_UpgradeDBInstance.Maintenance.html#AdjustingTheMaintenanceWindow) in the *Amazon RDS User Guide.*
  This property applies when AWS CloudFormation initially creates the DB instance. If you use AWS CloudFormation to update the DB instance, those updates are applied immediately.
  Constraints: Minimum 30-minute window.
- `processor_features` (Attributes List) The number of CPU cores and the number of threads per core for the DB instance class of the DB instance.
 This setting doesn't apply to Amazon Aurora or RDS Custom DB instances. (see [below for nested schema](#nestedatt--processor_features))
- `promotion_tier` (Number) The order of priority in which an Aurora Replica is promoted to the primary instance after a failure of the existing primary instance. For more information, see [Fault Tolerance for an Aurora DB Cluster](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Concepts.AuroraHighAvailability.html#Aurora.Managing.FaultTolerance) in the *Amazon Aurora User Guide*.
 This setting doesn't apply to RDS Custom DB instances.
 Default: ``1``
 Valid Values: ``0 - 15``
- `publicly_accessible` (Boolean) Indicates whether the DB instance is an internet-facing instance. If you specify true, AWS CloudFormation creates an instance with a publicly resolvable DNS name, which resolves to a public IP address. If you specify false, AWS CloudFormation creates an internal instance with a DNS name that resolves to a private IP address. 
 The default behavior value depends on your VPC setup and the database subnet group. For more information, see the ``PubliclyAccessible`` parameter in the [CreateDBInstance](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_CreateDBInstance.html) in the *Amazon RDS API Reference*.
- `replica_mode` (String) The open mode of an Oracle read replica. For more information, see [Working with Oracle Read Replicas for Amazon RDS](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/oracle-read-replicas.html) in the *Amazon RDS User Guide*.
 This setting is only supported in RDS for Oracle.
 Default: ``open-read-only``
 Valid Values: ``open-read-only`` or ``mounted``
- `restore_time` (String) The date and time to restore from. This parameter applies to point-in-time recovery. For more information, see [Restoring a DB instance to a specified time](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PIT.html) in the in the *Amazon RDS User Guide*.
 Constraints:
  +  Must be a time in Universal Coordinated Time (UTC) format.
  +  Must be before the latest restorable time for the DB instance.
  +  Can't be specified if the ``UseLatestRestorableTime`` parameter is enabled.
  
 Example: ``2009-09-07T23:45:00Z``
- `source_db_cluster_identifier` (String) The identifier of the Multi-AZ DB cluster that will act as the source for the read replica. Each DB cluster can have up to 15 read replicas.
 Constraints:
  +  Must be the identifier of an existing Multi-AZ DB cluster.
  +  Can't be specified if the ``SourceDBInstanceIdentifier`` parameter is also specified.
  +  The specified DB cluster must have automatic backups enabled, that is, its backup retention period must be greater than 0.
  +  The source DB cluster must be in the same AWS-Region as the read replica. Cross-Region replication isn't supported.
- `source_db_instance_automated_backups_arn` (String) The Amazon Resource Name (ARN) of the replicated automated backups from which to restore, for example, ``arn:aws:rds:us-east-1:123456789012:auto-backup:ab-L2IJCEXJP7XQ7HOJ4SIEXAMPLE``.
 This setting doesn't apply to RDS Custom.
- `source_db_instance_identifier` (String) If you want to create a read replica DB instance, specify the ID of the source DB instance. Each DB instance can have a limited number of read replicas. For more information, see [Working with Read Replicas](https://docs.aws.amazon.com/AmazonRDS/latest/DeveloperGuide/USER_ReadRepl.html) in the *Amazon RDS User Guide*.
 For information about constraints that apply to DB instance identifiers, see [Naming constraints in Amazon RDS](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Limits.html#RDS_Limits.Constraints) in the *Amazon RDS User Guide*.
 The ``SourceDBInstanceIdentifier`` property determines whether a DB instance is a read replica. If you remove the ``SourceDBInstanceIdentifier`` property from your template and then update your stack, AWS CloudFormation promotes the read replica to a standalone DB instance.
 If you specify the ``UseLatestRestorableTime`` or ``RestoreTime`` properties in conjunction with the ``SourceDBInstanceIdentifier`` property, RDS restores the DB instance to the requested point in time, thereby creating a new DB instance.
   +  If you specify a source DB instance that uses VPC security groups, we recommend that you specify the ``VPCSecurityGroups`` property. If you don't specify the property, the read replica inherits the value of the ``VPCSecurityGroups`` property from the source DB when you create the replica. However, if you update the stack, AWS CloudFormation reverts the replica's ``VPCSecurityGroups`` property to the default value because it's not defined in the stack's template. This change might cause unexpected issues.
  +  Read replicas don't support deletion policies. AWS CloudFormation ignores any deletion policy that's associated with a read replica.
  +  If you specify ``SourceDBInstanceIdentifier``, don't specify the ``DBSnapshotIdentifier`` property. You can't create a read replica from a snapshot.
  +  Don't set the ``BackupRetentionPeriod``, ``DBName``, ``MasterUsername``, ``MasterUserPassword``, and ``PreferredBackupWindow`` properties. The database attributes are inherited from the source DB instance, and backups are disabled for read replicas.
  +  If the source DB instance is in a different region than the read replica, specify the source region in ``SourceRegion``, and specify an ARN for a valid DB instance in ``SourceDBInstanceIdentifier``. For more information, see [Constructing a Amazon RDS Amazon Resource Name (ARN)](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html#USER_Tagging.ARN) in the *Amazon RDS User Guide*.
  +  For DB instances in Amazon Aurora clusters, don't specify this property. Amazon RDS automatically assigns writer and reader DB instances.
- `source_dbi_resource_id` (String) The resource ID of the source DB instance from which to restore.
- `source_region` (String) The ID of the region that contains the source DB instance for the read replica.
- `storage_encrypted` (Boolean) A value that indicates whether the DB instance is encrypted. By default, it isn't encrypted.
 If you specify the ``KmsKeyId`` property, then you must enable encryption.
 If you specify the ``SourceDBInstanceIdentifier`` or ``SourceDbiResourceId`` property, don't specify this property. The value is inherited from the source DB instance, and if the DB instance is encrypted, the specified ``KmsKeyId`` property is used.
 If you specify the ``SourceDBInstanceAutomatedBackupsArn`` property, don't specify this property. The value is inherited from the source DB instance automated backup. 
 If you specify ``DBSnapshotIdentifier`` property, don't specify this property. The value is inherited from the snapshot.
  *Amazon Aurora* 
 Not applicable. The encryption for DB instances is managed by the DB cluster.
- `storage_throughput` (Number) Specifies the storage throughput value, in mebibyte per second (MiBps), for the DB instance. This setting applies only to the ``gp3`` storage type. 
 This setting doesn't apply to RDS Custom or Amazon Aurora.
- `storage_type` (String) The storage type to associate with the DB instance.
 If you specify ``io1``, ``io2``, or ``gp3``, you must also include a value for the ``Iops`` parameter.
 This setting doesn't apply to Amazon Aurora DB instances. Storage is managed by the DB cluster.
 Valid Values: ``gp2 | gp3 | io1 | io2 | standard``
 Default: ``io1``, if the ``Iops`` parameter is specified. Otherwise, ``gp3``.
- `tags` (Attributes List) Tags to assign to the DB instance. (see [below for nested schema](#nestedatt--tags))
- `tde_credential_arn` (String)
- `tde_credential_password` (String)
- `timezone` (String) The time zone of the DB instance. The time zone parameter is currently supported only by [RDS for Db2](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/db2-time-zone) and [RDS for SQL Server](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_SQLServer.html#SQLServer.Concepts.General.TimeZone).
- `use_default_processor_features` (Boolean) Specifies whether the DB instance class of the DB instance uses its default processor features.
 This setting doesn't apply to RDS Custom DB instances.
- `use_latest_restorable_time` (Boolean) Specifies whether the DB instance is restored from the latest backup time. By default, the DB instance isn't restored from the latest backup time. This parameter applies to point-in-time recovery. For more information, see [Restoring a DB instance to a specified time](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PIT.html) in the in the *Amazon RDS User Guide*.
 Constraints:
  +  Can't be specified if the ``RestoreTime`` parameter is provided.
- `vpc_security_groups` (List of String) A list of the VPC security group IDs to assign to the DB instance. The list can include both the physical IDs of existing VPC security groups and references to [AWS::EC2::SecurityGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group.html) resources created in the template.
 If you plan to update the resource, don't specify VPC security groups in a shared VPC.
  If you set ``VPCSecurityGroups``, you must not set [DBSecurityGroups](https://docs.aws.amazon.com//AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html#cfn-rds-dbinstance-dbsecuritygroups), and vice versa.
  You can migrate a DB instance in your stack from an RDS DB security group to a VPC security group, but keep the following in mind:
  +  You can't revert to using an RDS security group after you establish a VPC security group membership.
  +  When you migrate your DB instance to VPC security groups, if your stack update rolls back because the DB instance update fails or because an update fails in another AWS CloudFormation resource, the rollback fails because it can't revert to an RDS security group.
  +  To use the properties that are available when you use a VPC security group, you must recreate the DB instance. If you don't, AWS CloudFormation submits only the property values that are listed in the [DBSecurityGroups](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html#cfn-rds-dbinstance-dbsecuritygroups) property.
  
  To avoid this situation, migrate your DB instance to using VPC security groups only when that is the only change in your stack template. 
  *Amazon Aurora* 
 Not applicable. The associated list of EC2 VPC security groups is managed by the DB cluster. If specified, the setting must match the DB cluster setting.

### Read-Only

- `certificate_details` (Attributes) The details of the DB instance?s server certificate.
 For more information, see [Using SSL/TLS to encrypt a connection to a DB instance](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.SSL.html) in the *Amazon RDS User Guide* and [Using SSL/TLS to encrypt a connection to a DB cluster](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/UsingWithRDS.SSL.html) in the *Amazon Aurora User Guide*. (see [below for nested schema](#nestedatt--certificate_details))
- `db_instance_arn` (String)
- `db_instance_status` (String)
- `dbi_resource_id` (String)
- `endpoint` (Attributes) This data type represents the information you need to connect to an Amazon RDS DB instance. This data type is used as a response element in the following actions:
  +   ``CreateDBInstance`` 
  +   ``DescribeDBInstances`` 
  +   ``DeleteDBInstance`` 
  
 For the data structure that represents Amazon Aurora DB cluster endpoints, see ``DBClusterEndpoint``. (see [below for nested schema](#nestedatt--endpoint))
- `id` (String) Uniquely identifies the resource.
- `instance_create_time` (String)
- `is_storage_config_upgrade_available` (Boolean)
- `latest_restorable_time` (String)
- `listener_endpoint` (Attributes) This data type represents the information you need to connect to an Amazon RDS DB instance. This data type is used as a response element in the following actions:
  +   ``CreateDBInstance`` 
  +   ``DescribeDBInstances`` 
  +   ``DeleteDBInstance`` 
  
 For the data structure that represents Amazon Aurora DB cluster endpoints, see ``DBClusterEndpoint``. (see [below for nested schema](#nestedatt--listener_endpoint))
- `read_replica_db_cluster_identifiers` (List of String)
- `read_replica_db_instance_identifiers` (List of String)

<a id="nestedatt--associated_roles"></a>
### Nested Schema for `associated_roles`

Optional:

- `feature_name` (String) The name of the feature associated with the AWS Identity and Access Management (IAM) role. IAM roles that are associated with a DB instance grant permission for the DB instance to access other AWS services on your behalf. For the list of supported feature names, see the ``SupportedFeatureNames`` description in [DBEngineVersion](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DBEngineVersion.html) in the *Amazon RDS API Reference*.
- `role_arn` (String) The Amazon Resource Name (ARN) of the IAM role that is associated with the DB instance.


<a id="nestedatt--master_user_secret"></a>
### Nested Schema for `master_user_secret`

Optional:

- `kms_key_id` (String) The AWS KMS key identifier that is used to encrypt the secret.

Read-Only:

- `secret_arn` (String) The Amazon Resource Name (ARN) of the secret. This parameter is a return value that you can retrieve using the ``Fn::GetAtt`` intrinsic function. For more information, see [Return values](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#aws-resource-rds-dbinstance-return-values).


<a id="nestedatt--processor_features"></a>
### Nested Schema for `processor_features`

Optional:

- `name` (String) The name of the processor feature. Valid names are ``coreCount`` and ``threadsPerCore``.
- `value` (String) The value of a processor feature.


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String) A key is the required name of the tag. The string value can be from 1 to 128 Unicode characters in length and can't be prefixed with ``aws:`` or ``rds:``. The string can only contain only the set of Unicode letters, digits, white-space, '_', '.', ':', '/', '=', '+', '-', '@' (Java regex: "^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-@]*)$").
- `value` (String) A value is the optional value of the tag. The string value can be from 1 to 256 Unicode characters in length and can't be prefixed with ``aws:`` or ``rds:``. The string can only contain only the set of Unicode letters, digits, white-space, '_', '.', ':', '/', '=', '+', '-', '@' (Java regex: "^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-@]*)$").


<a id="nestedatt--certificate_details"></a>
### Nested Schema for `certificate_details`

Read-Only:

- `ca_identifier` (String) The CA identifier of the CA certificate used for the DB instance's server certificate.
- `valid_till` (String) The expiration date of the DB instance?s server certificate.


<a id="nestedatt--endpoint"></a>
### Nested Schema for `endpoint`

Read-Only:

- `address` (String) Specifies the DNS address of the DB instance.
- `hosted_zone_id` (String) Specifies the ID that Amazon Route 53 assigns when you create a hosted zone.
- `port` (String) Specifies the port that the database engine is listening on.


<a id="nestedatt--listener_endpoint"></a>
### Nested Schema for `listener_endpoint`

Read-Only:

- `address` (String) Specifies the DNS address of the DB instance.
- `hosted_zone_id` (String) Specifies the ID that Amazon Route 53 assigns when you create a hosted zone.
- `port` (String) Specifies the port that the database engine is listening on.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_rds_db_instance.example
  id = "db_instance_identifier"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_rds_db_instance.example "db_instance_identifier"
```