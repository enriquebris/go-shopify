package goshopify

import "time"

type Metafield struct {
	ID            int        `json:"id,omitempty"`
	Namespace     string     `json:"namespace,omitempty"`
	Key           string     `json:"key,omitempty"`
	Value         string     `json:"value,omitempty"`
	ValueType     string     `json:"value_type,omitempty"`
	Description   string     `jsnon:"description,omitempty"`
	OwnerID       int        `json:"owner_id,omitempty"`
	CreatedAt     *time.Time `json:"created_at,omitempty"`
	UpdatedAt     *time.Time `json:"updated_at,omitempty"`
	OwnerResource string     `json:"owner_resource,omitempty"`
}

// Represents the results from the metafields.json endpoint
type MetafieldsResource struct {
	Metafields []Metafield `json:"metafields"`
}

// Represents the result from the metafields.json endpoint
type MetafieldResource struct {
	Metafield *Metafield `json:"metafield"`
}
