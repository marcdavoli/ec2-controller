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

package vpc_peering_connection

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

	if ackcompare.HasNilDifference(a.ko.Spec.DryRun, b.ko.Spec.DryRun) {
		delta.Add("Spec.DryRun", a.ko.Spec.DryRun, b.ko.Spec.DryRun)
	} else if a.ko.Spec.DryRun != nil && b.ko.Spec.DryRun != nil {
		if *a.ko.Spec.DryRun != *b.ko.Spec.DryRun {
			delta.Add("Spec.DryRun", a.ko.Spec.DryRun, b.ko.Spec.DryRun)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.PeerOwnerID, b.ko.Spec.PeerOwnerID) {
		delta.Add("Spec.PeerOwnerID", a.ko.Spec.PeerOwnerID, b.ko.Spec.PeerOwnerID)
	} else if a.ko.Spec.PeerOwnerID != nil && b.ko.Spec.PeerOwnerID != nil {
		if *a.ko.Spec.PeerOwnerID != *b.ko.Spec.PeerOwnerID {
			delta.Add("Spec.PeerOwnerID", a.ko.Spec.PeerOwnerID, b.ko.Spec.PeerOwnerID)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.PeerRegion, b.ko.Spec.PeerRegion) {
		delta.Add("Spec.PeerRegion", a.ko.Spec.PeerRegion, b.ko.Spec.PeerRegion)
	} else if a.ko.Spec.PeerRegion != nil && b.ko.Spec.PeerRegion != nil {
		if *a.ko.Spec.PeerRegion != *b.ko.Spec.PeerRegion {
			delta.Add("Spec.PeerRegion", a.ko.Spec.PeerRegion, b.ko.Spec.PeerRegion)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.PeerVPCID, b.ko.Spec.PeerVPCID) {
		delta.Add("Spec.PeerVPCID", a.ko.Spec.PeerVPCID, b.ko.Spec.PeerVPCID)
	} else if a.ko.Spec.PeerVPCID != nil && b.ko.Spec.PeerVPCID != nil {
		if *a.ko.Spec.PeerVPCID != *b.ko.Spec.PeerVPCID {
			delta.Add("Spec.PeerVPCID", a.ko.Spec.PeerVPCID, b.ko.Spec.PeerVPCID)
		}
	}
	if !reflect.DeepEqual(a.ko.Spec.PeerVPCRef, b.ko.Spec.PeerVPCRef) {
		delta.Add("Spec.PeerVPCRef", a.ko.Spec.PeerVPCRef, b.ko.Spec.PeerVPCRef)
	}
	if !reflect.DeepEqual(a.ko.Spec.TagSpecifications, b.ko.Spec.TagSpecifications) {
		delta.Add("Spec.TagSpecifications", a.ko.Spec.TagSpecifications, b.ko.Spec.TagSpecifications)
	}
	if ackcompare.HasNilDifference(a.ko.Spec.VPCID, b.ko.Spec.VPCID) {
		delta.Add("Spec.VPCID", a.ko.Spec.VPCID, b.ko.Spec.VPCID)
	} else if a.ko.Spec.VPCID != nil && b.ko.Spec.VPCID != nil {
		if *a.ko.Spec.VPCID != *b.ko.Spec.VPCID {
			delta.Add("Spec.VPCID", a.ko.Spec.VPCID, b.ko.Spec.VPCID)
		}
	}
	if !reflect.DeepEqual(a.ko.Spec.VPCRef, b.ko.Spec.VPCRef) {
		delta.Add("Spec.VPCRef", a.ko.Spec.VPCRef, b.ko.Spec.VPCRef)
	}

	return delta
}
