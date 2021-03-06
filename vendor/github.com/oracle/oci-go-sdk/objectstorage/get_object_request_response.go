// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package objectstorage

import (
	"github.com/oracle/oci-go-sdk/common"
	"io"
	"net/http"
)

// GetObjectRequest wrapper for the GetObject operation
type GetObjectRequest struct {

	// The top-level namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The name of the bucket. Avoid entering confidential information.
	// Example: `my-new-bucket1`
	BucketName *string `mandatory:"true" contributesTo:"path" name:"bucketName"`

	// The name of the object. Avoid entering confidential information.
	// Example: `test/object1.log`
	ObjectName *string `mandatory:"true" contributesTo:"path" name:"objectName"`

	// The entity tag to match. For creating and committing a multipart upload to an object, this is the entity tag of the target object.
	// For uploading a part, this is the entity tag of the target part.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// The entity tag to avoid matching. The only valid value is ‘*’, which indicates that the request should fail if the object already exists.
	// For creating and committing a multipart upload, this is the entity tag of the target object. For uploading a part, this is the entity tag of the target part.
	IfNoneMatch *string `mandatory:"false" contributesTo:"header" name:"if-none-match"`

	// The client request ID for tracing.
	OpcClientRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-client-request-id"`

	// Optional byte range to fetch, as described in RFC 7233 (https://tools.ietf.org/rfc/rfc7233), section 2.1.
	// Note that only a single range of bytes is supported.
	Range *string `mandatory:"false" contributesTo:"header" name:"range"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetObjectRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetObjectRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetObjectRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetObjectResponse wrapper for the GetObject operation
type GetObjectResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The io.ReadCloser instance
	Content io.ReadCloser `presentIn:"body" encoding:"binary"`

	// Echoes back the value passed in the opc-client-request-id header, for use by clients when debugging.
	OpcClientRequestId *string `presentIn:"header" name:"opc-client-request-id"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular
	// request, provide this request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// The entity tag for the object.
	ETag *string `presentIn:"header" name:"etag"`

	// The user-defined metadata for the object.
	OpcMeta map[string]string `presentIn:"header-collection" prefix:"opc-meta-"`

	// The object size in bytes.
	ContentLength *int `presentIn:"header" name:"content-length"`

	// Content-Range header for range requests, per RFC 7233 (https://tools.ietf.org/rfc/rfc7233), section 4.2.
	ContentRange *string `presentIn:"header" name:"content-range"`

	// Content-MD5 header, as described in RFC 2616 (https://tools.ietf.org/rfc/rfc2616), section 14.15.
	// Unavailable for objects uploaded using multipart upload.
	ContentMd5 *string `presentIn:"header" name:"content-md5"`

	// Only applicable to objects uploaded using multipart upload.
	// Base-64 representation of the multipart object hash.
	// The multipart object hash is calculated by taking the MD5 hashes of the parts,
	// concatenating the binary representation of those hashes in order of their part numbers,
	// and then calculating the MD5 hash of the concatenated values.
	OpcMultipartMd5 *string `presentIn:"header" name:"opc-multipart-md5"`

	// Content-Type header, as described in RFC 2616 (https://tools.ietf.org/rfc/rfc2616), section 14.17.
	ContentType *string `presentIn:"header" name:"content-type"`

	// Content-Language header, as described in RFC 2616 (https://tools.ietf.org/rfc/rfc2616), section 14.12.
	ContentLanguage *string `presentIn:"header" name:"content-language"`

	// Content-Encoding header, as described in RFC 2616 (https://tools.ietf.org/rfc/rfc2616), section 14.11.
	ContentEncoding *string `presentIn:"header" name:"content-encoding"`

	// The object modification time, as described in RFC 2616 (https://tools.ietf.org/rfc/rfc2616), section 14.29.
	LastModified *common.SDKTime `presentIn:"header" name:"last-modified"`

	// The current state of the object.
	ArchivalState GetObjectArchivalStateEnum `presentIn:"header" name:"archival-state"`

	// Time that the object is returned to the archived state. This field is only present for restored objects.
	TimeOfArchival *common.SDKTime `presentIn:"header" name:"time-of-archival"`

	// Flag to indicate whether or not the object was modified.  If this is true,
	// the getter for the object itself will return null.  Callers should check this
	// if they specified one of the request params that might result in a conditional
	// response (like 'if-match'/'if-none-match').
	IsNotModified bool
}

func (response GetObjectResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetObjectResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetObjectArchivalStateEnum Enum with underlying type: string
type GetObjectArchivalStateEnum string

// Set of constants representing the allowable values for GetObjectArchivalState
const (
	GetObjectArchivalStateAvailable GetObjectArchivalStateEnum = "AVAILABLE"
	GetObjectArchivalStateArchived  GetObjectArchivalStateEnum = "ARCHIVED"
	GetObjectArchivalStateRestoring GetObjectArchivalStateEnum = "RESTORING"
	GetObjectArchivalStateRestored  GetObjectArchivalStateEnum = "RESTORED"
)

var mappingGetObjectArchivalState = map[string]GetObjectArchivalStateEnum{
	"AVAILABLE": GetObjectArchivalStateAvailable,
	"ARCHIVED":  GetObjectArchivalStateArchived,
	"RESTORING": GetObjectArchivalStateRestoring,
	"RESTORED":  GetObjectArchivalStateRestored,
}

// GetGetObjectArchivalStateEnumValues Enumerates the set of values for GetObjectArchivalState
func GetGetObjectArchivalStateEnumValues() []GetObjectArchivalStateEnum {
	values := make([]GetObjectArchivalStateEnum, 0)
	for _, v := range mappingGetObjectArchivalState {
		values = append(values, v)
	}
	return values
}
