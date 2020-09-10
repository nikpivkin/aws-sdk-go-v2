// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes the null version (if there is one) of an object and inserts a delete
// marker, which becomes the latest version of the object. If there isn't a null
// version, Amazon S3 does not remove any objects.  <p>To remove a specific
// version, you must be the bucket owner and you must use the version Id
// subresource. Using this subresource permanently deletes the version. If the
// object deleted is a delete marker, Amazon S3 sets the response header,
// <code>x-amz-delete-marker</code>, to true. </p> <p>If the object you want to
// delete is in a bucket where the bucket versioning configuration is MFA Delete
// enabled, you must include the <code>x-amz-mfa</code> request header in the
// DELETE <code>versionId</code> request. Requests that include
// <code>x-amz-mfa</code> must use HTTPS. </p> <p> For more information about MFA
// Delete, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMFADelete.html">Using
// MFA Delete</a>. To see sample requests that use versioning, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/API/RESTObjectDELETE.html#ExampleVersionObjectDelete">Sample
// Request</a>. </p> <p>You can delete objects by explicitly calling the DELETE
// Object API or configure its lifecycle (<a>PutBucketLifecycle</a>) to enable
// Amazon S3 to remove them for you. If you want to block users or accounts from
// removing or deleting objects from your bucket, you must deny them the
// <code>s3:DeleteObject</code>, <code>s3:DeleteObjectVersion</code>, and
// <code>s3:PutLifeCycleConfiguration</code> actions. </p> <p>The following
// operation is related to <code>DeleteObject</code>:</p> <ul> <li> <p>
// <a>PutObject</a> </p> </li> </ul>
func (c *Client) DeleteObject(ctx context.Context, params *DeleteObjectInput, optFns ...func(*Options)) (*DeleteObjectOutput, error) {
	stack := middleware.NewStack("DeleteObject", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpDeleteObjectMiddlewares(stack)
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
	addOpDeleteObjectValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteObject(options.Region), middleware.Before)

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
			OperationName: "DeleteObject",
			Err:           err,
		}
	}
	out := result.(*DeleteObjectOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteObjectInput struct {
	// VersionId used to reference a specific version of the object.
	VersionId *string
	// Key name of the object to delete.
	Key *string
	// Confirms that the requester knows that they will be charged for the request.
	// Bucket owners need not specify this parameter in their requests. For information
	// about downloading objects from requester pays buckets, see Downloading Objects
	// in Requestor Pays Buckets
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html)
	// in the Amazon S3 Developer Guide.
	RequestPayer types.RequestPayer
	// The concatenation of the authentication device's serial number, a space, and the
	// value that is displayed on your authentication device. Required to permanently
	// delete a versioned object if versioning is configured with MFA delete enabled.
	MFA *string
	// Indicates whether S3 Object Lock should bypass Governance-mode restrictions to
	// process this operation.
	BypassGovernanceRetention *bool
	// The bucket name of the bucket containing the object. When using this API with an
	// access point, you must direct requests to the access point hostname. The access
	// point hostname takes the form
	// AccessPointName-AccountId.s3-accesspoint.Region.amazonaws.com. When using this
	// operation using an access point through the AWS SDKs, you provide the access
	// point ARN in place of the bucket name. For more information about access point
	// ARNs, see Using Access Points
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-access-points.html) in
	// the Amazon Simple Storage Service Developer Guide.
	Bucket *string
}

type DeleteObjectOutput struct {
	// Returns the version ID of the delete marker created as a result of the DELETE
	// operation.
	VersionId *string
	// If present, indicates that the requester was successfully charged for the
	// request.
	RequestCharged types.RequestCharged
	// Specifies whether the versioned object that was permanently deleted was (true)
	// or was not (false) a delete marker.
	DeleteMarker *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpDeleteObjectMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpDeleteObject{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpDeleteObject{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteObject(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "DeleteObject",
	}
}
