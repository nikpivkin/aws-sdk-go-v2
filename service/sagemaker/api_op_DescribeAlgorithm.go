// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns a description of the specified algorithm that is in your account.
func (c *Client) DescribeAlgorithm(ctx context.Context, params *DescribeAlgorithmInput, optFns ...func(*Options)) (*DescribeAlgorithmOutput, error) {
	if params == nil {
		params = &DescribeAlgorithmInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAlgorithm", params, optFns, c.addOperationDescribeAlgorithmMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAlgorithmOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAlgorithmInput struct {

	// The name of the algorithm to describe.
	//
	// This member is required.
	AlgorithmName *string

	noSmithyDocumentSerde
}

type DescribeAlgorithmOutput struct {

	// The Amazon Resource Name (ARN) of the algorithm.
	//
	// This member is required.
	AlgorithmArn *string

	// The name of the algorithm being described.
	//
	// This member is required.
	AlgorithmName *string

	// The current status of the algorithm.
	//
	// This member is required.
	AlgorithmStatus types.AlgorithmStatus

	// Details about the current status of the algorithm.
	//
	// This member is required.
	AlgorithmStatusDetails *types.AlgorithmStatusDetails

	// A timestamp specifying when the algorithm was created.
	//
	// This member is required.
	CreationTime *time.Time

	// Details about training jobs run by this algorithm.
	//
	// This member is required.
	TrainingSpecification *types.TrainingSpecification

	// A brief summary about the algorithm.
	AlgorithmDescription *string

	// Whether the algorithm is certified to be listed in Amazon Web Services
	// Marketplace.
	CertifyForMarketplace bool

	// Details about inference jobs that the algorithm runs.
	InferenceSpecification *types.InferenceSpecification

	// The product identifier of the algorithm.
	ProductId *string

	// Details about configurations for one or more training jobs that SageMaker runs
	// to test the algorithm.
	ValidationSpecification *types.AlgorithmValidationSpecification

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeAlgorithmMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeAlgorithm{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeAlgorithm{}, middleware.After)
	if err != nil {
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
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDescribeAlgorithmValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAlgorithm(options.Region), middleware.Before); err != nil {
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
	return nil
}

func newServiceMetadataMiddleware_opDescribeAlgorithm(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "DescribeAlgorithm",
	}
}
