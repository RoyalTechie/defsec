package redshift

import "github.com/aquasecurity/defsec/types"

type Redshift struct {
	Clusters       []Cluster
	SecurityGroups []SecurityGroup
}

type SecurityGroup struct {
	types.Metadata
	Description types.StringValue
}

type Cluster struct {
	Encryption      Encryption
	SubnetGroupName types.StringValue
}

type Encryption struct {
	Enabled  types.BoolValue
	KMSKeyID types.StringValue
}

func (g *SecurityGroup) GetMetadata() *types.Metadata {
	return &g.Metadata
}

func (g *SecurityGroup) GetRawValue() interface{} {
	return nil
}
