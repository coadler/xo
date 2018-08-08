// Package ischema contains the types for schema 'information_schema'.
package ischema

import "github.com/coadler/xo/examples/pgcatalog/pgtypes"

// Code generated by xo. DO NOT EDIT.

// UsagePrivilege represents a row from 'information_schema.usage_privileges'.
type UsagePrivilege struct {
	Grantor       pgtypes.SQLIdentifier `json:"grantor"`        // grantor
	Grantee       pgtypes.SQLIdentifier `json:"grantee"`        // grantee
	ObjectCatalog pgtypes.SQLIdentifier `json:"object_catalog"` // object_catalog
	ObjectSchema  pgtypes.SQLIdentifier `json:"object_schema"`  // object_schema
	ObjectName    pgtypes.SQLIdentifier `json:"object_name"`    // object_name
	ObjectType    pgtypes.CharacterData `json:"object_type"`    // object_type
	PrivilegeType pgtypes.CharacterData `json:"privilege_type"` // privilege_type
	IsGrantable   pgtypes.YesOrNo       `json:"is_grantable"`   // is_grantable
}
