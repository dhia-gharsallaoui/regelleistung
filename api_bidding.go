/*
IP RL BSP API

IP RL BSP API for participation in capacity/energy market tenders.

API version: 3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)


// BiddingApiService BiddingApi service
type BiddingApiService service

type ApiDeleteBidRequest struct {
	ctx context.Context
	ApiService *BiddingApiService
	id string
}

func (r ApiDeleteBidRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteBidExecute(r)
}

/*
DeleteBid Deletes a simple, complex bid or a bid component

This operation can be used to delete
  * a specific simple bid (with the unique ID of the bid),
  * a single component of a complex bid (with the unique ID of the bid component) or
  * the entire complex bid with all its components (with the group ID of the complex bid).


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Bid identification
 @return ApiDeleteBidRequest
*/
func (a *BiddingApiService) DeleteBid(ctx context.Context, id string) ApiDeleteBidRequest {
	return ApiDeleteBidRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *BiddingApiService) DeleteBidExecute(r ApiDeleteBidRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BiddingApiService.DeleteBid")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/bids/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetBidRequest struct {
	ctx context.Context
	ApiService *BiddingApiService
	id string
}

func (r ApiGetBidRequest) Execute() (*Bid, *http.Response, error) {
	return r.ApiService.GetBidExecute(r)
}

/*
GetBid Get a simple, complex bid or a complex bid component

Returns the single bid with requested ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Unique bid identification or group bid id of complex bid
 @return ApiGetBidRequest
*/
func (a *BiddingApiService) GetBid(ctx context.Context, id string) ApiGetBidRequest {
	return ApiGetBidRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return Bid
func (a *BiddingApiService) GetBidExecute(r ApiGetBidRequest) (*Bid, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Bid
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BiddingApiService.GetBid")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/bids/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetBidsRequest struct {
	ctx context.Context
	ApiService *BiddingApiService
	deliveryDate *string
	market *ReserveMarket
	tender *string
	productType *ProductType
	productName *ProductName
	state *string
	zone *string
	from *time.Time
	to *time.Time
	type_ *BidType
	cid *string
	gid *string
	link *string
	linkedTo *string
	tag *string
	backupsOnly *bool
	backupFor *string
	offset *int32
	limit *int32
}

// Only bids for the specified delivery date will be listed (ISO 8601 format YYYY-MM-DD). If not specified and also no tender ID is specified, then only bids for today&#39;s delivery day will be listed.
func (r ApiGetBidsRequest) DeliveryDate(deliveryDate string) ApiGetBidsRequest {
	r.deliveryDate = &deliveryDate
	return r
}

// Market filter. If not specified and also no tender ID is specified, then only bids for capacity market will be listed.
func (r ApiGetBidsRequest) Market(market ReserveMarket) ApiGetBidsRequest {
	r.market = &market
	return r
}

// By specifying the tender ID, only bids of this tender will be listed. For details on the tender id format, see the [reference guide.](/docs/guide#tender-id)
func (r ApiGetBidsRequest) Tender(tender string) ApiGetBidsRequest {
	r.tender = &tender
	return r
}

// Product type filter. If this query parameter is not set and also no tender ID is specified, bids for all product types are returned.
func (r ApiGetBidsRequest) ProductType(productType ProductType) ApiGetBidsRequest {
	r.productType = &productType
	return r
}

// Product name filter. If this query parameter is not set, bids for all product names are returned.
func (r ApiGetBidsRequest) ProductName(productName ProductName) ApiGetBidsRequest {
	r.productName = &productName
	return r
}

func (r ApiGetBidsRequest) State(state string) ApiGetBidsRequest {
	r.state = &state
	return r
}

func (r ApiGetBidsRequest) Zone(zone string) ApiGetBidsRequest {
	r.zone = &zone
	return r
}

func (r ApiGetBidsRequest) From(from time.Time) ApiGetBidsRequest {
	r.from = &from
	return r
}

func (r ApiGetBidsRequest) To(to time.Time) ApiGetBidsRequest {
	r.to = &to
	return r
}

// By specifying the type, only bids of this type will be listed.
func (r ApiGetBidsRequest) Type_(type_ BidType) ApiGetBidsRequest {
	r.type_ = &type_
	return r
}

// Filters bids that have the given complex bid ID.
func (r ApiGetBidsRequest) Cid(cid string) ApiGetBidsRequest {
	r.cid = &cid
	return r
}

// Filters bids that have the given bid group ID.
func (r ApiGetBidsRequest) Gid(gid string) ApiGetBidsRequest {
	r.gid = &gid
	return r
}

// Filters bids that have the given technical link ID.
func (r ApiGetBidsRequest) Link(link string) ApiGetBidsRequest {
	r.link = &link
	return r
}

// Filters bids that have a conditional link to the given bid ID.
func (r ApiGetBidsRequest) LinkedTo(linkedTo string) ApiGetBidsRequest {
	r.linkedTo = &linkedTo
	return r
}

// Filters bids that have the given tag.
func (r ApiGetBidsRequest) Tag(tag string) ApiGetBidsRequest {
	r.tag = &tag
	return r
}

func (r ApiGetBidsRequest) BackupsOnly(backupsOnly bool) ApiGetBidsRequest {
	r.backupsOnly = &backupsOnly
	return r
}

// Filters bids which are a backup for the provider with the given EIC
func (r ApiGetBidsRequest) BackupFor(backupFor string) ApiGetBidsRequest {
	r.backupFor = &backupFor
	return r
}

// The number of items to skip before starting to collect the result set.
func (r ApiGetBidsRequest) Offset(offset int32) ApiGetBidsRequest {
	r.offset = &offset
	return r
}

// The numbers of items to return
func (r ApiGetBidsRequest) Limit(limit int32) ApiGetBidsRequest {
	r.limit = &limit
	return r
}

func (r ApiGetBidsRequest) Execute() ([]Bid, *http.Response, error) {
	return r.ApiService.GetBidsExecute(r)
}

/*
GetBids List submitted bids (optionally capacity or energy market)

Returns the list of all submitted capacity/energy bids and complex bid components sorted in the following order:
  * `productType`,
  * `productName`,
  * `state`,
  * `revisionTimestamp`,
  * `id`.

The result list can be reduced as required by specifying combinations of different filter parameters.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetBidsRequest
*/
func (a *BiddingApiService) GetBids(ctx context.Context) ApiGetBidsRequest {
	return ApiGetBidsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []Bid
func (a *BiddingApiService) GetBidsExecute(r ApiGetBidsRequest) ([]Bid, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Bid
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BiddingApiService.GetBids")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/bids"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.deliveryDate != nil {
		parameterAddToQuery(localVarQueryParams, "delivery-date", r.deliveryDate, "")
	}
	if r.market != nil {
		parameterAddToQuery(localVarQueryParams, "market", r.market, "")
	}
	if r.tender != nil {
		parameterAddToQuery(localVarQueryParams, "tender", r.tender, "")
	}
	if r.productType != nil {
		parameterAddToQuery(localVarQueryParams, "product-type", r.productType, "")
	}
	if r.productName != nil {
		parameterAddToQuery(localVarQueryParams, "product-name", r.productName, "")
	}
	if r.state != nil {
		parameterAddToQuery(localVarQueryParams, "state", r.state, "")
	}
	if r.zone != nil {
		parameterAddToQuery(localVarQueryParams, "zone", r.zone, "")
	}
	if r.from != nil {
		parameterAddToQuery(localVarQueryParams, "from", r.from, "")
	}
	if r.to != nil {
		parameterAddToQuery(localVarQueryParams, "to", r.to, "")
	}
	if r.type_ != nil {
		parameterAddToQuery(localVarQueryParams, "type", r.type_, "")
	}
	if r.cid != nil {
		parameterAddToQuery(localVarQueryParams, "cid", r.cid, "")
	}
	if r.gid != nil {
		parameterAddToQuery(localVarQueryParams, "gid", r.gid, "")
	}
	if r.link != nil {
		parameterAddToQuery(localVarQueryParams, "link", r.link, "")
	}
	if r.linkedTo != nil {
		parameterAddToQuery(localVarQueryParams, "linked-to", r.linkedTo, "")
	}
	if r.tag != nil {
		parameterAddToQuery(localVarQueryParams, "tag", r.tag, "")
	}
	if r.backupsOnly != nil {
		parameterAddToQuery(localVarQueryParams, "backups-only", r.backupsOnly, "")
	}
	if r.backupFor != nil {
		parameterAddToQuery(localVarQueryParams, "backup-for", r.backupFor, "")
	}
	if r.offset != nil {
		parameterAddToQuery(localVarQueryParams, "offset", r.offset, "")
	}
	if r.limit != nil {
		parameterAddToQuery(localVarQueryParams, "limit", r.limit, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiSubmitBidRequest struct {
	ctx context.Context
	ApiService *BiddingApiService
	bidSubmission *BidSubmission
}

func (r ApiSubmitBidRequest) BidSubmission(bidSubmission BidSubmission) ApiSubmitBidRequest {
	r.bidSubmission = &bidSubmission
	return r
}

func (r ApiSubmitBidRequest) Execute() (*BidSubmissionResponse, *http.Response, error) {
	return r.ApiService.SubmitBidExecute(r)
}

/*
SubmitBid Submit a simple, complex bid or a complex bid component

Submission of a simple or complex bid for a open capacity/energy market tender. By specifying the product, delivery date and market, the bid is automatically assigned to the corresponding open tender. The gate closure time of the tender is relevant for the submission.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSubmitBidRequest
*/
func (a *BiddingApiService) SubmitBid(ctx context.Context) ApiSubmitBidRequest {
	return ApiSubmitBidRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BidSubmissionResponse
func (a *BiddingApiService) SubmitBidExecute(r ApiSubmitBidRequest) (*BidSubmissionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BidSubmissionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BiddingApiService.SubmitBid")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/bids"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.bidSubmission
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdateBidRequest struct {
	ctx context.Context
	ApiService *BiddingApiService
	id string
	bidModification *BidModification
}

func (r ApiUpdateBidRequest) BidModification(bidModification BidModification) ApiUpdateBidRequest {
	r.bidModification = &bidModification
	return r
}

func (r ApiUpdateBidRequest) Execute() (*Bid, *http.Response, error) {
	return r.ApiService.UpdateBidExecute(r)
}

/*
UpdateBid Updates a simple, complex bid or a bid component

Connecting zone, volumes and prices can be changed (only applies to simple bids and individual components of a complex bid). Furthermore, linkages can be defined. The validation rules apply as for bid submission. It is possible to only include necessary information (change) in the request body, thus omitting bid attributes which are not to be changed.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Bid identification
 @return ApiUpdateBidRequest
*/
func (a *BiddingApiService) UpdateBid(ctx context.Context, id string) ApiUpdateBidRequest {
	return ApiUpdateBidRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return Bid
func (a *BiddingApiService) UpdateBidExecute(r ApiUpdateBidRequest) (*Bid, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Bid
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BiddingApiService.UpdateBid")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/bids/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.bidModification
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
