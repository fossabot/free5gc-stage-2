/*
 * Npcf_PolicyAuthorization Service API
 *
 * This is the Policy Authorization Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// this data type is defined in the same way as the EventsSubscReqData data type, but with the OpenAPI nullable property set to true.
type EventsSubscReqDataRm struct {
	Events   []AfEventSubscription `json:"events" bson:"events"`
	NotifUri string                `json:"notifUri,omitempty" bson:"notifUri"`
	UsgThres *UsageThresholdRm     `json:"usgThres,omitempty" bson:"usgThres"`
}