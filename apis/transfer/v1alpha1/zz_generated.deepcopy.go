// +build !ignore_autogenerated

/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomEndpointDetails) DeepCopyInto(out *CustomEndpointDetails) {
	*out = *in
	if in.AddressAllocationIDs != nil {
		in, out := &in.AddressAllocationIDs, &out.AddressAllocationIDs
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SecurityGroupIDs != nil {
		in, out := &in.SecurityGroupIDs, &out.SecurityGroupIDs
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SecurityGroupIDRefs != nil {
		in, out := &in.SecurityGroupIDRefs, &out.SecurityGroupIDRefs
		*out = make([]v1.Reference, len(*in))
		copy(*out, *in)
	}
	if in.SecurityGroupIDSelector != nil {
		in, out := &in.SecurityGroupIDSelector, &out.SecurityGroupIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.SubnetIDs != nil {
		in, out := &in.SubnetIDs, &out.SubnetIDs
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SubnetIDRefs != nil {
		in, out := &in.SubnetIDRefs, &out.SubnetIDRefs
		*out = make([]v1.Reference, len(*in))
		copy(*out, *in)
	}
	if in.SubnetIDSelector != nil {
		in, out := &in.SubnetIDSelector, &out.SubnetIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.VPCEndpointID != nil {
		in, out := &in.VPCEndpointID, &out.VPCEndpointID
		*out = new(string)
		**out = **in
	}
	if in.VPCID != nil {
		in, out := &in.VPCID, &out.VPCID
		*out = new(string)
		**out = **in
	}
	if in.VPCIDRef != nil {
		in, out := &in.VPCIDRef, &out.VPCIDRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.VPCIDSelector != nil {
		in, out := &in.VPCIDSelector, &out.VPCIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomEndpointDetails.
func (in *CustomEndpointDetails) DeepCopy() *CustomEndpointDetails {
	if in == nil {
		return nil
	}
	out := new(CustomEndpointDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomServerParameters) DeepCopyInto(out *CustomServerParameters) {
	*out = *in
	if in.CustomEndpointDetails != nil {
		in, out := &in.CustomEndpointDetails, &out.CustomEndpointDetails
		*out = new(CustomEndpointDetails)
		(*in).DeepCopyInto(*out)
	}
	if in.CertificateRef != nil {
		in, out := &in.CertificateRef, &out.CertificateRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.CertificateSelector != nil {
		in, out := &in.CertificateSelector, &out.CertificateSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.LoggingRoleRef != nil {
		in, out := &in.LoggingRoleRef, &out.LoggingRoleRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.LoggingRoleSelector != nil {
		in, out := &in.LoggingRoleSelector, &out.LoggingRoleSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomServerParameters.
func (in *CustomServerParameters) DeepCopy() *CustomServerParameters {
	if in == nil {
		return nil
	}
	out := new(CustomServerParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomUserParameters) DeepCopyInto(out *CustomUserParameters) {
	*out = *in
	if in.ServerID != nil {
		in, out := &in.ServerID, &out.ServerID
		*out = new(string)
		**out = **in
	}
	if in.ServerIDRef != nil {
		in, out := &in.ServerIDRef, &out.ServerIDRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.ServerIDSelector != nil {
		in, out := &in.ServerIDSelector, &out.ServerIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(string)
		**out = **in
	}
	if in.RoleRef != nil {
		in, out := &in.RoleRef, &out.RoleRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.RoleSelector != nil {
		in, out := &in.RoleSelector, &out.RoleSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomUserParameters.
func (in *CustomUserParameters) DeepCopy() *CustomUserParameters {
	if in == nil {
		return nil
	}
	out := new(CustomUserParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DescribedSecurityPolicy) DeepCopyInto(out *DescribedSecurityPolicy) {
	*out = *in
	if in.SecurityPolicyName != nil {
		in, out := &in.SecurityPolicyName, &out.SecurityPolicyName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DescribedSecurityPolicy.
func (in *DescribedSecurityPolicy) DeepCopy() *DescribedSecurityPolicy {
	if in == nil {
		return nil
	}
	out := new(DescribedSecurityPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DescribedServer) DeepCopyInto(out *DescribedServer) {
	*out = *in
	if in.ARN != nil {
		in, out := &in.ARN, &out.ARN
		*out = new(string)
		**out = **in
	}
	if in.Certificate != nil {
		in, out := &in.Certificate, &out.Certificate
		*out = new(string)
		**out = **in
	}
	if in.Domain != nil {
		in, out := &in.Domain, &out.Domain
		*out = new(string)
		**out = **in
	}
	if in.EndpointDetails != nil {
		in, out := &in.EndpointDetails, &out.EndpointDetails
		*out = new(EndpointDetails)
		(*in).DeepCopyInto(*out)
	}
	if in.EndpointType != nil {
		in, out := &in.EndpointType, &out.EndpointType
		*out = new(string)
		**out = **in
	}
	if in.HostKeyFingerprint != nil {
		in, out := &in.HostKeyFingerprint, &out.HostKeyFingerprint
		*out = new(string)
		**out = **in
	}
	if in.IdentityProviderDetails != nil {
		in, out := &in.IdentityProviderDetails, &out.IdentityProviderDetails
		*out = new(IdentityProviderDetails)
		(*in).DeepCopyInto(*out)
	}
	if in.IdentityProviderType != nil {
		in, out := &in.IdentityProviderType, &out.IdentityProviderType
		*out = new(string)
		**out = **in
	}
	if in.LoggingRole != nil {
		in, out := &in.LoggingRole, &out.LoggingRole
		*out = new(string)
		**out = **in
	}
	if in.Protocols != nil {
		in, out := &in.Protocols, &out.Protocols
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SecurityPolicyName != nil {
		in, out := &in.SecurityPolicyName, &out.SecurityPolicyName
		*out = new(string)
		**out = **in
	}
	if in.ServerID != nil {
		in, out := &in.ServerID, &out.ServerID
		*out = new(string)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*Tag, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Tag)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.UserCount != nil {
		in, out := &in.UserCount, &out.UserCount
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DescribedServer.
func (in *DescribedServer) DeepCopy() *DescribedServer {
	if in == nil {
		return nil
	}
	out := new(DescribedServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DescribedUser) DeepCopyInto(out *DescribedUser) {
	*out = *in
	if in.ARN != nil {
		in, out := &in.ARN, &out.ARN
		*out = new(string)
		**out = **in
	}
	if in.HomeDirectory != nil {
		in, out := &in.HomeDirectory, &out.HomeDirectory
		*out = new(string)
		**out = **in
	}
	if in.HomeDirectoryMappings != nil {
		in, out := &in.HomeDirectoryMappings, &out.HomeDirectoryMappings
		*out = make([]*HomeDirectoryMapEntry, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(HomeDirectoryMapEntry)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.HomeDirectoryType != nil {
		in, out := &in.HomeDirectoryType, &out.HomeDirectoryType
		*out = new(string)
		**out = **in
	}
	if in.Policy != nil {
		in, out := &in.Policy, &out.Policy
		*out = new(string)
		**out = **in
	}
	if in.PosixProfile != nil {
		in, out := &in.PosixProfile, &out.PosixProfile
		*out = new(PosixProfile)
		(*in).DeepCopyInto(*out)
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(string)
		**out = **in
	}
	if in.SshPublicKeys != nil {
		in, out := &in.SshPublicKeys, &out.SshPublicKeys
		*out = make([]*SshPublicKey, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(SshPublicKey)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*Tag, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Tag)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.UserName != nil {
		in, out := &in.UserName, &out.UserName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DescribedUser.
func (in *DescribedUser) DeepCopy() *DescribedUser {
	if in == nil {
		return nil
	}
	out := new(DescribedUser)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointDetails) DeepCopyInto(out *EndpointDetails) {
	*out = *in
	if in.AddressAllocationIDs != nil {
		in, out := &in.AddressAllocationIDs, &out.AddressAllocationIDs
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SecurityGroupIDs != nil {
		in, out := &in.SecurityGroupIDs, &out.SecurityGroupIDs
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SubnetIDs != nil {
		in, out := &in.SubnetIDs, &out.SubnetIDs
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.VPCEndpointID != nil {
		in, out := &in.VPCEndpointID, &out.VPCEndpointID
		*out = new(string)
		**out = **in
	}
	if in.VPCID != nil {
		in, out := &in.VPCID, &out.VPCID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointDetails.
func (in *EndpointDetails) DeepCopy() *EndpointDetails {
	if in == nil {
		return nil
	}
	out := new(EndpointDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HomeDirectoryMapEntry) DeepCopyInto(out *HomeDirectoryMapEntry) {
	*out = *in
	if in.Entry != nil {
		in, out := &in.Entry, &out.Entry
		*out = new(string)
		**out = **in
	}
	if in.Target != nil {
		in, out := &in.Target, &out.Target
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HomeDirectoryMapEntry.
func (in *HomeDirectoryMapEntry) DeepCopy() *HomeDirectoryMapEntry {
	if in == nil {
		return nil
	}
	out := new(HomeDirectoryMapEntry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityProviderDetails) DeepCopyInto(out *IdentityProviderDetails) {
	*out = *in
	if in.InvocationRole != nil {
		in, out := &in.InvocationRole, &out.InvocationRole
		*out = new(string)
		**out = **in
	}
	if in.URL != nil {
		in, out := &in.URL, &out.URL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityProviderDetails.
func (in *IdentityProviderDetails) DeepCopy() *IdentityProviderDetails {
	if in == nil {
		return nil
	}
	out := new(IdentityProviderDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ListedServer) DeepCopyInto(out *ListedServer) {
	*out = *in
	if in.ARN != nil {
		in, out := &in.ARN, &out.ARN
		*out = new(string)
		**out = **in
	}
	if in.Domain != nil {
		in, out := &in.Domain, &out.Domain
		*out = new(string)
		**out = **in
	}
	if in.EndpointType != nil {
		in, out := &in.EndpointType, &out.EndpointType
		*out = new(string)
		**out = **in
	}
	if in.IdentityProviderType != nil {
		in, out := &in.IdentityProviderType, &out.IdentityProviderType
		*out = new(string)
		**out = **in
	}
	if in.LoggingRole != nil {
		in, out := &in.LoggingRole, &out.LoggingRole
		*out = new(string)
		**out = **in
	}
	if in.ServerID != nil {
		in, out := &in.ServerID, &out.ServerID
		*out = new(string)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	if in.UserCount != nil {
		in, out := &in.UserCount, &out.UserCount
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ListedServer.
func (in *ListedServer) DeepCopy() *ListedServer {
	if in == nil {
		return nil
	}
	out := new(ListedServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ListedUser) DeepCopyInto(out *ListedUser) {
	*out = *in
	if in.ARN != nil {
		in, out := &in.ARN, &out.ARN
		*out = new(string)
		**out = **in
	}
	if in.HomeDirectory != nil {
		in, out := &in.HomeDirectory, &out.HomeDirectory
		*out = new(string)
		**out = **in
	}
	if in.HomeDirectoryType != nil {
		in, out := &in.HomeDirectoryType, &out.HomeDirectoryType
		*out = new(string)
		**out = **in
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(string)
		**out = **in
	}
	if in.SshPublicKeyCount != nil {
		in, out := &in.SshPublicKeyCount, &out.SshPublicKeyCount
		*out = new(int64)
		**out = **in
	}
	if in.UserName != nil {
		in, out := &in.UserName, &out.UserName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ListedUser.
func (in *ListedUser) DeepCopy() *ListedUser {
	if in == nil {
		return nil
	}
	out := new(ListedUser)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PosixProfile) DeepCopyInto(out *PosixProfile) {
	*out = *in
	if in.Gid != nil {
		in, out := &in.Gid, &out.Gid
		*out = new(int64)
		**out = **in
	}
	if in.SecondaryGids != nil {
		in, out := &in.SecondaryGids, &out.SecondaryGids
		*out = make([]*int64, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(int64)
				**out = **in
			}
		}
	}
	if in.Uid != nil {
		in, out := &in.Uid, &out.Uid
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PosixProfile.
func (in *PosixProfile) DeepCopy() *PosixProfile {
	if in == nil {
		return nil
	}
	out := new(PosixProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Server) DeepCopyInto(out *Server) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Server.
func (in *Server) DeepCopy() *Server {
	if in == nil {
		return nil
	}
	out := new(Server)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Server) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerList) DeepCopyInto(out *ServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Server, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerList.
func (in *ServerList) DeepCopy() *ServerList {
	if in == nil {
		return nil
	}
	out := new(ServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerObservation) DeepCopyInto(out *ServerObservation) {
	*out = *in
	if in.ServerID != nil {
		in, out := &in.ServerID, &out.ServerID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerObservation.
func (in *ServerObservation) DeepCopy() *ServerObservation {
	if in == nil {
		return nil
	}
	out := new(ServerObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerParameters) DeepCopyInto(out *ServerParameters) {
	*out = *in
	if in.Certificate != nil {
		in, out := &in.Certificate, &out.Certificate
		*out = new(string)
		**out = **in
	}
	if in.Domain != nil {
		in, out := &in.Domain, &out.Domain
		*out = new(string)
		**out = **in
	}
	if in.EndpointType != nil {
		in, out := &in.EndpointType, &out.EndpointType
		*out = new(string)
		**out = **in
	}
	if in.HostKey != nil {
		in, out := &in.HostKey, &out.HostKey
		*out = new(string)
		**out = **in
	}
	if in.IdentityProviderDetails != nil {
		in, out := &in.IdentityProviderDetails, &out.IdentityProviderDetails
		*out = new(IdentityProviderDetails)
		(*in).DeepCopyInto(*out)
	}
	if in.IdentityProviderType != nil {
		in, out := &in.IdentityProviderType, &out.IdentityProviderType
		*out = new(string)
		**out = **in
	}
	if in.LoggingRole != nil {
		in, out := &in.LoggingRole, &out.LoggingRole
		*out = new(string)
		**out = **in
	}
	if in.Protocols != nil {
		in, out := &in.Protocols, &out.Protocols
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SecurityPolicyName != nil {
		in, out := &in.SecurityPolicyName, &out.SecurityPolicyName
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*Tag, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Tag)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	in.CustomServerParameters.DeepCopyInto(&out.CustomServerParameters)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerParameters.
func (in *ServerParameters) DeepCopy() *ServerParameters {
	if in == nil {
		return nil
	}
	out := new(ServerParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerSpec) DeepCopyInto(out *ServerSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerSpec.
func (in *ServerSpec) DeepCopy() *ServerSpec {
	if in == nil {
		return nil
	}
	out := new(ServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerStatus) DeepCopyInto(out *ServerStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerStatus.
func (in *ServerStatus) DeepCopy() *ServerStatus {
	if in == nil {
		return nil
	}
	out := new(ServerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SshPublicKey) DeepCopyInto(out *SshPublicKey) {
	*out = *in
	if in.DateImported != nil {
		in, out := &in.DateImported, &out.DateImported
		*out = (*in).DeepCopy()
	}
	if in.SshPublicKeyBody != nil {
		in, out := &in.SshPublicKeyBody, &out.SshPublicKeyBody
		*out = new(string)
		**out = **in
	}
	if in.SshPublicKeyID != nil {
		in, out := &in.SshPublicKeyID, &out.SshPublicKeyID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SshPublicKey.
func (in *SshPublicKey) DeepCopy() *SshPublicKey {
	if in == nil {
		return nil
	}
	out := new(SshPublicKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Tag) DeepCopyInto(out *Tag) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tag.
func (in *Tag) DeepCopy() *Tag {
	if in == nil {
		return nil
	}
	out := new(Tag)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *User) DeepCopyInto(out *User) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new User.
func (in *User) DeepCopy() *User {
	if in == nil {
		return nil
	}
	out := new(User)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *User) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserList) DeepCopyInto(out *UserList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]User, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserList.
func (in *UserList) DeepCopy() *UserList {
	if in == nil {
		return nil
	}
	out := new(UserList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UserList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserObservation) DeepCopyInto(out *UserObservation) {
	*out = *in
	if in.ServerID != nil {
		in, out := &in.ServerID, &out.ServerID
		*out = new(string)
		**out = **in
	}
	if in.UserName != nil {
		in, out := &in.UserName, &out.UserName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserObservation.
func (in *UserObservation) DeepCopy() *UserObservation {
	if in == nil {
		return nil
	}
	out := new(UserObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserParameters) DeepCopyInto(out *UserParameters) {
	*out = *in
	if in.HomeDirectory != nil {
		in, out := &in.HomeDirectory, &out.HomeDirectory
		*out = new(string)
		**out = **in
	}
	if in.HomeDirectoryMappings != nil {
		in, out := &in.HomeDirectoryMappings, &out.HomeDirectoryMappings
		*out = make([]*HomeDirectoryMapEntry, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(HomeDirectoryMapEntry)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.HomeDirectoryType != nil {
		in, out := &in.HomeDirectoryType, &out.HomeDirectoryType
		*out = new(string)
		**out = **in
	}
	if in.Policy != nil {
		in, out := &in.Policy, &out.Policy
		*out = new(string)
		**out = **in
	}
	if in.PosixProfile != nil {
		in, out := &in.PosixProfile, &out.PosixProfile
		*out = new(PosixProfile)
		(*in).DeepCopyInto(*out)
	}
	if in.SshPublicKeyBody != nil {
		in, out := &in.SshPublicKeyBody, &out.SshPublicKeyBody
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*Tag, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Tag)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	in.CustomUserParameters.DeepCopyInto(&out.CustomUserParameters)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserParameters.
func (in *UserParameters) DeepCopy() *UserParameters {
	if in == nil {
		return nil
	}
	out := new(UserParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserSpec) DeepCopyInto(out *UserSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserSpec.
func (in *UserSpec) DeepCopy() *UserSpec {
	if in == nil {
		return nil
	}
	out := new(UserSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserStatus) DeepCopyInto(out *UserStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserStatus.
func (in *UserStatus) DeepCopy() *UserStatus {
	if in == nil {
		return nil
	}
	out := new(UserStatus)
	in.DeepCopyInto(out)
	return out
}
