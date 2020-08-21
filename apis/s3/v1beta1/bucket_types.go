/*
Copyright 2020 The Crossplane Authors.

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

package v1beta1

import (
	runtimev1alpha1 "github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BucketParameters are parameters for configuring the calls made to AWS Bucket API.
type BucketParameters struct {

	// The canned ACL to apply to the bucket.
	// +kubebuilder:validation:Enum=private;public-read;public-read-write;authenticated-read
	// +optional
	ACL *string `json:"acl,omitempty"`

	// Specifies the Region where the bucket will be created. If you don't specify
	// a Region, the bucket is created in the US East (N. Virginia) Region (us-east-1).
	// +optional
	LocationConstraint *string `json:"locationConstraint,omitempty"`

	// Allows grantee the read, write, read ACP, and write ACP permissions on the
	// bucket.
	// +optional
	GrantFullControl *string `json:"grantFullControl,omitempty"`

	// Allows grantee to list the objects in the bucket.
	// +optional
	GrantRead *string `json:"grantRead,omitempty"`

	// Allows grantee to read the bucket ACL.
	// +optional
	GrantReadACP *string `json:"grantReadAcp,omitempty"`

	// Allows grantee to create, overwrite, and delete any object in the bucket.
	// +optional
	GrantWrite *string `json:"grantWrite,omitempty"`

	// Allows grantee to write the ACL for the applicable bucket.
	// +optional
	GrantWriteACP *string `json:"grantWriteAcp,omitempty"`

	// Specifies whether you want S3 Object Lock to be enabled for the new bucket.
	// +optional
	ObjectLockEnabledForBucket *bool `json:"objectLockEnabledForBucket,omitempty"`
}

// BucketObservation is observation of Bucket properties.
type BucketObservation struct {
}

// BucketSpec represents the desired state of the Bucket.
type BucketSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  BucketParameters `json:"forProvider"`
}

// BucketStatus represents the observed state of the Bucket.
type BucketStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     BucketObservation `json:"atProvider"`
}

// +kubebuilder:object:root=true

// An Bucket is a managed resource that represents an AWS S3 Bucket.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Bucket struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BucketSpec   `json:"spec"`
	Status BucketStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketList contains a list of Buckets
type BucketList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Bucket `json:"items"`
}
