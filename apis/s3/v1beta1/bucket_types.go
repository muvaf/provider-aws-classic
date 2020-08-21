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

	// VersioningConfiguration describes the versioning state of an Amazon S3 bucket.
	// +optional
	VersioningConfiguration *VersioningConfiguration `json:"versioningConfiguration,omitempty"`

	// AccelerateConfiguration configures the transfer acceleration state for an
	// Amazon S3 bucket. For more information, see Amazon S3 Transfer Acceleration
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/transfer-acceleration.html)
	// in the Amazon Simple Storage Service Developer Guide.
	// +optional
	AccelerateConfiguration *AccelerateConfiguration `json:"accelerateConfiguration,omitempty"`

	// Describes the cross-origin access configuration for objects in an Amazon
	// S3 bucket. For more information, see Enabling Cross-Origin Resource Sharing
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) in the Amazon
	// Simple Storage Service Developer Guide.
	// +optional
	CORSConfiguration *CORSConfiguration `json:"corsConfiguration,omitempty"`

	// Specifies website configuration parameters for an Amazon S3 bucket.
	WebsiteConfiguration *WebsiteConfiguration `json:"websiteConfiguration,omitempty"`
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

// VersioningConfiguration describes the versioning state of an Amazon S3 bucket.
type VersioningConfiguration struct {
	// MFADelete specifies whether MFA delete is enabled in the bucket versioning configuration.
	// This element is only returned if the bucket has been configured with MFA
	// delete. If the bucket has never been so configured, this element is not returned.
	// +kubebuilder:validation:Enum=Enabled;Disabled
	MFADelete *string `json:"mfaDelete"`

	// Status is the desired versioning state of the bucket.
	// +kubebuilder:validation:Enum=Enabled;Suspended
	Status *string `json:"status"`
}

// AccelerateConfiguration configures the transfer acceleration state for an
// Amazon S3 bucket. For more information, see Amazon S3 Transfer Acceleration
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/transfer-acceleration.html)
// in the Amazon Simple Storage Service Developer Guide.
type AccelerateConfiguration struct {
	// Status specifies the transfer acceleration status of the bucket.
	// +kubebuilder:validation:Enum=Enabled;Suspended
	Status string `json:"status"`
}

// CORSConfiguration describes the cross-origin access configuration for objects
// in an Amazon S3 bucket. For more information, see Enabling Cross-Origin Resource Sharing
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) in the Amazon
// Simple Storage Service Developer Guide.
type CORSConfiguration struct {
	// A set of origins and methods (cross-origin access that you want to allow).
	// You can add up to 100 rules to the configuration.
	CORSRules []CORSRule `json:"corsRules"`
}

// CORSRule specifies a cross-origin access rule for an Amazon S3 bucket.
type CORSRule struct {
	// Headers that are specified in the Access-Control-Request-Headers header.
	// These headers are allowed in a preflight OPTIONS request. In response to
	// any preflight OPTIONS request, Amazon S3 returns any requested headers that
	// are allowed.
	// +optional
	AllowedHeaders []string `json:"allowedHeaders,omitempty"`

	// An HTTP method that you allow the origin to execute. Valid values are GET,
	// PUT, HEAD, POST, and DELETE.
	AllowedMethods []string `json:"allowedMethods"`

	// One or more origins you want customers to be able to access the bucket from.
	AllowedOrigins []string `json:"allowedOrigins"`

	// One or more headers in the response that you want customers to be able to
	// access from their applications (for example, from a JavaScript XMLHttpRequest
	// object).
	// +optional
	ExposeHeaders []string `json:"exposeHeaders,omitempty"`

	// The time in seconds that your browser is to cache the preflight response
	// for the specified resource.
	// +optional
	MaxAgeSeconds *int64 `json:"maxAgeSeconds,omitempty"`
}

// WebsiteConfiguration specifies website configuration parameters for an Amazon S3 bucket.
type WebsiteConfiguration struct {
	// The name of the error document for the website.
	// +optional
	ErrorDocument *ErrorDocument `json:"errorDocument,omitempty"`

	// The name of the index document for the website.
	// +optional
	IndexDocument *IndexDocument `json:"indexDocument,omitempty"`

	// The redirect behavior for every request to this bucket's website endpoint.
	// If you specify this property, you can't specify any other property.
	// +optional
	RedirectAllRequestsTo *RedirectAllRequestsTo `json:"redirectAllRequestsTo,omitempty"`

	// Rules that define when a redirect is applied and the redirect behavior.
	// +optional
	RoutingRules []RoutingRule `json:"routingRules,omitempty"`
}

// ErrorDocument is the error information.
type ErrorDocument struct {
	// The object key name to use when a 4XX class error occurs.
	Key string `json:"key"`
}

// IndexDocument is container for the Suffix element.
type IndexDocument struct {
	// A suffix that is appended to a request that is for a directory on the website
	// endpoint (for example,if the suffix is index.html and you make a request
	// to samplebucket/images/ the data that is returned will be for the object
	// with the key name images/index.html) The suffix must not be empty and must
	// not include a slash character.
	Suffix string `json:"suffix"`
}

// RedirectAllRequestsTo specifies the redirect behavior of all requests to a
// website endpoint of an Amazon S3 bucket.
type RedirectAllRequestsTo struct {
	// Name of the host where requests are redirected.
	HostName string `json:"hostName"`

	// Protocol to use when redirecting requests. The default is the protocol that
	// is used in the original request.
	// +kubebuilder:validation:Enum=http;https
	Protocol string `json:"protocol"`
}

// RoutingRule specifies the redirect behavior and when a redirect is applied.
type RoutingRule struct {
	// A container for describing a condition that must be met for the specified
	// redirect to apply. For example, 1. If request is for pages in the /docs folder,
	// redirect to the /documents folder. 2. If request results in HTTP error 4xx,
	// redirect request to another host where you might process the error.
	// +optional
	Condition *Condition `json:"condition,omitempty"`

	// Container for redirect information. You can redirect requests to another
	// host, to another page, or with another protocol. In the event of an error,
	// you can specify a different error code to return.
	Redirect Redirect `json:"redirect"`
}

// Condition is a container for describing a condition that must be met for the specified
// redirect to apply. For example, 1. If request is for pages in the /docs folder,
// redirect to the /documents folder. 2. If request results in HTTP error 4xx,
// redirect request to another host where you might process the error.
type Condition struct {
	// The HTTP error code when the redirect is applied. In the event of an error,
	// if the error code equals this value, then the specified redirect is applied.
	// Required when parent element Condition is specified and sibling KeyPrefixEquals
	// is not specified. If both are specified, then both must be true for the redirect
	// to be applied.
	// +optional
	HttpErrorCodeReturnedEquals *string `json:"httpErrorCodeReturnedEquals,omitempty"`

	// The object key name prefix when the redirect is applied. For example, to
	// redirect requests for ExamplePage.html, the key prefix will be ExamplePage.html.
	// To redirect request for all pages with the prefix docs/, the key prefix will
	// be /docs, which identifies all objects in the docs/ folder. Required when
	// the parent element Condition is specified and sibling HttpErrorCodeReturnedEquals
	// is not specified. If both conditions are specified, both must be true for
	// the redirect to be applied.
	// +optional
	KeyPrefixEquals *string `json:"keyPrefixEquals,omitempty"`
}

// Redirect specifies how requests are redirected. In the event of an error, you can
// specify a different error code to return.
type Redirect struct {
	// The host name to use in the redirect request.
	// +optional
	HostName *string `json:"keyPrefixEquals,omitempty"`

	// The HTTP redirect code to use on the response. Not required if one of the
	// siblings is present.
	// +optional
	HttpRedirectCode *string `json:"httpRedirectCode,omitempty"`

	// Protocol to use when redirecting requests. The default is the protocol that
	// is used in the original request.
	// +kubebuilder:validation:Enum=http;https
	Protocol string `json:"protocol"`

	// The object key prefix to use in the redirect request. For example, to redirect
	// requests for all pages with prefix docs/ (objects in the docs/ folder) to
	// documents/, you can set a condition block with KeyPrefixEquals set to docs/
	// and in the Redirect set ReplaceKeyPrefixWith to /documents. Not required
	// if one of the siblings is present. Can be present only if ReplaceKeyWith
	// is not provided.
	// +optional
	ReplaceKeyPrefixWith *string `json:"replaceKeyPrefixWith,omitempty"`

	// The specific object key to use in the redirect request. For example, redirect
	// request to error.html. Not required if one of the siblings is present. Can
	// be present only if ReplaceKeyPrefixWith is not provided.
	// +optional
	ReplaceKeyWith *string `json:"replaceKeyWith,omitempty"`
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
