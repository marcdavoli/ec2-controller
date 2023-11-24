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

// VpcSpec defines the desired state of Vpc.
//
// Describes a VPC.
type VPCSpec struct {
	AcceptVPCPeeringRequestsFromVPCID   []*string                                  `json:"acceptVPCPeeringRequestsFromVPCID,omitempty"`
	AcceptVPCPeeringRequestsFromVPCRefs []*ackv1alpha1.AWSResourceReferenceWrapper `json:"acceptVPCPeeringRequestsFromVPCRefs,omitempty"`
	// Requests an Amazon-provided IPv6 CIDR block with a /56 prefix length for
	// the VPC. You cannot specify the range of IP addresses, or the size of the
	// CIDR block.
	AmazonProvidedIPv6CIDRBlock *bool `json:"amazonProvidedIPv6CIDRBlock,omitempty"`
	// +kubebuilder:validation:Required
	CIDRBlocks []*string `json:"cidrBlocks"`
	// The attribute value. The valid values are true or false.
	EnableDNSHostnames *bool `json:"enableDNSHostnames,omitempty"`
	// The attribute value. The valid values are true or false.
	EnableDNSSupport *bool `json:"enableDNSSupport,omitempty"`
	// The tenancy options for instances launched into the VPC. For default, instances
	// are launched with shared tenancy by default. You can launch instances with
	// any tenancy into a shared tenancy VPC. For dedicated, instances are launched
	// as dedicated tenancy instances by default. You can only launch instances
	// with a tenancy of dedicated or host into a dedicated tenancy VPC.
	//
	// Important: The host value cannot be used with this parameter. Use the default
	// or dedicated values only.
	//
	// Default: default
	InstanceTenancy *string `json:"instanceTenancy,omitempty"`
	// The ID of an IPv4 IPAM pool you want to use for allocating this VPC's CIDR.
	// For more information, see What is IPAM? (https://docs.aws.amazon.com/vpc/latest/ipam/what-is-it-ipam.html)
	// in the Amazon VPC IPAM User Guide.
	IPv4IPAMPoolID *string `json:"ipv4IPAMPoolID,omitempty"`
	// The netmask length of the IPv4 CIDR you want to allocate to this VPC from
	// an Amazon VPC IP Address Manager (IPAM) pool. For more information about
	// IPAM, see What is IPAM? (https://docs.aws.amazon.com/vpc/latest/ipam/what-is-it-ipam.html)
	// in the Amazon VPC IPAM User Guide.
	IPv4NetmaskLength *int64 `json:"ipv4NetmaskLength,omitempty"`
	// The IPv6 CIDR block from the IPv6 address pool. You must also specify Ipv6Pool
	// in the request.
	//
	// To let Amazon choose the IPv6 CIDR block for you, omit this parameter.
	IPv6CIDRBlock *string `json:"ipv6CIDRBlock,omitempty"`
	// The name of the location from which we advertise the IPV6 CIDR block. Use
	// this parameter to limit the address to this location.
	//
	// You must set AmazonProvidedIpv6CidrBlock to true to use this parameter.
	IPv6CIDRBlockNetworkBorderGroup *string `json:"ipv6CIDRBlockNetworkBorderGroup,omitempty"`
	// The ID of an IPv6 IPAM pool which will be used to allocate this VPC an IPv6
	// CIDR. IPAM is a VPC feature that you can use to automate your IP address
	// management workflows including assigning, tracking, troubleshooting, and
	// auditing IP addresses across Amazon Web Services Regions and accounts throughout
	// your Amazon Web Services Organization. For more information, see What is
	// IPAM? (https://docs.aws.amazon.com/vpc/latest/ipam/what-is-it-ipam.html)
	// in the Amazon VPC IPAM User Guide.
	IPv6IPAMPoolID *string `json:"ipv6IPAMPoolID,omitempty"`
	// The netmask length of the IPv6 CIDR you want to allocate to this VPC from
	// an Amazon VPC IP Address Manager (IPAM) pool. For more information about
	// IPAM, see What is IPAM? (https://docs.aws.amazon.com/vpc/latest/ipam/what-is-it-ipam.html)
	// in the Amazon VPC IPAM User Guide.
	IPv6NetmaskLength *int64 `json:"ipv6NetmaskLength,omitempty"`
	// The ID of an IPv6 address pool from which to allocate the IPv6 CIDR block.
	IPv6Pool *string `json:"ipv6Pool,omitempty"`
	// The tags. The value parameter is required, but if you don't want the tag
	// to have a value, specify the parameter with no value, and we set the value
	// to an empty string.
	Tags []*Tag `json:"tags,omitempty"`
}

// VPCStatus defines the observed state of VPC
type VPCStatus struct {
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
	// Information about the IPv4 CIDR blocks associated with the VPC.
	// +kubebuilder:validation:Optional
	CIDRBlockAssociationSet []*VPCCIDRBlockAssociation `json:"cidrBlockAssociationSet,omitempty"`
	// The ID of the set of DHCP options you've associated with the VPC.
	// +kubebuilder:validation:Optional
	DHCPOptionsID *string `json:"dhcpOptionsID,omitempty"`
	// Information about the IPv6 CIDR blocks associated with the VPC.
	// +kubebuilder:validation:Optional
	IPv6CIDRBlockAssociationSet []*VPCIPv6CIDRBlockAssociation `json:"ipv6CIDRBlockAssociationSet,omitempty"`
	// Indicates whether the VPC is the default VPC.
	// +kubebuilder:validation:Optional
	IsDefault *bool `json:"isDefault,omitempty"`
	// The ID of the Amazon Web Services account that owns the VPC.
	// +kubebuilder:validation:Optional
	OwnerID *string `json:"ownerID,omitempty"`
	// The current state of the VPC.
	// +kubebuilder:validation:Optional
	State *string `json:"state,omitempty"`
	// The ID of the VPC.
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcID,omitempty"`
}

// VPC is the Schema for the VPCS API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="ID",type=string,priority=0,JSONPath=`.status.vpcID`
// +kubebuilder:printcolumn:name="state",type=string,priority=0,JSONPath=`.status.state`
type VPC struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VPCSpec   `json:"spec,omitempty"`
	Status            VPCStatus `json:"status,omitempty"`
}

// VPCList contains a list of VPC
// +kubebuilder:object:root=true
type VPCList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPC `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VPC{}, &VPCList{})
}
