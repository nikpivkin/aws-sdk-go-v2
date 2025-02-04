// Code generated by smithy-go-codegen DO NOT EDIT.

package firehose

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/firehose/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Writes multiple data records into a delivery stream in a single call, which can
// achieve higher throughput per producer than when writing single records. To
// write single data records into a delivery stream, use PutRecord . Applications
// using these operations are referred to as producers. Kinesis Data Firehose
// accumulates and publishes a particular metric for a customer account in one
// minute intervals. It is possible that the bursts of incoming bytes/records
// ingested to a delivery stream last only for a few seconds. Due to this, the
// actual spikes in the traffic might not be fully visible in the customer's 1
// minute CloudWatch metrics. For information about service quota, see Amazon
// Kinesis Data Firehose Quota (https://docs.aws.amazon.com/firehose/latest/dev/limits.html)
// . Each PutRecordBatch request supports up to 500 records. Each record in the
// request can be as large as 1,000 KB (before base64 encoding), up to a limit of 4
// MB for the entire request. These limits cannot be changed. You must specify the
// name of the delivery stream and the data record when using PutRecord . The data
// record consists of a data blob that can be up to 1,000 KB in size, and any kind
// of data. For example, it could be a segment from a log file, geographic location
// data, website clickstream data, and so on. Kinesis Data Firehose buffers records
// before delivering them to the destination. To disambiguate the data blobs at the
// destination, a common solution is to use delimiters in the data, such as a
// newline ( \n ) or some other character unique within the data. This allows the
// consumer application to parse individual data items when reading the data from
// the destination. The PutRecordBatch response includes a count of failed
// records, FailedPutCount , and an array of responses, RequestResponses . Even if
// the PutRecordBatch call succeeds, the value of FailedPutCount may be greater
// than 0, indicating that there are records for which the operation didn't
// succeed. Each entry in the RequestResponses array provides additional
// information about the processed record. It directly correlates with a record in
// the request array using the same ordering, from the top to the bottom. The
// response array always includes the same number of records as the request array.
// RequestResponses includes both successfully and unsuccessfully processed
// records. Kinesis Data Firehose tries to process all records in each
// PutRecordBatch request. A single record failure does not stop the processing of
// subsequent records. A successfully processed record includes a RecordId value,
// which is unique for the record. An unsuccessfully processed record includes
// ErrorCode and ErrorMessage values. ErrorCode reflects the type of error, and is
// one of the following values: ServiceUnavailableException or InternalFailure .
// ErrorMessage provides more detailed information about the error. If there is an
// internal server error or a timeout, the write might have completed or it might
// have failed. If FailedPutCount is greater than 0, retry the request, resending
// only those records that might have failed processing. This minimizes the
// possible duplicate records and also reduces the total bytes sent (and
// corresponding charges). We recommend that you handle any duplicates at the
// destination. If PutRecordBatch throws ServiceUnavailableException , the API is
// automatically reinvoked (retried) 3 times. If the exception persists, it is
// possible that the throughput limits have been exceeded for the delivery stream.
// Re-invoking the Put API operations (for example, PutRecord and PutRecordBatch)
// can result in data duplicates. For larger data assets, allow for a longer time
// out before retrying Put API operations. Data records sent to Kinesis Data
// Firehose are stored for 24 hours from the time they are added to a delivery
// stream as it attempts to send the records to the destination. If the destination
// is unreachable for more than 24 hours, the data is no longer available. Don't
// concatenate two or more base64 strings to form the data fields of your records.
// Instead, concatenate the raw data, then perform base64 encoding.
func (c *Client) PutRecordBatch(ctx context.Context, params *PutRecordBatchInput, optFns ...func(*Options)) (*PutRecordBatchOutput, error) {
	if params == nil {
		params = &PutRecordBatchInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutRecordBatch", params, optFns, c.addOperationPutRecordBatchMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutRecordBatchOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutRecordBatchInput struct {

	// The name of the delivery stream.
	//
	// This member is required.
	DeliveryStreamName *string

	// One or more records.
	//
	// This member is required.
	Records []types.Record

	noSmithyDocumentSerde
}

type PutRecordBatchOutput struct {

	// The number of records that might have failed processing. This number might be
	// greater than 0 even if the PutRecordBatch call succeeds. Check FailedPutCount
	// to determine whether there are records that you need to resend.
	//
	// This member is required.
	FailedPutCount *int32

	// The results array. For each record, the index of the response element is the
	// same as the index used in the request array.
	//
	// This member is required.
	RequestResponses []types.PutRecordBatchResponseEntry

	// Indicates whether server-side encryption (SSE) was enabled during this
	// operation.
	Encrypted *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutRecordBatchMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutRecordBatch{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutRecordBatch{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addPutRecordBatchResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpPutRecordBatchValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutRecordBatch(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addendpointDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opPutRecordBatch(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "firehose",
		OperationName: "PutRecordBatch",
	}
}

type opPutRecordBatchResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opPutRecordBatchResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opPutRecordBatchResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	var resolvedEndpoint smithyendpoints.Endpoint
	resolvedEndpoint, err = m.EndpointResolver.ResolveEndpoint(ctx, params)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	}

	req.URL = &resolvedEndpoint.URI

	for k := range resolvedEndpoint.Headers {
		req.Header.Set(
			k,
			resolvedEndpoint.Headers.Get(k),
		)
	}

	authSchemes, err := internalauth.GetAuthenticationSchemes(&resolvedEndpoint.Properties)
	if err != nil {
		var nfe *internalauth.NoAuthenticationSchemesFoundError
		if errors.As(err, &nfe) {
			// if no auth scheme is found, default to sigv4
			signingName := "firehose"
			signingRegion := m.BuiltInResolver.(*builtInResolver).Region
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)

		}
		var ue *internalauth.UnSupportedAuthenticationSchemeSpecifiedError
		if errors.As(err, &ue) {
			return out, metadata, fmt.Errorf(
				"This operation requests signer version(s) %v but the client only supports %v",
				ue.UnsupportedSchemes,
				internalauth.SupportedSchemes,
			)
		}
	}

	for _, authScheme := range authSchemes {
		switch authScheme.(type) {
		case *internalauth.AuthenticationSchemeV4:
			v4Scheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4)
			var signingName, signingRegion string
			if v4Scheme.SigningName == nil {
				signingName = "firehose"
			} else {
				signingName = *v4Scheme.SigningName
			}
			if v4Scheme.SigningRegion == nil {
				signingRegion = m.BuiltInResolver.(*builtInResolver).Region
			} else {
				signingRegion = *v4Scheme.SigningRegion
			}
			if v4Scheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4Scheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
			break
		case *internalauth.AuthenticationSchemeV4A:
			v4aScheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4A)
			if v4aScheme.SigningName == nil {
				v4aScheme.SigningName = aws.String("firehose")
			}
			if v4aScheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, *v4aScheme.SigningName)
			ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	return next.HandleSerialize(ctx, in)
}

func addPutRecordBatchResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opPutRecordBatchResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
