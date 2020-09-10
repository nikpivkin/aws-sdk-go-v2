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

// Returns metadata about all of the versions of objects in a bucket. You can also
// use request parameters as selection criteria to return metadata about a subset
// of all the object versions.  <note> <p> A 200 OK response can contain valid or
// invalid XML. Make sure to design your application to parse the contents of the
// response and handle it appropriately.</p> </note> <p>To use this operation, you
// must have READ access to the bucket.</p> <p>The following operations are related
// to <code>ListObjectVersions</code>:</p> <ul> <li> <p> <a>ListObjectsV2</a> </p>
// </li> <li> <p> <a>GetObject</a> </p> </li> <li> <p> <a>PutObject</a> </p> </li>
// <li> <p> <a>DeleteObject</a> </p> </li> </ul>
func (c *Client) ListObjectVersions(ctx context.Context, params *ListObjectVersionsInput, optFns ...func(*Options)) (*ListObjectVersionsOutput, error) {
	stack := middleware.NewStack("ListObjectVersions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpListObjectVersionsMiddlewares(stack)
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
	addOpListObjectVersionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListObjectVersions(options.Region), middleware.Before)

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
			OperationName: "ListObjectVersions",
			Err:           err,
		}
	}
	out := result.(*ListObjectVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListObjectVersionsInput struct {
	// Use this parameter to select only those keys that begin with the specified
	// prefix. You can use prefixes to separate a bucket into different groupings of
	// keys. (You can think of using prefix to make groups in the same way you'd use a
	// folder in a file system.) You can use prefix with delimiter to roll up numerous
	// objects into a single result under CommonPrefixes.
	Prefix *string
	// Requests Amazon S3 to encode the object keys in the response and specifies the
	// encoding method to use. An object key may contain any Unicode character;
	// however, XML 1.0 parser cannot parse some characters, such as characters with an
	// ASCII value from 0 to 10. For characters that are not supported in XML 1.0, you
	// can add this parameter to request that Amazon S3 encode the keys in the
	// response.
	EncodingType types.EncodingType
	// Specifies the key to start with when listing objects in a bucket.
	KeyMarker *string
	// Sets the maximum number of keys returned in the response. By default the API
	// returns up to 1,000 key names. The response might contain fewer keys but will
	// never contain more. If additional keys satisfy the search criteria, but were not
	// returned because max-keys was exceeded, the response contains true. To return
	// the additional keys, see key-marker and version-id-marker.
	MaxKeys *int32
	// Specifies the object version you want to start listing from.
	VersionIdMarker *string
	// The bucket name that contains the objects. When using this API with an access
	// point, you must direct requests to the access point hostname. The access point
	// hostname takes the form
	// AccessPointName-AccountId.s3-accesspoint.Region.amazonaws.com. When using this
	// operation using an access point through the AWS SDKs, you provide the access
	// point ARN in place of the bucket name. For more information about access point
	// ARNs, see Using Access Points
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-access-points.html) in
	// the Amazon Simple Storage Service Developer Guide.
	Bucket *string
	// A delimiter is a character that you specify to group keys. All keys that contain
	// the same string between the prefix and the first occurrence of the delimiter are
	// grouped under a single result element in CommonPrefixes. These groups are
	// counted as one result against the max-keys limitation. These keys are not
	// returned elsewhere in the response.
	Delimiter *string
}

type ListObjectVersionsOutput struct {
	// Container for an object that is a delete marker.
	DeleteMarkers []*types.DeleteMarkerEntry
	// All of the keys rolled up into a common prefix count as a single return when
	// calculating the number of returns.
	CommonPrefixes []*types.CommonPrefix
	// Selects objects that start with the value supplied by this parameter.
	Prefix *string
	// When the number of responses exceeds the value of MaxKeys, NextKeyMarker
	// specifies the first key not returned that satisfies the search criteria. Use
	// this value for the key-marker request parameter in a subsequent request.
	NextKeyMarker *string
	// Bucket name.
	Name *string
	// A flag that indicates whether Amazon S3 returned all of the results that
	// satisfied the search criteria. If your results were truncated, you can make a
	// follow-up paginated request using the NextKeyMarker and NextVersionIdMarker
	// response parameters as a starting place in another request to return the rest of
	// the results.
	IsTruncated *bool
	// Encoding type used by Amazon S3 to encode object key names in the XML response.
	// <p>If you specify encoding-type request parameter, Amazon S3 includes this
	// element in the response, and returns encoded key name values in the following
	// response elements:</p> <p> <code>KeyMarker, NextKeyMarker, Prefix, Key</code>,
	// and <code>Delimiter</code>.</p>
	EncodingType types.EncodingType
	// Marks the last version of the key returned in a truncated response.
	VersionIdMarker *string
	// When the number of responses exceeds the value of MaxKeys, NextVersionIdMarker
	// specifies the first object version not returned that satisfies the search
	// criteria. Use this value for the version-id-marker request parameter in a
	// subsequent request.
	NextVersionIdMarker *string
	// Specifies the maximum number of objects to return.
	MaxKeys *int32
	// The delimiter grouping the included keys. A delimiter is a character that you
	// specify to group keys. All keys that contain the same string between the prefix
	// and the first occurrence of the delimiter are grouped under a single result
	// element in CommonPrefixes. These groups are counted as one result against the
	// max-keys limitation. These keys are not returned elsewhere in the response.
	Delimiter *string
	// Marks the last key returned in a truncated response.
	KeyMarker *string
	// Container for version information.
	Versions []*types.ObjectVersion

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpListObjectVersionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpListObjectVersions{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpListObjectVersions{}, middleware.After)
}

func newServiceMetadataMiddleware_opListObjectVersions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "ListObjectVersions",
	}
}
