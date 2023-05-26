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

// Updates properties of a FUOTA task.
func (c *Client) UpdateFuotaTask(ctx context.Context, params *UpdateFuotaTaskInput, optFns ...func(*Options)) (*UpdateFuotaTaskOutput, error) {
	if params == nil {
		params = &UpdateFuotaTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateFuotaTask", params, optFns, c.addOperationUpdateFuotaTaskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateFuotaTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateFuotaTaskInput struct {

	// The ID of a FUOTA task.
	//
	// This member is required.
	Id *string

	// The description of the new resource.
	Description *string

	// The S3 URI points to a firmware update image that is to be used with a FUOTA
	// task.
	FirmwareUpdateImage *string

	// The firmware update role that is to be used with a FUOTA task.
	FirmwareUpdateRole *string

	// The interval for sending fragments in milliseconds, rounded to the nearest
	// second. This interval only determines the timing for when the Cloud sends down
	// the fragments to yor device. There can be a delay for when your device will
	// receive these fragments. This delay depends on the device's class and the
	// communication delay with the cloud.
	FragmentIntervalMS *int32

	// The size of each fragment in bytes. This parameter is supported only for FUOTA
	// tasks with multicast groups.
	FragmentSizeBytes *int32

	// The LoRaWAN information used with a FUOTA task.
	LoRaWAN *types.LoRaWANFuotaTask

	// The name of a FUOTA task.
	Name *string

	// The percentage of the added fragments that are redundant. For example, if the
	// size of the firmware image file is 100 bytes and the fragment size is 10 bytes,
	// with RedundancyPercent set to 50(%), the final number of encoded fragments is
	// (100 / 10) + (100 / 10 * 50%) = 15.
	RedundancyPercent *int32

	noSmithyDocumentSerde
}

type UpdateFuotaTaskOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateFuotaTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateFuotaTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateFuotaTask{}, middleware.After)
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
	if err = addOpUpdateFuotaTaskValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateFuotaTask(options.Region), middleware.Before); err != nil {
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
	return nil
}

func newServiceMetadataMiddleware_opUpdateFuotaTask(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotwireless",
		OperationName: "UpdateFuotaTask",
	}
}
