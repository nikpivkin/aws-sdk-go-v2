// Code generated by smithy-go-codegen DO NOT EDIT.

package inspector2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/inspector2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Activates or deactivates Amazon Inspector deep inspection for the provided
// member accounts in your organization. You must be the delegated administrator of
// an organization in Amazon Inspector to use this API.
func (c *Client) BatchUpdateMemberEc2DeepInspectionStatus(ctx context.Context, params *BatchUpdateMemberEc2DeepInspectionStatusInput, optFns ...func(*Options)) (*BatchUpdateMemberEc2DeepInspectionStatusOutput, error) {
	if params == nil {
		params = &BatchUpdateMemberEc2DeepInspectionStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchUpdateMemberEc2DeepInspectionStatus", params, optFns, c.addOperationBatchUpdateMemberEc2DeepInspectionStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchUpdateMemberEc2DeepInspectionStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchUpdateMemberEc2DeepInspectionStatusInput struct {

	// The unique identifiers for the Amazon Web Services accounts to change Amazon
	// Inspector deep inspection status for.
	//
	// This member is required.
	AccountIds []types.MemberAccountEc2DeepInspectionStatus

	noSmithyDocumentSerde
}

type BatchUpdateMemberEc2DeepInspectionStatusOutput struct {

	// An array of objects that provide details for each of the accounts that Amazon
	// Inspector deep inspection status was successfully changed for.
	AccountIds []types.MemberAccountEc2DeepInspectionStatusState

	// An array of objects that provide details for each of the accounts that Amazon
	// Inspector deep inspection status could not be successfully changed for.
	FailedAccountIds []types.FailedMemberAccountEc2DeepInspectionStatusState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchUpdateMemberEc2DeepInspectionStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchUpdateMemberEc2DeepInspectionStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchUpdateMemberEc2DeepInspectionStatus{}, middleware.After)
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
	if err = addOpBatchUpdateMemberEc2DeepInspectionStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchUpdateMemberEc2DeepInspectionStatus(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchUpdateMemberEc2DeepInspectionStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "inspector2",
		OperationName: "BatchUpdateMemberEc2DeepInspectionStatus",
	}
}
