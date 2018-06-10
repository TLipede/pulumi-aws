// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elasticache

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to get information about an Elasticache Cluster
func Lookupluster(ctx *pulumi.Context, args *GetClusterArgs) (*GetClusterResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["clusterId"] = args.ClusterId
		inputs["tags"] = args.Tags
	}
	outputs, err := ctx.Invoke("aws:elasticache/getCluster:getCluster", inputs)
	if err != nil {
		return nil, err
	}
	ret := GetClusterResult{}
	if v, ok := outputs["arn"]; ok {
		ret.Arn = v
	}
	if v, ok := outputs["availabilityZone"]; ok {
		ret.AvailabilityZone = v
	}
	if v, ok := outputs["cacheNodes"]; ok {
		ret.CacheNodes = v
	}
	if v, ok := outputs["clusterAddress"]; ok {
		ret.ClusterAddress = v
	}
	if v, ok := outputs["configurationEndpoint"]; ok {
		ret.ConfigurationEndpoint = v
	}
	if v, ok := outputs["engine"]; ok {
		ret.Engine = v
	}
	if v, ok := outputs["engineVersion"]; ok {
		ret.EngineVersion = v
	}
	if v, ok := outputs["maintenanceWindow"]; ok {
		ret.MaintenanceWindow = v
	}
	if v, ok := outputs["nodeType"]; ok {
		ret.NodeType = v
	}
	if v, ok := outputs["notificationTopicArn"]; ok {
		ret.NotificationTopicArn = v
	}
	if v, ok := outputs["numCacheNodes"]; ok {
		ret.NumCacheNodes = v
	}
	if v, ok := outputs["parameterGroupName"]; ok {
		ret.ParameterGroupName = v
	}
	if v, ok := outputs["port"]; ok {
		ret.Port = v
	}
	if v, ok := outputs["replicationGroupId"]; ok {
		ret.ReplicationGroupId = v
	}
	if v, ok := outputs["securityGroupIds"]; ok {
		ret.SecurityGroupIds = v
	}
	if v, ok := outputs["securityGroupNames"]; ok {
		ret.SecurityGroupNames = v
	}
	if v, ok := outputs["snapshotRetentionLimit"]; ok {
		ret.SnapshotRetentionLimit = v
	}
	if v, ok := outputs["snapshotWindow"]; ok {
		ret.SnapshotWindow = v
	}
	if v, ok := outputs["subnetGroupName"]; ok {
		ret.SubnetGroupName = v
	}
	if v, ok := outputs["tags"]; ok {
		ret.Tags = v
	}
	return &ret, nil
}

// A collection of arguments for invoking getCluster.
type GetClusterArgs struct {
	// Group identifier.
	ClusterId interface{}
	Tags interface{}
}

// A collection of values returned by getCluster.
type GetClusterResult struct {
	Arn interface{}
	// The Availability Zone for the cache cluster.
	AvailabilityZone interface{}
	// List of node objects including `id`, `address`, `port` and `availability_zone`.
	// Referenceable e.g. as `${data.aws_elasticache_cluster.bar.cache_nodes.0.address}`
	CacheNodes interface{}
	// The DNS name of the cache cluster without the port appended.
	ClusterAddress interface{}
	// The configuration endpoint to allow host discovery.
	ConfigurationEndpoint interface{}
	// Name of the cache engine.
	Engine interface{}
	// Version number of the cache engine.
	EngineVersion interface{}
	// Specifies the weekly time range for when maintenance
	// on the cache cluster is performed.
	MaintenanceWindow interface{}
	// The cluster node type.
	NodeType interface{}
	// An Amazon Resource Name (ARN) of an
	// SNS topic that ElastiCache notifications get sent to.
	NotificationTopicArn interface{}
	// The number of cache nodes that the cache cluster has.
	NumCacheNodes interface{}
	// Name of the parameter group associated with this cache cluster.
	ParameterGroupName interface{}
	// The port number on which each of the cache nodes will
	// accept connections.
	Port interface{}
	// The replication group to which this cache cluster belongs.
	ReplicationGroupId interface{}
	// List VPC security groups associated with the cache cluster.
	SecurityGroupIds interface{}
	// List of security group names associated with this cache cluster.
	SecurityGroupNames interface{}
	// The number of days for which ElastiCache will
	// retain automatic cache cluster snapshots before deleting them.
	SnapshotRetentionLimit interface{}
	// The daily time range (in UTC) during which ElastiCache will
	// begin taking a daily snapshot of the cache cluster.
	SnapshotWindow interface{}
	// Name of the subnet group associated to the cache cluster.
	SubnetGroupName interface{}
	// The tags assigned to the resource
	Tags interface{}
}