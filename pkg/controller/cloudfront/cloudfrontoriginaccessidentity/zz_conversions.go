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

// Code generated by ack-generate. DO NOT EDIT.

package cloudfrontoriginaccessidentity

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	svcsdk "github.com/aws/aws-sdk-go/service/cloudfront"

	svcapitypes "github.com/crossplane-contrib/provider-aws/apis/cloudfront/v1alpha1"
)

// NOTE(muvaf): We return pointers in case the function needs to start with an
// empty object, hence need to return a new pointer.

// GenerateGetCloudFrontOriginAccessIdentityInput returns input for read
// operation.
func GenerateGetCloudFrontOriginAccessIdentityInput(cr *svcapitypes.CloudFrontOriginAccessIdentity) *svcsdk.GetCloudFrontOriginAccessIdentityInput {
	res := &svcsdk.GetCloudFrontOriginAccessIdentityInput{}

	return res
}

// GenerateCloudFrontOriginAccessIdentity returns the current state in the form of *svcapitypes.CloudFrontOriginAccessIdentity.
func GenerateCloudFrontOriginAccessIdentity(resp *svcsdk.GetCloudFrontOriginAccessIdentityOutput) *svcapitypes.CloudFrontOriginAccessIdentity {
	cr := &svcapitypes.CloudFrontOriginAccessIdentity{}

	if resp.CloudFrontOriginAccessIdentity != nil {
		f0 := &svcapitypes.OriginAccessIdentity{}
		if resp.CloudFrontOriginAccessIdentity.CloudFrontOriginAccessIdentityConfig != nil {
			f0f0 := &svcapitypes.OriginAccessIdentityConfig{}
			if resp.CloudFrontOriginAccessIdentity.CloudFrontOriginAccessIdentityConfig.Comment != nil {
				f0f0.Comment = resp.CloudFrontOriginAccessIdentity.CloudFrontOriginAccessIdentityConfig.Comment
			}
			f0.CloudFrontOriginAccessIdentityConfig = f0f0
		}
		if resp.CloudFrontOriginAccessIdentity.Id != nil {
			f0.ID = resp.CloudFrontOriginAccessIdentity.Id
		}
		if resp.CloudFrontOriginAccessIdentity.S3CanonicalUserId != nil {
			f0.S3CanonicalUserID = resp.CloudFrontOriginAccessIdentity.S3CanonicalUserId
		}
		cr.Status.AtProvider.CloudFrontOriginAccessIdentity = f0
	} else {
		cr.Status.AtProvider.CloudFrontOriginAccessIdentity = nil
	}
	if resp.ETag != nil {
		cr.Status.AtProvider.ETag = resp.ETag
	} else {
		cr.Status.AtProvider.ETag = nil
	}

	return cr
}

// GenerateCreateCloudFrontOriginAccessIdentityInput returns a create input.
func GenerateCreateCloudFrontOriginAccessIdentityInput(cr *svcapitypes.CloudFrontOriginAccessIdentity) *svcsdk.CreateCloudFrontOriginAccessIdentityInput {
	res := &svcsdk.CreateCloudFrontOriginAccessIdentityInput{}

	if cr.Spec.ForProvider.CloudFrontOriginAccessIdentityConfig != nil {
		f0 := &svcsdk.OriginAccessIdentityConfig{}
		if cr.Spec.ForProvider.CloudFrontOriginAccessIdentityConfig.Comment != nil {
			f0.SetComment(*cr.Spec.ForProvider.CloudFrontOriginAccessIdentityConfig.Comment)
		}
		res.SetCloudFrontOriginAccessIdentityConfig(f0)
	}

	return res
}

// GenerateUpdateCloudFrontOriginAccessIdentityInput returns an update input.
func GenerateUpdateCloudFrontOriginAccessIdentityInput(cr *svcapitypes.CloudFrontOriginAccessIdentity) *svcsdk.UpdateCloudFrontOriginAccessIdentityInput {
	res := &svcsdk.UpdateCloudFrontOriginAccessIdentityInput{}

	if cr.Spec.ForProvider.CloudFrontOriginAccessIdentityConfig != nil {
		f0 := &svcsdk.OriginAccessIdentityConfig{}
		if cr.Spec.ForProvider.CloudFrontOriginAccessIdentityConfig.Comment != nil {
			f0.SetComment(*cr.Spec.ForProvider.CloudFrontOriginAccessIdentityConfig.Comment)
		}
		res.SetCloudFrontOriginAccessIdentityConfig(f0)
	}

	return res
}

// GenerateDeleteCloudFrontOriginAccessIdentityInput returns a deletion input.
func GenerateDeleteCloudFrontOriginAccessIdentityInput(cr *svcapitypes.CloudFrontOriginAccessIdentity) *svcsdk.DeleteCloudFrontOriginAccessIdentityInput {
	res := &svcsdk.DeleteCloudFrontOriginAccessIdentityInput{}

	return res
}

// IsNotFound returns whether the given error is of type NotFound or not.
func IsNotFound(err error) bool {
	awsErr, ok := err.(awserr.Error)
	return ok && awsErr.Code() == "NoSuchCloudFrontOriginAccessIdentity"
}
