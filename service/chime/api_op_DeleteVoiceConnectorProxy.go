// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

type DeleteVoiceConnectorProxyInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Chime Voice Connector ID.
	//
	// VoiceConnectorId is a required field
	VoiceConnectorId *string `location:"uri" locationName:"voiceConnectorId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteVoiceConnectorProxyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteVoiceConnectorProxyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteVoiceConnectorProxyInput"}

	if s.VoiceConnectorId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VoiceConnectorId"))
	}
	if s.VoiceConnectorId != nil && len(*s.VoiceConnectorId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("VoiceConnectorId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteVoiceConnectorProxyInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.VoiceConnectorId != nil {
		v := *s.VoiceConnectorId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "voiceConnectorId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteVoiceConnectorProxyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteVoiceConnectorProxyOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteVoiceConnectorProxyOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteVoiceConnectorProxy = "DeleteVoiceConnectorProxy"

// DeleteVoiceConnectorProxyRequest returns a request value for making API operation for
// Amazon Chime.
//
// Deletes the proxy configuration from the specified Amazon Chime Voice Connector.
//
//    // Example sending a request using DeleteVoiceConnectorProxyRequest.
//    req := client.DeleteVoiceConnectorProxyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/DeleteVoiceConnectorProxy
func (c *Client) DeleteVoiceConnectorProxyRequest(input *DeleteVoiceConnectorProxyInput) DeleteVoiceConnectorProxyRequest {
	op := &aws.Operation{
		Name:       opDeleteVoiceConnectorProxy,
		HTTPMethod: "DELETE",
		HTTPPath:   "/voice-connectors/{voiceConnectorId}/programmable-numbers/proxy",
	}

	if input == nil {
		input = &DeleteVoiceConnectorProxyInput{}
	}

	req := c.newRequest(op, input, &DeleteVoiceConnectorProxyOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteVoiceConnectorProxyRequest{Request: req, Input: input, Copy: c.DeleteVoiceConnectorProxyRequest}
}

// DeleteVoiceConnectorProxyRequest is the request type for the
// DeleteVoiceConnectorProxy API operation.
type DeleteVoiceConnectorProxyRequest struct {
	*aws.Request
	Input *DeleteVoiceConnectorProxyInput
	Copy  func(*DeleteVoiceConnectorProxyInput) DeleteVoiceConnectorProxyRequest
}

// Send marshals and sends the DeleteVoiceConnectorProxy API request.
func (r DeleteVoiceConnectorProxyRequest) Send(ctx context.Context) (*DeleteVoiceConnectorProxyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteVoiceConnectorProxyResponse{
		DeleteVoiceConnectorProxyOutput: r.Request.Data.(*DeleteVoiceConnectorProxyOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteVoiceConnectorProxyResponse is the response type for the
// DeleteVoiceConnectorProxy API operation.
type DeleteVoiceConnectorProxyResponse struct {
	*DeleteVoiceConnectorProxyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteVoiceConnectorProxy request.
func (r *DeleteVoiceConnectorProxyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}