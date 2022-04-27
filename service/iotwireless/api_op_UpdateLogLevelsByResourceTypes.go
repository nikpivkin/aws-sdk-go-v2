// Code generated by smithy-go-codegen DO NOT EDIT.

package iotwireless

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotwireless/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Set default log level, or log levels by resource types. This can be for wireless
// device log options or wireless gateways log options and is used to control the
// log messages that'll be displayed in CloudWatch.
func (c *Client) UpdateLogLevelsByResourceTypes(ctx context.Context, params *UpdateLogLevelsByResourceTypesInput, optFns ...func(*Options)) (*UpdateLogLevelsByResourceTypesOutput, error) {
	if params == nil {
		params = &UpdateLogLevelsByResourceTypesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateLogLevelsByResourceTypes", params, optFns, c.addOperationUpdateLogLevelsByResourceTypesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateLogLevelsByResourceTypesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateLogLevelsByResourceTypesInput struct {

	// The log level for a log message. The log levels can be disabled, or set to ERROR
	// to display less verbose logs containing only error information, or to INFO for
	// more detailed logs.
	DefaultLogLevel types.LogLevel

	// The list of wireless device log options.
	WirelessDeviceLogOptions []types.WirelessDeviceLogOption

	// The list of wireless gateway log options.
	WirelessGatewayLogOptions []types.WirelessGatewayLogOption

	noSmithyDocumentSerde
}

type UpdateLogLevelsByResourceTypesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateLogLevelsByResourceTypesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateLogLevelsByResourceTypes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateLogLevelsByResourceTypes{}, middleware.After)
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
	if err = addOpUpdateLogLevelsByResourceTypesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateLogLevelsByResourceTypes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateLogLevelsByResourceTypes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotwireless",
		OperationName: "UpdateLogLevelsByResourceTypes",
	}
}
