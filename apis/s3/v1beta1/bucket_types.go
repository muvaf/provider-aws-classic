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

	// Specifies default encryption for a bucket using server-side encryption with
	// Amazon S3-managed keys (SSE-S3) or customer master keys stored in AWS KMS
	// (SSE-KMS). For information about the Amazon S3 default encryption feature,
	// see Amazon S3 Default Bucket Encryption (https://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-encryption.html)
	// in the Amazon Simple Storage Service Developer Guide.
	// +optional
	ServerSideEncryptionConfiguration *ServerSideEncryptionConfiguration `json:"serverSideEncryptionConfiguration,omitempty"`
}

// ServerSideEncryptionConfiguration specifies the default server-side-encryption configuration.
type ServerSideEncryptionConfiguration struct {
	// Container for information about a particular server-side encryption configuration
	// rule.
	Rules []ServerSideEncryptionRule `json:"rules"`
}

// ServerSideEncryptionRule Specifies the default server-side encryption configuration.
type ServerSideEncryptionRule struct {
	// Specifies the default server-side encryption to apply to new objects in the
	// bucket. If a PUT Object request doesn't specify any server-side encryption,
	// this default encryption will be applied.
	ApplyServerSideEncryptionByDefault ServerSideEncryptionByDefault `json:"applyServerSideEncryptionByDefault"`
}

// ServerSideEncryptionByDefault describes the default server-side encryption to
// apply to new objects in the bucket. If a PUT Object request doesn't specify
// any server-side encryption, this default encryption will be applied.
type ServerSideEncryptionByDefault struct {

	// AWS Key Management Service (KMS) customer master key ID to use for the default
	// encryption. This parameter is allowed if and only if SSEAlgorithm is set
	// to aws:kms.
	//
	// You can specify the key ID or the Amazon Resource Name (ARN) of the CMK.
	// However, if you are using encryption with cross-account operations, you must
	// use a fully qualified CMK ARN. For more information, see Using encryption
	// for cross-account operations (https://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-encryption.html#bucket-encryption-update-bucket-policy).
	//
	// For example:
	//
	//    * Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab
	//
	//    * Key ARN: arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//
	// Amazon S3 only supports symmetric CMKs and not asymmetric CMKs. For more
	// information, see Using Symmetric and Asymmetric Keys (https://docs.aws.amazon.com/kms/latest/developerguide/symmetric-asymmetric.html)
	// in the AWS Key Management Service Developer Guide.
	// +optional
	KMSMasterKeyID *string `json:"kmsMasterKeyId,omitempty"`

	// NOTE(muvaf): aws:kms is not accepted by kubebuilder enum.

	// Server-side encryption algorithm to use for the default encryption.
	// Options are AES256 or aws:kms
	SSEAlgorithm string `json:"sseAlgorithm"`
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
