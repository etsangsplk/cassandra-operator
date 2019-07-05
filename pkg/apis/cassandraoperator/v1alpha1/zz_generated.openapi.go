// +build !ignore_autogenerated

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/instaclustr/cassandra-operator/pkg/apis/cassandraoperator/v1alpha1.CassandraBackup":           schema_pkg_apis_cassandraoperator_v1alpha1_CassandraBackup(ref),
		"github.com/instaclustr/cassandra-operator/pkg/apis/cassandraoperator/v1alpha1.CassandraBackupSpec":       schema_pkg_apis_cassandraoperator_v1alpha1_CassandraBackupSpec(ref),
		"github.com/instaclustr/cassandra-operator/pkg/apis/cassandraoperator/v1alpha1.CassandraBackupStatus":     schema_pkg_apis_cassandraoperator_v1alpha1_CassandraBackupStatus(ref),
		"github.com/instaclustr/cassandra-operator/pkg/apis/cassandraoperator/v1alpha1.CassandraCluster":          schema_pkg_apis_cassandraoperator_v1alpha1_CassandraCluster(ref),
		"github.com/instaclustr/cassandra-operator/pkg/apis/cassandraoperator/v1alpha1.CassandraClusterSpec":      schema_pkg_apis_cassandraoperator_v1alpha1_CassandraClusterSpec(ref),
		"github.com/instaclustr/cassandra-operator/pkg/apis/cassandraoperator/v1alpha1.CassandraClusterStatus":    schema_pkg_apis_cassandraoperator_v1alpha1_CassandraClusterStatus(ref),
		"github.com/instaclustr/cassandra-operator/pkg/apis/cassandraoperator/v1alpha1.CassandraDataCenter":       schema_pkg_apis_cassandraoperator_v1alpha1_CassandraDataCenter(ref),
		"github.com/instaclustr/cassandra-operator/pkg/apis/cassandraoperator/v1alpha1.CassandraDataCenterSpec":   schema_pkg_apis_cassandraoperator_v1alpha1_CassandraDataCenterSpec(ref),
		"github.com/instaclustr/cassandra-operator/pkg/apis/cassandraoperator/v1alpha1.CassandraDataCenterStatus": schema_pkg_apis_cassandraoperator_v1alpha1_CassandraDataCenterStatus(ref),
	}
}

func schema_pkg_apis_cassandraoperator_v1alpha1_CassandraBackup(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CassandraBackup is the Schema for the cassandrabackups API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/instaclustr/cassandra-operator/pkg/apis/cassandraoperator/v1alpha1.CassandraBackupSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/instaclustr/cassandra-operator/pkg/apis/cassandraoperator/v1alpha1.CassandraBackupStatus"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/instaclustr/cassandra-operator/pkg/apis/cassandraoperator/v1alpha1.CassandraBackupSpec", "github.com/instaclustr/cassandra-operator/pkg/apis/cassandraoperator/v1alpha1.CassandraBackupStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_cassandraoperator_v1alpha1_CassandraBackupSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CassandraBackupSpec defines the desired state of CassandraBackup",
				Properties: map[string]spec.Schema{
					"cdc": {
						SchemaProps: spec.SchemaProps{
							Description: "Cassandra DC name to back up. Used to find the pods in the CDC",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"destinationUri": {
						SchemaProps: spec.SchemaProps{
							Description: "The uri for the backup target location e.g. s3 bucket, filepath",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"keyspaces": {
						SchemaProps: spec.SchemaProps{
							Description: "The list of keyspaces to back up",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"snapshotName": {
						SchemaProps: spec.SchemaProps{
							Description: "The snapshot name for the backup",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"cdc", "destinationUri", "keyspaces", "snapshotName"},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_cassandraoperator_v1alpha1_CassandraBackupStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CassandraBackupStatus defines the observed state of CassandraBackup",
				Properties: map[string]spec.Schema{
					"state": {
						SchemaProps: spec.SchemaProps{
							Description: "State shows the status of the operation",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"progress": {
						SchemaProps: spec.SchemaProps{
							Description: "Progress shows the percentage of the operation done",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"state", "progress"},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_cassandraoperator_v1alpha1_CassandraCluster(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CassandraCluster is the Schema for the cassandraclusters API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/instaclustr/cassandra-operator/pkg/apis/cassandraoperator/v1alpha1.CassandraClusterSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/instaclustr/cassandra-operator/pkg/apis/cassandraoperator/v1alpha1.CassandraClusterStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/instaclustr/cassandra-operator/pkg/apis/cassandraoperator/v1alpha1.CassandraClusterSpec", "github.com/instaclustr/cassandra-operator/pkg/apis/cassandraoperator/v1alpha1.CassandraClusterStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_cassandraoperator_v1alpha1_CassandraClusterSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CassandraClusterSpec defines the desired state of CassandraCluster",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_cassandraoperator_v1alpha1_CassandraClusterStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CassandraClusterStatus defines the observed state of CassandraCluster",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_cassandraoperator_v1alpha1_CassandraDataCenter(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CassandraDataCenter is the Schema for the cassandradatacenters API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/instaclustr/cassandra-operator/pkg/apis/cassandraoperator/v1alpha1.CassandraDataCenterSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/instaclustr/cassandra-operator/pkg/apis/cassandraoperator/v1alpha1.CassandraDataCenterStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/instaclustr/cassandra-operator/pkg/apis/cassandraoperator/v1alpha1.CassandraDataCenterSpec", "github.com/instaclustr/cassandra-operator/pkg/apis/cassandraoperator/v1alpha1.CassandraDataCenterStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_cassandraoperator_v1alpha1_CassandraDataCenterSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CassandraDataCenterSpec defines the desired state of CassandraDataCenter",
				Properties: map[string]spec.Schema{
					"cluster": {
						SchemaProps: spec.SchemaProps{
							Description: "Cluster is either a string or v1.LocalObjectReference Cluster interface{} `json:\"cluster,omitempty\"`",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"nodes": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"cassandraImage": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"sidecarImage": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"imagePullPolicy": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"imagePullSecrets": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/api/core/v1.LocalObjectReference"),
									},
								},
							},
						},
					},
					"resources": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/api/core/v1.ResourceRequirements"),
						},
					},
					"dataVolumeClaimSpec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/api/core/v1.PersistentVolumeClaimSpec"),
						},
					},
					"prometheusSupport": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
				},
				Required: []string{"nodes", "cassandraImage", "sidecarImage", "imagePullPolicy", "resources", "dataVolumeClaimSpec", "prometheusSupport"},
			},
		},
		Dependencies: []string{
			"k8s.io/api/core/v1.LocalObjectReference", "k8s.io/api/core/v1.PersistentVolumeClaimSpec", "k8s.io/api/core/v1.ResourceRequirements"},
	}
}

func schema_pkg_apis_cassandraoperator_v1alpha1_CassandraDataCenterStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CassandraDataCenterStatus defines the observed state of CassandraDataCenter",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}