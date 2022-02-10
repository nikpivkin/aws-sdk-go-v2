// Code generated by smithy-go-codegen DO NOT EDIT.

package restxml

import (
	"bytes"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	awshttp "github.com/aws/aws-sdk-go-v2/aws/transport/http"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/restxml/types"
	smithydocument "github.com/aws/smithy-go/document"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyrand "github.com/aws/smithy-go/rand"
	smithytesting "github.com/aws/smithy-go/testing"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
)

func TestClient_XmlUnions_awsRestxmlSerialize(t *testing.T) {
	cases := map[string]struct {
		Params        *XmlUnionsInput
		ExpectMethod  string
		ExpectURIPath string
		ExpectQuery   []smithytesting.QueryItem
		RequireQuery  []string
		ForbidQuery   []string
		ExpectHeader  http.Header
		RequireHeader []string
		ForbidHeader  []string
		Host          *url.URL
		BodyMediaType string
		BodyAssert    func(io.Reader) error
	}{
		// Serializes union struct member
		"XmlUnionsWithStructMember": {
			Params: &XmlUnionsInput{
				UnionValue: &types.XmlUnionShapeMemberStructValue{Value: types.XmlNestedUnionStruct{
					StringValue:  ptr.String("string"),
					BooleanValue: ptr.Bool(true),
					ByteValue:    ptr.Int8(1),
					ShortValue:   ptr.Int16(2),
					IntegerValue: ptr.Int32(3),
					LongValue:    ptr.Int64(4),
					FloatValue:   ptr.Float32(5.5),
					DoubleValue:  ptr.Float64(6.5),
				}},
			},
			ExpectMethod:  "PUT",
			ExpectURIPath: "/XmlUnions",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/xml"},
			},
			BodyMediaType: "application/xml",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareXMLReaderBytes(actual, []byte(`<XmlUnionsInputOutput>
			    <unionValue>
			       <structValue>
			          <stringValue>string</stringValue>
			          <booleanValue>true</booleanValue>
			          <byteValue>1</byteValue>
			          <shortValue>2</shortValue>
			          <integerValue>3</integerValue>
			          <longValue>4</longValue>
			          <floatValue>5.5</floatValue>
			          <doubleValue>6.5</doubleValue>
			       </structValue>
			    </unionValue>
			</XmlUnionsInputOutput>
			`))
			},
		},
		// serialize union string member
		"XmlUnionsWithStringMember": {
			Params: &XmlUnionsInput{
				UnionValue: &types.XmlUnionShapeMemberStringValue{Value: "some string"},
			},
			ExpectMethod:  "PUT",
			ExpectURIPath: "/XmlUnions",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/xml"},
			},
			BodyMediaType: "application/xml",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareXMLReaderBytes(actual, []byte(`<XmlUnionsInputOutput>
			   <unionValue>
			      <stringValue>some string</stringValue>
			   </unionValue>
			</XmlUnionsInputOutput>
			`))
			},
		},
		// Serializes union boolean member
		"XmlUnionsWithBooleanMember": {
			Params: &XmlUnionsInput{
				UnionValue: &types.XmlUnionShapeMemberBooleanValue{Value: true},
			},
			ExpectMethod:  "PUT",
			ExpectURIPath: "/XmlUnions",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/xml"},
			},
			BodyMediaType: "application/xml",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareXMLReaderBytes(actual, []byte(`<XmlUnionsInputOutput>
			   <unionValue>
			      <booleanValue>true</booleanValue>
			   </unionValue>
			</XmlUnionsInputOutput>
			`))
			},
		},
		// Serializes union member
		"XmlUnionsWithUnionMember": {
			Params: &XmlUnionsInput{
				UnionValue: &types.XmlUnionShapeMemberUnionValue{Value: &types.XmlUnionShapeMemberBooleanValue{Value: true}},
			},
			ExpectMethod:  "PUT",
			ExpectURIPath: "/XmlUnions",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/xml"},
			},
			BodyMediaType: "application/xml",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareXMLReaderBytes(actual, []byte(`<XmlUnionsInputOutput>
			   <unionValue>
			      <unionValue>
			         <booleanValue>true</booleanValue>
			      </unionValue>
			   </unionValue>
			</XmlUnionsInputOutput>
			`))
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			var actualReq *http.Request
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				actualReq = r.Clone(r.Context())
				if len(actualReq.URL.RawPath) == 0 {
					actualReq.URL.RawPath = actualReq.URL.Path
				}
				if v := actualReq.ContentLength; v != 0 {
					actualReq.Header.Set("Content-Length", strconv.FormatInt(v, 10))
				}
				var buf bytes.Buffer
				if _, err := io.Copy(&buf, r.Body); err != nil {
					t.Errorf("failed to read request body, %v", err)
				}
				actualReq.Body = ioutil.NopCloser(&buf)

				w.WriteHeader(200)
			}))
			defer server.Close()
			serverURL := server.URL
			if c.Host != nil {
				u, err := url.Parse(serverURL)
				if err != nil {
					t.Fatalf("expect no error, got %v", err)
				}
				u.Path = c.Host.Path
				u.RawPath = c.Host.RawPath
				u.RawQuery = c.Host.RawQuery
				serverURL = u.String()
			}
			client := New(Options{
				APIOptions: []func(*middleware.Stack) error{
					func(s *middleware.Stack) error {
						s.Finalize.Clear()
						s.Initialize.Remove(`OperationInputValidation`)
						return nil
					},
				},
				EndpointResolver: EndpointResolverFunc(func(region string, options EndpointResolverOptions) (e aws.Endpoint, err error) {
					e.URL = serverURL
					e.SigningRegion = "us-west-2"
					return e, err
				}),
				HTTPClient:               awshttp.NewBuildableClient(),
				IdempotencyTokenProvider: smithyrand.NewUUIDIdempotencyToken(&smithytesting.ByteLoop{}),
				Region:                   "us-west-2",
			})
			result, err := client.XmlUnions(context.Background(), c.Params)
			if err != nil {
				t.Fatalf("expect nil err, got %v", err)
			}
			if result == nil {
				t.Fatalf("expect not nil result")
			}
			if e, a := c.ExpectMethod, actualReq.Method; e != a {
				t.Errorf("expect %v method, got %v", e, a)
			}
			if e, a := c.ExpectURIPath, actualReq.URL.RawPath; e != a {
				t.Errorf("expect %v path, got %v", e, a)
			}
			queryItems := smithytesting.ParseRawQuery(actualReq.URL.RawQuery)
			smithytesting.AssertHasQuery(t, c.ExpectQuery, queryItems)
			smithytesting.AssertHasQueryKeys(t, c.RequireQuery, queryItems)
			smithytesting.AssertNotHaveQueryKeys(t, c.ForbidQuery, queryItems)
			smithytesting.AssertHasHeader(t, c.ExpectHeader, actualReq.Header)
			smithytesting.AssertHasHeaderKeys(t, c.RequireHeader, actualReq.Header)
			smithytesting.AssertNotHaveHeaderKeys(t, c.ForbidHeader, actualReq.Header)
			if c.BodyAssert != nil {
				if err := c.BodyAssert(actualReq.Body); err != nil {
					t.Errorf("expect body equal, got %v", err)
				}
			}
		})
	}
}

func TestClient_XmlUnions_awsRestxmlDeserialize(t *testing.T) {
	cases := map[string]struct {
		StatusCode    int
		Header        http.Header
		BodyMediaType string
		Body          []byte
		ExpectResult  *XmlUnionsOutput
	}{
		// Serializes union struct member
		"XmlUnionsWithStructMember": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/xml"},
			},
			BodyMediaType: "application/xml",
			Body: []byte(`<XmlUnionsInputOutput>
			    <unionValue>
			       <structValue>
			          <stringValue>string</stringValue>
			          <booleanValue>true</booleanValue>
			          <byteValue>1</byteValue>
			          <shortValue>2</shortValue>
			          <integerValue>3</integerValue>
			          <longValue>4</longValue>
			          <floatValue>5.5</floatValue>
			          <doubleValue>6.5</doubleValue>
			       </structValue>
			    </unionValue>
			</XmlUnionsInputOutput>
			`),
			ExpectResult: &XmlUnionsOutput{
				UnionValue: &types.XmlUnionShapeMemberStructValue{Value: types.XmlNestedUnionStruct{
					StringValue:  ptr.String("string"),
					BooleanValue: ptr.Bool(true),
					ByteValue:    ptr.Int8(1),
					ShortValue:   ptr.Int16(2),
					IntegerValue: ptr.Int32(3),
					LongValue:    ptr.Int64(4),
					FloatValue:   ptr.Float32(5.5),
					DoubleValue:  ptr.Float64(6.5),
				}},
			},
		},
		// Serializes union string member
		"XmlUnionsWithStringMember": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/xml"},
			},
			BodyMediaType: "application/xml",
			Body: []byte(`<XmlUnionsInputOutput>
			   <unionValue>
			      <stringValue>some string</stringValue>
			   </unionValue>
			</XmlUnionsInputOutput>
			`),
			ExpectResult: &XmlUnionsOutput{
				UnionValue: &types.XmlUnionShapeMemberStringValue{Value: "some string"},
			},
		},
		// Serializes union boolean member
		"XmlUnionsWithBooleanMember": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/xml"},
			},
			BodyMediaType: "application/xml",
			Body: []byte(`<XmlUnionsInputOutput>
			   <unionValue>
			      <booleanValue>true</booleanValue>
			   </unionValue>
			</XmlUnionsInputOutput>
			`),
			ExpectResult: &XmlUnionsOutput{
				UnionValue: &types.XmlUnionShapeMemberBooleanValue{Value: true},
			},
		},
		// Serializes union member
		"XmlUnionsWithUnionMember": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/xml"},
			},
			BodyMediaType: "application/xml",
			Body: []byte(`<XmlUnionsInputOutput>
			   <unionValue>
			      <unionValue>
			         <booleanValue>true</booleanValue>
			      </unionValue>
			   </unionValue>
			</XmlUnionsInputOutput>
			`),
			ExpectResult: &XmlUnionsOutput{
				UnionValue: &types.XmlUnionShapeMemberUnionValue{Value: &types.XmlUnionShapeMemberBooleanValue{Value: true}},
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			serverURL := "http://localhost:8888/"
			client := New(Options{
				HTTPClient: smithyhttp.ClientDoFunc(func(r *http.Request) (*http.Response, error) {
					headers := http.Header{}
					for k, vs := range c.Header {
						for _, v := range vs {
							headers.Add(k, v)
						}
					}
					if len(c.BodyMediaType) != 0 && len(headers.Values("Content-Type")) == 0 {
						headers.Set("Content-Type", c.BodyMediaType)
					}
					response := &http.Response{
						StatusCode: c.StatusCode,
						Header:     headers,
						Request:    r,
					}
					if len(c.Body) != 0 {
						response.ContentLength = int64(len(c.Body))
						response.Body = ioutil.NopCloser(bytes.NewReader(c.Body))
					} else {

						response.Body = http.NoBody
					}
					return response, nil
				}),
				APIOptions: []func(*middleware.Stack) error{
					func(s *middleware.Stack) error {
						s.Finalize.Clear()
						s.Initialize.Remove(`OperationInputValidation`)
						return nil
					},
				},
				EndpointResolver: EndpointResolverFunc(func(region string, options EndpointResolverOptions) (e aws.Endpoint, err error) {
					e.URL = serverURL
					e.SigningRegion = "us-west-2"
					return e, err
				}),
				IdempotencyTokenProvider: smithyrand.NewUUIDIdempotencyToken(&smithytesting.ByteLoop{}),
				Region:                   "us-west-2",
			})
			var params XmlUnionsInput
			result, err := client.XmlUnions(context.Background(), &params)
			if err != nil {
				t.Fatalf("expect nil err, got %v", err)
			}
			if result == nil {
				t.Fatalf("expect not nil result")
			}
			opts := cmp.Options{
				cmpopts.IgnoreUnexported(
					middleware.Metadata{},
				),
				cmp.FilterValues(func(x, y float64) bool {
					return math.IsNaN(x) && math.IsNaN(y)
				}, cmp.Comparer(func(_, _ interface{}) bool { return true })),
				cmp.FilterValues(func(x, y float32) bool {
					return math.IsNaN(float64(x)) && math.IsNaN(float64(y))
				}, cmp.Comparer(func(_, _ interface{}) bool { return true })),
				cmpopts.IgnoreTypes(smithydocument.NoSerde{}),
			}
			if err := smithytesting.CompareValues(c.ExpectResult, result, opts...); err != nil {
				t.Errorf("expect c.ExpectResult value match:\n%v", err)
			}
		})
	}
}
