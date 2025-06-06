// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

//go:build go1.23

package compute

import (
	"iter"

	computepb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/googleapis/gax-go/v2/iterator"
)

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *AcceleratorTypeIterator) All() iter.Seq2[*computepb.AcceleratorType, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *AcceleratorTypesScopedListPairIterator) All() iter.Seq2[AcceleratorTypesScopedListPair, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *AddressIterator) All() iter.Seq2[*computepb.Address, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *AddressesScopedListPairIterator) All() iter.Seq2[AddressesScopedListPair, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *AutoscalerIterator) All() iter.Seq2[*computepb.Autoscaler, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *AutoscalersScopedListPairIterator) All() iter.Seq2[AutoscalersScopedListPair, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *BackendBucketIterator) All() iter.Seq2[*computepb.BackendBucket, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *BackendServiceIterator) All() iter.Seq2[*computepb.BackendService, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *BackendServicesScopedListPairIterator) All() iter.Seq2[BackendServicesScopedListPair, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *BgpRouteIterator) All() iter.Seq2[*computepb.BgpRoute, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *CommitmentIterator) All() iter.Seq2[*computepb.Commitment, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *CommitmentsScopedListPairIterator) All() iter.Seq2[CommitmentsScopedListPair, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *DiskIterator) All() iter.Seq2[*computepb.Disk, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *DiskTypeIterator) All() iter.Seq2[*computepb.DiskType, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *DiskTypesScopedListPairIterator) All() iter.Seq2[DiskTypesScopedListPair, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *DisksScopedListPairIterator) All() iter.Seq2[DisksScopedListPair, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *ExchangedPeeringRouteIterator) All() iter.Seq2[*computepb.ExchangedPeeringRoute, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *ExternalVpnGatewayIterator) All() iter.Seq2[*computepb.ExternalVpnGateway, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *FirewallIterator) All() iter.Seq2[*computepb.Firewall, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *FirewallPoliciesScopedListPairIterator) All() iter.Seq2[FirewallPoliciesScopedListPair, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *FirewallPolicyIterator) All() iter.Seq2[*computepb.FirewallPolicy, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *ForwardingRuleIterator) All() iter.Seq2[*computepb.ForwardingRule, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *ForwardingRulesScopedListPairIterator) All() iter.Seq2[ForwardingRulesScopedListPair, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *HealthCheckIterator) All() iter.Seq2[*computepb.HealthCheck, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *HealthCheckServiceIterator) All() iter.Seq2[*computepb.HealthCheckService, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *HealthChecksScopedListPairIterator) All() iter.Seq2[HealthChecksScopedListPair, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *ImageIterator) All() iter.Seq2[*computepb.Image, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *InstanceGroupIterator) All() iter.Seq2[*computepb.InstanceGroup, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *InstanceGroupManagerIterator) All() iter.Seq2[*computepb.InstanceGroupManager, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *InstanceGroupManagerResizeRequestIterator) All() iter.Seq2[*computepb.InstanceGroupManagerResizeRequest, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *InstanceGroupManagersScopedListPairIterator) All() iter.Seq2[InstanceGroupManagersScopedListPair, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *InstanceGroupsScopedListPairIterator) All() iter.Seq2[InstanceGroupsScopedListPair, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *InstanceIterator) All() iter.Seq2[*computepb.Instance, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *InstanceManagedByIgmErrorIterator) All() iter.Seq2[*computepb.InstanceManagedByIgmError, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *InstanceTemplateIterator) All() iter.Seq2[*computepb.InstanceTemplate, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *InstanceTemplatesScopedListPairIterator) All() iter.Seq2[InstanceTemplatesScopedListPair, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *InstanceWithNamedPortsIterator) All() iter.Seq2[*computepb.InstanceWithNamedPorts, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *InstancesScopedListPairIterator) All() iter.Seq2[InstancesScopedListPair, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *InstantSnapshotIterator) All() iter.Seq2[*computepb.InstantSnapshot, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *InstantSnapshotsScopedListPairIterator) All() iter.Seq2[InstantSnapshotsScopedListPair, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *InterconnectAttachmentGroupIterator) All() iter.Seq2[*computepb.InterconnectAttachmentGroup, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *InterconnectAttachmentIterator) All() iter.Seq2[*computepb.InterconnectAttachment, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *InterconnectAttachmentsScopedListPairIterator) All() iter.Seq2[InterconnectAttachmentsScopedListPair, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *InterconnectGroupIterator) All() iter.Seq2[*computepb.InterconnectGroup, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *InterconnectIterator) All() iter.Seq2[*computepb.Interconnect, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *InterconnectLocationIterator) All() iter.Seq2[*computepb.InterconnectLocation, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *InterconnectRemoteLocationIterator) All() iter.Seq2[*computepb.InterconnectRemoteLocation, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *LicenseIterator) All() iter.Seq2[*computepb.License, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *MachineImageIterator) All() iter.Seq2[*computepb.MachineImage, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *MachineTypeIterator) All() iter.Seq2[*computepb.MachineType, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *MachineTypesScopedListPairIterator) All() iter.Seq2[MachineTypesScopedListPair, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *ManagedInstanceIterator) All() iter.Seq2[*computepb.ManagedInstance, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *NetworkAttachmentIterator) All() iter.Seq2[*computepb.NetworkAttachment, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *NetworkAttachmentsScopedListPairIterator) All() iter.Seq2[NetworkAttachmentsScopedListPair, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *NetworkEdgeSecurityServicesScopedListPairIterator) All() iter.Seq2[NetworkEdgeSecurityServicesScopedListPair, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *NetworkEndpointGroupIterator) All() iter.Seq2[*computepb.NetworkEndpointGroup, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *NetworkEndpointGroupsScopedListPairIterator) All() iter.Seq2[NetworkEndpointGroupsScopedListPair, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *NetworkEndpointWithHealthStatusIterator) All() iter.Seq2[*computepb.NetworkEndpointWithHealthStatus, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *NetworkIterator) All() iter.Seq2[*computepb.Network, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *NetworkProfileIterator) All() iter.Seq2[*computepb.NetworkProfile, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *NodeGroupIterator) All() iter.Seq2[*computepb.NodeGroup, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *NodeGroupNodeIterator) All() iter.Seq2[*computepb.NodeGroupNode, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *NodeGroupsScopedListPairIterator) All() iter.Seq2[NodeGroupsScopedListPair, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *NodeTemplateIterator) All() iter.Seq2[*computepb.NodeTemplate, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *NodeTemplatesScopedListPairIterator) All() iter.Seq2[NodeTemplatesScopedListPair, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *NodeTypeIterator) All() iter.Seq2[*computepb.NodeType, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *NodeTypesScopedListPairIterator) All() iter.Seq2[NodeTypesScopedListPair, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *NotificationEndpointIterator) All() iter.Seq2[*computepb.NotificationEndpoint, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *OperationIterator) All() iter.Seq2[*computepb.Operation, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *OperationsScopedListPairIterator) All() iter.Seq2[OperationsScopedListPair, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *PacketMirroringIterator) All() iter.Seq2[*computepb.PacketMirroring, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *PacketMirroringsScopedListPairIterator) All() iter.Seq2[PacketMirroringsScopedListPair, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *PerInstanceConfigIterator) All() iter.Seq2[*computepb.PerInstanceConfig, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *ProjectIterator) All() iter.Seq2[*computepb.Project, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *PublicAdvertisedPrefixIterator) All() iter.Seq2[*computepb.PublicAdvertisedPrefix, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *PublicDelegatedPrefixIterator) All() iter.Seq2[*computepb.PublicDelegatedPrefix, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *PublicDelegatedPrefixesScopedListPairIterator) All() iter.Seq2[PublicDelegatedPrefixesScopedListPair, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *ReferenceIterator) All() iter.Seq2[*computepb.Reference, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *RegionIterator) All() iter.Seq2[*computepb.Region, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *ReservationBlockIterator) All() iter.Seq2[*computepb.ReservationBlock, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *ReservationIterator) All() iter.Seq2[*computepb.Reservation, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *ReservationsScopedListPairIterator) All() iter.Seq2[ReservationsScopedListPair, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *ResourcePoliciesScopedListPairIterator) All() iter.Seq2[ResourcePoliciesScopedListPair, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *ResourcePolicyIterator) All() iter.Seq2[*computepb.ResourcePolicy, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *RouteIterator) All() iter.Seq2[*computepb.Route, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *RoutePolicyIterator) All() iter.Seq2[*computepb.RoutePolicy, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *RouterIterator) All() iter.Seq2[*computepb.Router, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *RoutersScopedListPairIterator) All() iter.Seq2[RoutersScopedListPair, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *SecurityPoliciesScopedListPairIterator) All() iter.Seq2[SecurityPoliciesScopedListPair, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *SecurityPolicyIterator) All() iter.Seq2[*computepb.SecurityPolicy, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *ServiceAttachmentIterator) All() iter.Seq2[*computepb.ServiceAttachment, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *ServiceAttachmentsScopedListPairIterator) All() iter.Seq2[ServiceAttachmentsScopedListPair, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *SnapshotIterator) All() iter.Seq2[*computepb.Snapshot, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *SslCertificateIterator) All() iter.Seq2[*computepb.SslCertificate, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *SslCertificatesScopedListPairIterator) All() iter.Seq2[SslCertificatesScopedListPair, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *SslPoliciesScopedListPairIterator) All() iter.Seq2[SslPoliciesScopedListPair, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *SslPolicyIterator) All() iter.Seq2[*computepb.SslPolicy, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *StoragePoolDiskIterator) All() iter.Seq2[*computepb.StoragePoolDisk, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *StoragePoolIterator) All() iter.Seq2[*computepb.StoragePool, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *StoragePoolTypeIterator) All() iter.Seq2[*computepb.StoragePoolType, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *StoragePoolTypesScopedListPairIterator) All() iter.Seq2[StoragePoolTypesScopedListPair, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *StoragePoolsScopedListPairIterator) All() iter.Seq2[StoragePoolsScopedListPair, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *SubnetworkIterator) All() iter.Seq2[*computepb.Subnetwork, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *SubnetworksScopedListPairIterator) All() iter.Seq2[SubnetworksScopedListPair, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *TargetGrpcProxyIterator) All() iter.Seq2[*computepb.TargetGrpcProxy, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *TargetHttpProxiesScopedListPairIterator) All() iter.Seq2[TargetHttpProxiesScopedListPair, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *TargetHttpProxyIterator) All() iter.Seq2[*computepb.TargetHttpProxy, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *TargetHttpsProxiesScopedListPairIterator) All() iter.Seq2[TargetHttpsProxiesScopedListPair, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *TargetHttpsProxyIterator) All() iter.Seq2[*computepb.TargetHttpsProxy, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *TargetInstanceIterator) All() iter.Seq2[*computepb.TargetInstance, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *TargetInstancesScopedListPairIterator) All() iter.Seq2[TargetInstancesScopedListPair, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *TargetPoolIterator) All() iter.Seq2[*computepb.TargetPool, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *TargetPoolsScopedListPairIterator) All() iter.Seq2[TargetPoolsScopedListPair, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *TargetSslProxyIterator) All() iter.Seq2[*computepb.TargetSslProxy, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *TargetTcpProxiesScopedListPairIterator) All() iter.Seq2[TargetTcpProxiesScopedListPair, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *TargetTcpProxyIterator) All() iter.Seq2[*computepb.TargetTcpProxy, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *TargetVpnGatewayIterator) All() iter.Seq2[*computepb.TargetVpnGateway, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *TargetVpnGatewaysScopedListPairIterator) All() iter.Seq2[TargetVpnGatewaysScopedListPair, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *UrlMapIterator) All() iter.Seq2[*computepb.UrlMap, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *UrlMapsScopedListPairIterator) All() iter.Seq2[UrlMapsScopedListPair, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *UsableSubnetworkIterator) All() iter.Seq2[*computepb.UsableSubnetwork, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *VmEndpointNatMappingsIterator) All() iter.Seq2[*computepb.VmEndpointNatMappings, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *VpnGatewayIterator) All() iter.Seq2[*computepb.VpnGateway, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *VpnGatewaysScopedListPairIterator) All() iter.Seq2[VpnGatewaysScopedListPair, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *VpnTunnelIterator) All() iter.Seq2[*computepb.VpnTunnel, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *VpnTunnelsScopedListPairIterator) All() iter.Seq2[VpnTunnelsScopedListPair, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *XpnResourceIdIterator) All() iter.Seq2[*computepb.XpnResourceId, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *ZoneIterator) All() iter.Seq2[*computepb.Zone, error] {
	return iterator.RangeAdapter(it.Next)
}
