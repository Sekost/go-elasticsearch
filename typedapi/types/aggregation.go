// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741

package types

// Aggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741/specification/_types/aggregations/Aggregation.ts#L22-L25
type Aggregation struct {
	Meta *Metadata `json:"meta,omitempty"`
	Name *string   `json:"name,omitempty"`
}

// AggregationBuilder holds Aggregation struct and provides a builder API.
type AggregationBuilder struct {
	v *Aggregation
}

// NewAggregation provides a builder for the Aggregation struct.
func NewAggregationBuilder() *AggregationBuilder {
	r := AggregationBuilder{
		&Aggregation{},
	}

	return &r
}

// Build finalize the chain and returns the Aggregation struct
func (rb *AggregationBuilder) Build() Aggregation {
	return *rb.v
}

func (rb *AggregationBuilder) Meta(meta *MetadataBuilder) *AggregationBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}

func (rb *AggregationBuilder) Name(name string) *AggregationBuilder {
	rb.v.Name = &name
	return rb
}