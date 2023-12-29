// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// FlowLogSpec defines the desired state of FlowLog.
//
// Describes a flow log.
type FlowLogSpec struct {

	// The ARN for the IAM role that permits Amazon EC2 to publish flow logs to
	// a CloudWatch Logs log group in your account.
	//
	// If you specify LogDestinationType as s3, do not specify DeliverLogsPermissionArn
	// or LogGroupName.
	DeliverLogsPermissionARN *string `json:"deliverLogsPermissionARN,omitempty"`
	// The destination options.
	DestinationOptions *DestinationOptionsRequest `json:"destinationOptions,omitempty"`
	// The destination to which the flow log data is to be published. Flow log data
	// can be published to a CloudWatch Logs log group or an Amazon S3 bucket. The
	// value specified for this parameter depends on the value specified for LogDestinationType.
	//
	// If LogDestinationType is not specified or cloud-watch-logs, specify the Amazon
	// Resource Name (ARN) of the CloudWatch Logs log group. For example, to publish
	// to a log group called my-logs, specify arn:aws:logs:us-east-1:123456789012:log-group:my-logs.
	// Alternatively, use LogGroupName instead.
	//
	// If LogDestinationType is s3, specify the ARN of the Amazon S3 bucket. You
	// can also specify a subfolder in the bucket. To specify a subfolder in the
	// bucket, use the following ARN format: bucket_ARN/subfolder_name/. For example,
	// to specify a subfolder named my-logs in a bucket named my-bucket, use the
	// following ARN: arn:aws:s3:::my-bucket/my-logs/. You cannot use AWSLogs as
	// a subfolder name. This is a reserved term.
	LogDestination *string `json:"logDestination,omitempty"`
	// The type of destination to which the flow log data is to be published. Flow
	// log data can be published to CloudWatch Logs or Amazon S3. To publish flow
	// log data to CloudWatch Logs, specify cloud-watch-logs. To publish flow log
	// data to Amazon S3, specify s3.
	//
	// If you specify LogDestinationType as s3, do not specify DeliverLogsPermissionArn
	// or LogGroupName.
	//
	// Default: cloud-watch-logs
	LogDestinationType *string `json:"logDestinationType,omitempty"`
	// The fields to include in the flow log record, in the order in which they
	// should appear. For a list of available fields, see Flow log records (https://docs.aws.amazon.com/vpc/latest/userguide/flow-logs.html#flow-log-records).
	// If you omit this parameter, the flow log is created using the default format.
	// If you specify this parameter, you must specify at least one field.
	//
	// Specify the fields using the ${field-id} format, separated by spaces. For
	// the CLI, surround this parameter value with single quotes on Linux or double
	// quotes on Windows.
	LogFormat *string `json:"logFormat,omitempty"`
	// The name of a new or existing CloudWatch Logs log group where Amazon EC2
	// publishes your flow logs.
	//
	// If you specify LogDestinationType as s3, do not specify DeliverLogsPermissionArn
	// or LogGroupName.
	LogGroupName *string `json:"logGroupName,omitempty"`
	// The maximum interval of time during which a flow of packets is captured and
	// aggregated into a flow log record. You can specify 60 seconds (1 minute)
	// or 600 seconds (10 minutes).
	//
	// When a network interface is attached to a Nitro-based instance (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html#ec2-nitro-instances),
	// the aggregation interval is always 60 seconds or less, regardless of the
	// value that you specify.
	//
	// Default: 600
	MaxAggregationInterval *int64 `json:"maxAggregationInterval,omitempty"`
	// +kubebuilder:validation:Required
	ResourceID *string `json:"resourceID"`
	// The type of resource for which to create the flow log. For example, if you
	// specified a VPC ID for the ResourceId property, specify VPC for this property.
	// +kubebuilder:validation:Required
	ResourceType *string `json:"resourceType"`
	// The tags. The value parameter is required, but if you don't want the tag
	// to have a value, specify the parameter with no value, and we set the value
	// to an empty string.
	Tags []*Tag `json:"tags,omitempty"`
	// The type of traffic to log. You can log traffic that the resource accepts
	// or rejects, or all traffic.
	TrafficType *string `json:"trafficType,omitempty"`
}

// FlowLogStatus defines the observed state of FlowLog
type FlowLogStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	// +kubebuilder:validation:Optional
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	// +kubebuilder:validation:Optional
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
	// Unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	// +kubebuilder:validation:Optional
	ClientToken *string `json:"clientToken,omitempty"`
	// +kubebuilder:validation:Optional
	FlowLogID *string `json:"flowLogID,omitempty"`
	// Information about the flow logs that could not be created successfully.
	// +kubebuilder:validation:Optional
	Unsuccessful []*UnsuccessfulItem `json:"unsuccessful,omitempty"`
}

// FlowLog is the Schema for the FlowLogs API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type FlowLog struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FlowLogSpec   `json:"spec,omitempty"`
	Status            FlowLogStatus `json:"status,omitempty"`
}

// FlowLogList contains a list of FlowLog
// +kubebuilder:object:root=true
type FlowLogList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FlowLog `json:"items"`
}

func init() {
	SchemeBuilder.Register(&FlowLog{}, &FlowLogList{})
}
