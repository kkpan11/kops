// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates or updates a scheduled scaling action for an Auto Scaling group.
//
// For more information, see [Scheduled scaling] in the Amazon EC2 Auto Scaling User Guide.
//
// You can view the scheduled actions for an Auto Scaling group using the DescribeScheduledActions API
// call. If you are no longer using a scheduled action, you can delete it by
// calling the DeleteScheduledActionAPI.
//
// If you try to schedule your action in the past, Amazon EC2 Auto Scaling returns
// an error message.
//
// [Scheduled scaling]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-scheduled-scaling.html
func (c *Client) PutScheduledUpdateGroupAction(ctx context.Context, params *PutScheduledUpdateGroupActionInput, optFns ...func(*Options)) (*PutScheduledUpdateGroupActionOutput, error) {
	if params == nil {
		params = &PutScheduledUpdateGroupActionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutScheduledUpdateGroupAction", params, optFns, c.addOperationPutScheduledUpdateGroupActionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutScheduledUpdateGroupActionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutScheduledUpdateGroupActionInput struct {

	// The name of the Auto Scaling group.
	//
	// This member is required.
	AutoScalingGroupName *string

	// The name of this scaling action.
	//
	// This member is required.
	ScheduledActionName *string

	// The desired capacity is the initial capacity of the Auto Scaling group after
	// the scheduled action runs and the capacity it attempts to maintain. It can scale
	// beyond this capacity if you add more scaling conditions.
	//
	// You must specify at least one of the following properties: MaxSize , MinSize ,
	// or DesiredCapacity .
	DesiredCapacity *int32

	// The date and time for the recurring schedule to end, in UTC. For example,
	// "2021-06-01T00:00:00Z" .
	EndTime *time.Time

	// The maximum size of the Auto Scaling group.
	MaxSize *int32

	// The minimum size of the Auto Scaling group.
	MinSize *int32

	// The recurring schedule for this action. This format consists of five fields
	// separated by white spaces: [Minute] [Hour] [Day_of_Month] [Month_of_Year]
	// [Day_of_Week]. The value must be in quotes (for example, "30 0 1 1,6,12 *" ).
	// For more information about this format, see [Crontab].
	//
	// When StartTime and EndTime are specified with Recurrence , they form the
	// boundaries of when the recurring action starts and stops.
	//
	// Cron expressions use Universal Coordinated Time (UTC) by default.
	//
	// [Crontab]: http://crontab.org
	Recurrence *string

	// The date and time for this action to start, in YYYY-MM-DDThh:mm:ssZ format in
	// UTC/GMT only and in quotes (for example, "2021-06-01T00:00:00Z" ).
	//
	// If you specify Recurrence and StartTime , Amazon EC2 Auto Scaling performs the
	// action at this time, and then performs the action based on the specified
	// recurrence.
	StartTime *time.Time

	// This property is no longer used.
	Time *time.Time

	// Specifies the time zone for a cron expression. If a time zone is not provided,
	// UTC is used by default.
	//
	// Valid values are the canonical names of the IANA time zones, derived from the
	// IANA Time Zone Database (such as Etc/GMT+9 or Pacific/Tahiti ). For more
	// information, see [https://en.wikipedia.org/wiki/List_of_tz_database_time_zones].
	//
	// [https://en.wikipedia.org/wiki/List_of_tz_database_time_zones]: https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
	TimeZone *string

	noSmithyDocumentSerde
}

type PutScheduledUpdateGroupActionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutScheduledUpdateGroupActionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpPutScheduledUpdateGroupAction{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpPutScheduledUpdateGroupAction{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutScheduledUpdateGroupAction"); err != nil {
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
	if err = addOpPutScheduledUpdateGroupActionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutScheduledUpdateGroupAction(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutScheduledUpdateGroupAction(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutScheduledUpdateGroupAction",
	}
}
