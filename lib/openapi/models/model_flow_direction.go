/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type FlowDirection string

// List of FlowDirection
const (
	FlowDirection_DOWNLINK      FlowDirection = "DOWNLINK"
	FlowDirection_UPLINK        FlowDirection = "UPLINK"
	FlowDirection_BIDIRECTIONAL FlowDirection = "BIDIRECTIONAL"
	FlowDirection_UNSPECIFIED   FlowDirection = "UNSPECIFIED"
)