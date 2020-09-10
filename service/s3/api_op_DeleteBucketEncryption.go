// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This implementation of the DELETE operation removes default encryption from the
// bucket. For information about the Amazon S3 default encryption feature, see
// Amazon S3 Default Bucket Encryption
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-encryption.html) in the
// Amazon Simple Storage Service Developer Guide. To use this operation, you must
// have permissions to perform the s3:PutEncryptionConfiguration action. The bucket
// owner has this permission by default. The bucket owner can grant this permission
// to others. For more information about permissions, see Permissions Related to
// Bucket Subresource Operations
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources)
// and Managing Access Permissions to your Amazon S3 Resources
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-access-control.html) in the
// Amazon Simple Storage Service Developer Guide.  <p class="title"> <b>Related
// Resources</b> </p> <ul> <li> <p> <a>PutBucketEncryption</a> </p> </li> <li> <p>
// <a>GetBucketEncryption</a> </p> </li> </ul>
func (c *Client) DeleteBucketEncryption(ctx context.Context, params *DeleteBucketEncryptionInput, optFns ...func(*Options)) (*DeleteBucketEncryptionOutput, error) {
	stack := middleware.NewStack("DeleteBucketEncryption", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpDeleteBucketEncryptionMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpDeleteBucketEncryptionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteBucketEncryption(options.Region), middleware.Before)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "DeleteBucketEncryption",
			Err:           err,
		}
	}
	out := result.(*DeleteBucketEncryptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteBucketEncryptionInput struct {
	// The name of the bucket containing the server-side encryption configuration to
	// delete.
	Bucket *string
}

type DeleteBucketEncryptionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpDeleteBucketEncryptionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpDeleteBucketEncryption{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpDeleteBucketEncryption{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteBucketEncryption(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "DeleteBucketEncryption",
	}
}
