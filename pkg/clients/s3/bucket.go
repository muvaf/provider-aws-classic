/*
Copyright 2019 The Crossplane Authors.

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

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws/awserr"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/crossplane/provider-aws/apis/s3/v1beta1"

	"github.com/aws/aws-sdk-go-v2/service/s3"

	awsclients "github.com/crossplane/provider-aws/pkg/clients"
)

type BucketClient interface {
	HeadBucketRequest(input *s3.HeadBucketInput) s3.HeadBucketRequest
	CreateBucketRequest(input *s3.CreateBucketInput) s3.CreateBucketRequest
	DeleteBucketRequest(input *s3.DeleteBucketInput) s3.DeleteBucketRequest

	PutBucketEncryptionRequest(input *s3.PutBucketEncryptionInput) s3.PutBucketEncryptionRequest
	GetBucketEncryptionRequest(input *s3.GetBucketEncryptionInput) s3.GetBucketEncryptionRequest

	PutBucketVersioningRequest(input *s3.PutBucketVersioningInput) s3.PutBucketVersioningRequest
	GetBucketVersioningRequest(input *s3.GetBucketVersioningInput) s3.GetBucketVersioningRequest

	PutBucketAccelerateConfigurationRequest(input *s3.PutBucketAccelerateConfigurationInput) s3.PutBucketAccelerateConfigurationRequest
	GetBucketAccelerateConfigurationRequest(input *s3.GetBucketAccelerateConfigurationInput) s3.GetBucketAccelerateConfigurationRequest

	PutBucketCorsRequest(input *s3.PutBucketCorsInput) s3.PutBucketCorsRequest
	GetBucketCorsRequest(input *s3.GetBucketCorsInput) s3.GetBucketCorsRequest
}

// NewVpcClient returns a new client using AWS credentials as JSON encoded data.
func NewClient(ctx context.Context, credentials []byte, region string, auth awsclients.AuthMethod) (BucketClient, error) {
	cfg, err := auth(ctx, credentials, awsclients.DefaultSection, region)
	if cfg == nil {
		return nil, err
	}
	return s3.New(*cfg), nil
}

func GenerateCreateBucketInput(name string, s v1beta1.BucketParameters) *s3.CreateBucketInput {
	cbi := &s3.CreateBucketInput{
		ACL:                        s3.BucketCannedACL(aws.StringValue(s.ACL)),
		Bucket:                     aws.String(name),
		GrantFullControl:           s.GrantFullControl,
		GrantRead:                  s.GrantRead,
		GrantReadACP:               s.GrantReadACP,
		GrantWrite:                 s.GrantWrite,
		GrantWriteACP:              s.GrantWriteACP,
		ObjectLockEnabledForBucket: s.ObjectLockEnabledForBucket,
	}
	if len(awsclients.StringValue(s.LocationConstraint)) != 0 {
		cbi.CreateBucketConfiguration = &s3.CreateBucketConfiguration{LocationConstraint: s3.BucketLocationConstraint(awsclients.StringValue(s.LocationConstraint))}
	}
	return cbi
}

func GeneratePutBucketEncryptionInput(name string, s v1beta1.BucketParameters) *s3.PutBucketEncryptionInput {
	if s.ServerSideEncryptionConfiguration == nil {
		return nil
	}
	bei := &s3.PutBucketEncryptionInput{
		Bucket:                            aws.String(name),
		ServerSideEncryptionConfiguration: &s3.ServerSideEncryptionConfiguration{},
	}
	for _, rule := range s.ServerSideEncryptionConfiguration.Rules {
		bei.ServerSideEncryptionConfiguration.Rules = append(bei.ServerSideEncryptionConfiguration.Rules, s3.ServerSideEncryptionRule{
			ApplyServerSideEncryptionByDefault: &s3.ServerSideEncryptionByDefault{
				KMSMasterKeyID: rule.ApplyServerSideEncryptionByDefault.KMSMasterKeyID,
				SSEAlgorithm:   s3.ServerSideEncryption(rule.ApplyServerSideEncryptionByDefault.SSEAlgorithm),
			},
		})
	}
	return bei
}

func GeneratePutBucketVersioningInput(name string, s v1beta1.BucketParameters) *s3.PutBucketVersioningInput {
	if s.VersioningConfiguration == nil {
		return nil
	}
	return &s3.PutBucketVersioningInput{
		Bucket: aws.String(name),
		VersioningConfiguration: &s3.VersioningConfiguration{
			MFADelete: s3.MFADelete(aws.StringValue(s.VersioningConfiguration.MFADelete)),
			Status:    s3.BucketVersioningStatus(aws.StringValue(s.VersioningConfiguration.Status)),
		},
	}
}

func GenerateAccelerateConfigurationInput(name string, s v1beta1.BucketParameters) *s3.PutBucketAccelerateConfigurationInput {
	if s.AccelerateConfiguration == nil {
		return nil
	}
	return &s3.PutBucketAccelerateConfigurationInput{
		Bucket:                  aws.String(name),
		AccelerateConfiguration: &s3.AccelerateConfiguration{Status: s3.BucketAccelerateStatus(s.AccelerateConfiguration.Status)},
	}
}

func GeneratePutBucketCorsInput(name string, s v1beta1.BucketParameters) *s3.PutBucketCorsInput {
	if s.CORSConfiguration == nil {
		return nil
	}
	bci := &s3.PutBucketCorsInput{CORSConfiguration: &s3.CORSConfiguration{}}
	for _, cors := range s.CORSConfiguration.CORSRules {
		bci.CORSConfiguration.CORSRules = append(bci.CORSConfiguration.CORSRules, s3.CORSRule{
			AllowedHeaders: cors.AllowedHeaders,
			AllowedMethods: cors.AllowedMethods,
			AllowedOrigins: cors.AllowedOrigins,
			ExposeHeaders:  cors.ExposeHeaders,
			MaxAgeSeconds:  cors.MaxAgeSeconds,
		})
	}
	return bci
}

// IsNotFound helper function to test for ErrCodeNoSuchEntityException error
func IsNotFound(err error) bool {
	if err == nil {
		return false
	}
	if bucketErr, ok := err.(awserr.Error); ok && bucketErr.Code() == "NotFound" {
		return true
	}
	return false
}
