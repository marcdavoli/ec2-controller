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

package vpc_endpoint_service_configuration

import (
	"bytes"
	"reflect"

	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	acktags "github.com/aws-controllers-k8s/runtime/pkg/tags"
)

// Hack to avoid import errors during build...
var (
	_ = &bytes.Buffer{}
	_ = &reflect.Method{}
	_ = &acktags.Tags{}
)

// newResourceDelta returns a new `ackcompare.Delta` used to compare two
// resources
func newResourceDelta(
	a *resource,
	b *resource,
) *ackcompare.Delta {
	delta := ackcompare.NewDelta()
	if (a == nil && b != nil) ||
		(a != nil && b == nil) {
		delta.Add("", a, b)
		return delta
	}
	compareTags(delta, a, b)

	if ackcompare.HasNilDifference(a.ko.Spec.AcceptanceRequired, b.ko.Spec.AcceptanceRequired) {
		delta.Add("Spec.AcceptanceRequired", a.ko.Spec.AcceptanceRequired, b.ko.Spec.AcceptanceRequired)
	} else if a.ko.Spec.AcceptanceRequired != nil && b.ko.Spec.AcceptanceRequired != nil {
		if *a.ko.Spec.AcceptanceRequired != *b.ko.Spec.AcceptanceRequired {
			delta.Add("Spec.AcceptanceRequired", a.ko.Spec.AcceptanceRequired, b.ko.Spec.AcceptanceRequired)
		}
	}
	if !ackcompare.SliceStringPEqual(a.ko.Spec.AllowPrincipals, b.ko.Spec.AllowPrincipals) {
		delta.Add("Spec.AllowPrincipals", a.ko.Spec.AllowPrincipals, b.ko.Spec.AllowPrincipals)
	}
	if !ackcompare.SliceStringPEqual(a.ko.Spec.GatewayLoadBalancerARNs, b.ko.Spec.GatewayLoadBalancerARNs) {
		delta.Add("Spec.GatewayLoadBalancerARNs", a.ko.Spec.GatewayLoadBalancerARNs, b.ko.Spec.GatewayLoadBalancerARNs)
	}
	if !ackcompare.SliceStringPEqual(a.ko.Spec.NetworkLoadBalancerARNs, b.ko.Spec.NetworkLoadBalancerARNs) {
		delta.Add("Spec.NetworkLoadBalancerARNs", a.ko.Spec.NetworkLoadBalancerARNs, b.ko.Spec.NetworkLoadBalancerARNs)
	}
	if ackcompare.HasNilDifference(a.ko.Spec.PrivateDNSName, b.ko.Spec.PrivateDNSName) {
		delta.Add("Spec.PrivateDNSName", a.ko.Spec.PrivateDNSName, b.ko.Spec.PrivateDNSName)
	} else if a.ko.Spec.PrivateDNSName != nil && b.ko.Spec.PrivateDNSName != nil {
		if *a.ko.Spec.PrivateDNSName != *b.ko.Spec.PrivateDNSName {
			delta.Add("Spec.PrivateDNSName", a.ko.Spec.PrivateDNSName, b.ko.Spec.PrivateDNSName)
		}
	}
	if !ackcompare.SliceStringPEqual(a.ko.Spec.SupportedIPAddressTypes, b.ko.Spec.SupportedIPAddressTypes) {
		delta.Add("Spec.SupportedIPAddressTypes", a.ko.Spec.SupportedIPAddressTypes, b.ko.Spec.SupportedIPAddressTypes)
	}

	return delta
}
