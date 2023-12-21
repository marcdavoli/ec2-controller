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

// VpcEndpointServiceConfigurationSpec defines the desired state of VpcEndpointServiceConfiguration.
type VPCEndpointServiceConfigurationSpec struct {

	// Indicates whether requests from service consumers to create an endpoint to
	// your service must be accepted manually.
	AcceptanceRequired *bool `json:"acceptanceRequired,omitempty"`
	// The Amazon Resource Names (ARN) of one or more principals. Permissions are
	// granted to the principals in this list. To grant permissions to all principals,
	// specify an asterisk (*).
	AllowPrincipals []*string `json:"allowPrincipals,omitempty"`
	// The Amazon Resource Names (ARNs) of one or more Gateway Load Balancers.
	GatewayLoadBalancerARNs []*string `json:"gatewayLoadBalancerARNs,omitempty"`
	// The Amazon Resource Names (ARNs) of one or more Network Load Balancers for
	// your service.
	NetworkLoadBalancerARNs []*string `json:"networkLoadBalancerARNs,omitempty"`
	// (Interface endpoint configuration) The private DNS name to assign to the
	// VPC endpoint service.
	PrivateDNSName *string `json:"privateDNSName,omitempty"`
	// The supported IP address types. The possible values are ipv4 and ipv6.
	SupportedIPAddressTypes []*string `json:"supportedIPAddressTypes,omitempty"`
	// The tags. The value parameter is required, but if you don't want the tag
	// to have a value, specify the parameter with no value, and we set the value
	// to an empty string.
	Tags                 []*Tag `json:"tags,omitempty"`
	VerifyPrivateDNSName *bool  `json:"verifyPrivateDNSName,omitempty"`
}

// VPCEndpointServiceConfigurationStatus defines the observed state of VPCEndpointServiceConfiguration
type VPCEndpointServiceConfigurationStatus struct {
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
	// The Availability Zones in which the service is available.
	// +kubebuilder:validation:Optional
	AvailabilityZones []*string `json:"availabilityZones,omitempty"`
	// The DNS names for the service.
	// +kubebuilder:validation:Optional
	BaseEndpointDNSNames []*string `json:"baseEndpointDNSNames,omitempty"`
	// Indicates whether the service manages its VPC endpoints. Management of the
	// service VPC endpoints using the VPC endpoint API is restricted.
	// +kubebuilder:validation:Optional
	ManagesVPCEndpoints *bool `json:"managesVPCEndpoints,omitempty"`
	// The payer responsibility.
	// +kubebuilder:validation:Optional
	PayerResponsibility *string `json:"payerResponsibility,omitempty"`
	// Information about the endpoint service private DNS name configuration.
	// +kubebuilder:validation:Optional
	PrivateDNSNameConfiguration *PrivateDNSNameConfiguration `json:"privateDNSNameConfiguration,omitempty"`
	// The ID of the service.
	// +kubebuilder:validation:Optional
	ServiceID *string `json:"serviceID,omitempty"`
	// The name of the service.
	// +kubebuilder:validation:Optional
	ServiceName *string `json:"serviceName,omitempty"`
	// The service state.
	// +kubebuilder:validation:Optional
	ServiceState *string `json:"serviceState,omitempty"`
	// The type of service.
	// +kubebuilder:validation:Optional
	ServiceType []*ServiceTypeDetail `json:"serviceType,omitempty"`
}

// VPCEndpointServiceConfiguration is the Schema for the VPCEndpointServiceConfigurations API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="ServiceID",type=string,priority=0,JSONPath=`.status.serviceID`
// +kubebuilder:printcolumn:name="ServiceState",type=string,priority=0,JSONPath=`.status.serviceState`
type VPCEndpointServiceConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VPCEndpointServiceConfigurationSpec   `json:"spec,omitempty"`
	Status            VPCEndpointServiceConfigurationStatus `json:"status,omitempty"`
}

// VPCEndpointServiceConfigurationList contains a list of VPCEndpointServiceConfiguration
// +kubebuilder:object:root=true
type VPCEndpointServiceConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPCEndpointServiceConfiguration `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VPCEndpointServiceConfiguration{}, &VPCEndpointServiceConfigurationList{})
}
