// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets network performance data.
func (c *Client) GetAwsNetworkPerformanceData(ctx context.Context, params *GetAwsNetworkPerformanceDataInput, optFns ...func(*Options)) (*GetAwsNetworkPerformanceDataOutput, error) {
	if params == nil {
		params = &GetAwsNetworkPerformanceDataInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAwsNetworkPerformanceData", params, optFns, c.addOperationGetAwsNetworkPerformanceDataMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAwsNetworkPerformanceDataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAwsNetworkPerformanceDataInput struct {

	// A list of network performance data queries.
	DataQueries []types.DataQuery

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The ending time for the performance data request. The end time must be
	// formatted as yyyy-mm-ddThh:mm:ss . For example, 2022-06-12T12:00:00.000Z .
	EndTime *time.Time

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	// The starting time for the performance data request. The starting time must be
	// formatted as yyyy-mm-ddThh:mm:ss . For example, 2022-06-10T12:00:00.000Z .
	StartTime *time.Time

	noSmithyDocumentSerde
}

type GetAwsNetworkPerformanceDataOutput struct {

	// The list of data responses.
	DataResponses []types.DataResponse

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetAwsNetworkPerformanceDataMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpGetAwsNetworkPerformanceData{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpGetAwsNetworkPerformanceData{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetAwsNetworkPerformanceData"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addSpanRetryLoop(stack, options); err != nil {
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
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetAwsNetworkPerformanceData(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
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
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = addSpanInitializeStart(stack); err != nil {
		return err
	}
	if err = addSpanInitializeEnd(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestStart(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestEnd(stack); err != nil {
		return err
	}
	return nil
}

// GetAwsNetworkPerformanceDataPaginatorOptions is the paginator options for
// GetAwsNetworkPerformanceData
type GetAwsNetworkPerformanceDataPaginatorOptions struct {
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetAwsNetworkPerformanceDataPaginator is a paginator for
// GetAwsNetworkPerformanceData
type GetAwsNetworkPerformanceDataPaginator struct {
	options   GetAwsNetworkPerformanceDataPaginatorOptions
	client    GetAwsNetworkPerformanceDataAPIClient
	params    *GetAwsNetworkPerformanceDataInput
	nextToken *string
	firstPage bool
}

// NewGetAwsNetworkPerformanceDataPaginator returns a new
// GetAwsNetworkPerformanceDataPaginator
func NewGetAwsNetworkPerformanceDataPaginator(client GetAwsNetworkPerformanceDataAPIClient, params *GetAwsNetworkPerformanceDataInput, optFns ...func(*GetAwsNetworkPerformanceDataPaginatorOptions)) *GetAwsNetworkPerformanceDataPaginator {
	if params == nil {
		params = &GetAwsNetworkPerformanceDataInput{}
	}

	options := GetAwsNetworkPerformanceDataPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetAwsNetworkPerformanceDataPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetAwsNetworkPerformanceDataPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetAwsNetworkPerformanceData page.
func (p *GetAwsNetworkPerformanceDataPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetAwsNetworkPerformanceDataOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.GetAwsNetworkPerformanceData(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// GetAwsNetworkPerformanceDataAPIClient is a client that implements the
// GetAwsNetworkPerformanceData operation.
type GetAwsNetworkPerformanceDataAPIClient interface {
	GetAwsNetworkPerformanceData(context.Context, *GetAwsNetworkPerformanceDataInput, ...func(*Options)) (*GetAwsNetworkPerformanceDataOutput, error)
}

var _ GetAwsNetworkPerformanceDataAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opGetAwsNetworkPerformanceData(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetAwsNetworkPerformanceData",
	}
}
