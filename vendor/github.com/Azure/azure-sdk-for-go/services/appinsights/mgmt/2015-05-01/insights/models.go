package insights

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"net/http"
)

// ApplicationType enumerates the values for application type.
type ApplicationType string

const (
	// Other ...
	Other ApplicationType = "other"
	// Web ...
	Web ApplicationType = "web"
)

// FlowType enumerates the values for flow type.
type FlowType string

const (
	// Bluefield ...
	Bluefield FlowType = "Bluefield"
)

// RequestSource enumerates the values for request source.
type RequestSource string

const (
	// Rest ...
	Rest RequestSource = "rest"
)

// WebTestKind enumerates the values for web test kind.
type WebTestKind string

const (
	// Multistep ...
	Multistep WebTestKind = "multistep"
	// Ping ...
	Ping WebTestKind = "ping"
)

// APIKeyRequest an Application Insights component API Key createion request definition.
type APIKeyRequest struct {
	// Name - The name of the API Key.
	Name *string `json:"name,omitempty"`
	// LinkedReadProperties - The read access rights of this API Key.
	LinkedReadProperties *[]string `json:"linkedReadProperties,omitempty"`
	// LinkedWriteProperties - The write access rights of this API Key.
	LinkedWriteProperties *[]string `json:"linkedWriteProperties,omitempty"`
}

// ApplicationInsightsComponent an Application Insights component definition.
type ApplicationInsightsComponent struct {
	autorest.Response `json:"-"`
	// ID - Azure resource Id
	ID *string `json:"id,omitempty"`
	// Name - Azure resource name
	Name *string `json:"name,omitempty"`
	// Type - Azure resource type
	Type *string `json:"type,omitempty"`
	// Location - Resource location
	Location *string `json:"location,omitempty"`
	// Tags - Resource tags
	Tags *map[string]*string `json:"tags,omitempty"`
	// Kind - The kind of application that this component refers to, used to customize UI. This value is a freeform string, values should typically be one of the following: web, ios, other, store, java, phone.
	Kind *string `json:"kind,omitempty"`
	// ApplicationInsightsComponentProperties - Properties that define an Application Insights component resource.
	*ApplicationInsightsComponentProperties `json:"properties,omitempty"`
}

// UnmarshalJSON is the custom unmarshaler for ApplicationInsightsComponent struct.
func (aic *ApplicationInsightsComponent) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	var v *json.RawMessage

	v = m["kind"]
	if v != nil {
		var kind string
		err = json.Unmarshal(*m["kind"], &kind)
		if err != nil {
			return err
		}
		aic.Kind = &kind
	}

	v = m["properties"]
	if v != nil {
		var properties ApplicationInsightsComponentProperties
		err = json.Unmarshal(*m["properties"], &properties)
		if err != nil {
			return err
		}
		aic.ApplicationInsightsComponentProperties = &properties
	}

	v = m["id"]
	if v != nil {
		var ID string
		err = json.Unmarshal(*m["id"], &ID)
		if err != nil {
			return err
		}
		aic.ID = &ID
	}

	v = m["name"]
	if v != nil {
		var name string
		err = json.Unmarshal(*m["name"], &name)
		if err != nil {
			return err
		}
		aic.Name = &name
	}

	v = m["type"]
	if v != nil {
		var typeVar string
		err = json.Unmarshal(*m["type"], &typeVar)
		if err != nil {
			return err
		}
		aic.Type = &typeVar
	}

	v = m["location"]
	if v != nil {
		var location string
		err = json.Unmarshal(*m["location"], &location)
		if err != nil {
			return err
		}
		aic.Location = &location
	}

	v = m["tags"]
	if v != nil {
		var tags map[string]*string
		err = json.Unmarshal(*m["tags"], &tags)
		if err != nil {
			return err
		}
		aic.Tags = &tags
	}

	return nil
}

// ApplicationInsightsComponentAPIKey properties that define an API key of an Application Insights Component.
type ApplicationInsightsComponentAPIKey struct {
	autorest.Response `json:"-"`
	// ID - The unique ID of the API key inside an Applciation Insights component. It is auto generated when the API key is created.
	ID *string `json:"id,omitempty"`
	// APIKey - The API key value. It will be only return once when the API Key was created.
	APIKey *string `json:"apiKey,omitempty"`
	// CreatedDate - The create date of this API key.
	CreatedDate *string `json:"createdDate,omitempty"`
	// Name - The name of the API key.
	Name *string `json:"name,omitempty"`
	// LinkedReadProperties - The read access rights of this API Key.
	LinkedReadProperties *[]string `json:"linkedReadProperties,omitempty"`
	// LinkedWriteProperties - The write access rights of this API Key.
	LinkedWriteProperties *[]string `json:"linkedWriteProperties,omitempty"`
}

// ApplicationInsightsComponentAPIKeyListResult describes the list of API Keys of an Application Insights Component.
type ApplicationInsightsComponentAPIKeyListResult struct {
	autorest.Response `json:"-"`
	// Value - List of API Key definitions.
	Value *[]ApplicationInsightsComponentAPIKey `json:"value,omitempty"`
}

// ApplicationInsightsComponentBillingFeatures an Application Insights component billing features
type ApplicationInsightsComponentBillingFeatures struct {
	autorest.Response `json:"-"`
	// DataVolumeCap - An Application Insights component daily data volumne cap
	DataVolumeCap *ApplicationInsightsComponentDataVolumeCap `json:"DataVolumeCap,omitempty"`
	// CurrentBillingFeatures - Current enabled pricing plan. When the component is in the Enterprise plan, this will list both 'Basic' and 'Application Insights Enterprise'.
	CurrentBillingFeatures *[]string `json:"CurrentBillingFeatures,omitempty"`
}

// ApplicationInsightsComponentDataVolumeCap an Application Insights component daily data volumne cap
type ApplicationInsightsComponentDataVolumeCap struct {
	// Cap - Daily data volume cap in GB.
	Cap *float64 `json:"Cap,omitempty"`
	// ResetTime - Daily data volume cap UTC reset hour.
	ResetTime *int32 `json:"ResetTime,omitempty"`
	// WarningThreshold - Reserved, not used for now.
	WarningThreshold *int32 `json:"WarningThreshold,omitempty"`
	// StopSendNotificationWhenHitThreshold - Reserved, not used for now.
	StopSendNotificationWhenHitThreshold *bool `json:"StopSendNotificationWhenHitThreshold,omitempty"`
	// StopSendNotificationWhenHitCap - Do not send a notification email when the daily data volume cap is met.
	StopSendNotificationWhenHitCap *bool `json:"StopSendNotificationWhenHitCap,omitempty"`
	// MaxHistoryCap - Maximum daily data volume cap that the user can set for this component.
	MaxHistoryCap *float64 `json:"MaxHistoryCap,omitempty"`
}

// ApplicationInsightsComponentExportConfiguration properties that define a Continuous Export configuration.
type ApplicationInsightsComponentExportConfiguration struct {
	autorest.Response `json:"-"`
	// ExportID - The unique ID of the export configuration inside an Applciation Insights component. It is auto generated when the Continuous Export configuration is created.
	ExportID *string `json:"ExportId,omitempty"`
	// InstrumentationKey - The instrumentation key of the Application Insights component.
	InstrumentationKey *string `json:"InstrumentationKey,omitempty"`
	// RecordTypes - This comma separated list of document types that will be exported. The possible values include 'Requests', 'Event', 'Exceptions', 'Metrics', 'PageViews', 'PageViewPerformance', 'Rdd', 'PerformanceCounters', 'Availability', 'Messages'.
	RecordTypes *string `json:"RecordTypes,omitempty"`
	// ApplicationName - The name of the Application Insights component.
	ApplicationName *string `json:"ApplicationName,omitempty"`
	// SubscriptionID - The subscription of the Application Insights component.
	SubscriptionID *string `json:"SubscriptionId,omitempty"`
	// ResourceGroup - The resource group of the Application Insights component.
	ResourceGroup *string `json:"ResourceGroup,omitempty"`
	// DestinationStorageSubscriptionID - The destination storage account subscription ID.
	DestinationStorageSubscriptionID *string `json:"DestinationStorageSubscriptionId,omitempty"`
	// DestinationStorageLocationID - The destination account location ID.
	DestinationStorageLocationID *string `json:"DestinationStorageLocationId,omitempty"`
	// DestinationAccountID - The name of destination account.
	DestinationAccountID *string `json:"DestinationAccountId,omitempty"`
	// DestinationType - The destination type.
	DestinationType *string `json:"DestinationType,omitempty"`
	// IsUserEnabled - This will be 'true' if the Continuous Export configuration is enabled, otherwise it will be 'false'.
	IsUserEnabled *string `json:"IsUserEnabled,omitempty"`
	// LastUserUpdate - Last time the Continuous Export configuration was updated.
	LastUserUpdate *string `json:"LastUserUpdate,omitempty"`
	// NotificationQueueEnabled - Deprecated
	NotificationQueueEnabled *string `json:"NotificationQueueEnabled,omitempty"`
	// ExportStatus - This indicates current Continuous Export configuration status. The possible values are 'Preparing', 'Success', 'Failure'.
	ExportStatus *string `json:"ExportStatus,omitempty"`
	// LastSuccessTime - The last time data was successfully delivered to the destination storage container for this Continuous Export configuration.
	LastSuccessTime *string `json:"LastSuccessTime,omitempty"`
	// LastGapTime - The last time the Continuous Export configuration started failing.
	LastGapTime *string `json:"LastGapTime,omitempty"`
	// PermanentErrorReason - This is the reason the Continuous Export configuration started failing. It can be 'AzureStorageNotFound' or 'AzureStorageAccessDenied'.
	PermanentErrorReason *string `json:"PermanentErrorReason,omitempty"`
	// StorageName - The name of the destination storage account.
	StorageName *string `json:"StorageName,omitempty"`
	// ContainerName - The name of the destination storage container.
	ContainerName *string `json:"ContainerName,omitempty"`
}

// ApplicationInsightsComponentExportRequest an Application Insights component Continuous Export configuration request
// definition.
type ApplicationInsightsComponentExportRequest struct {
	// RecordTypes - The document types to be exported, as comma separated values. Allowed values include 'Requests', 'Event', 'Exceptions', 'Metrics', 'PageViews', 'PageViewPerformance', 'Rdd', 'PerformanceCounters', 'Availability', 'Messages'.
	RecordTypes *string `json:"RecordTypes,omitempty"`
	// DestinationType - The Continuous Export destination type. This has to be 'Blob'.
	DestinationType *string `json:"DestinationType,omitempty"`
	// DestinationAddress - The SAS URL for the destination storage container. It must grant write permission.
	DestinationAddress *string `json:"DestinationAddress,omitempty"`
	// IsEnabled - Set to 'true' to create a Continuous Export configuration as enabled, otherwise set it to 'false'.
	IsEnabled *string `json:"IsEnabled,omitempty"`
	// NotificationQueueEnabled - Deprecated
	NotificationQueueEnabled *string `json:"NotificationQueueEnabled,omitempty"`
	// NotificationQueueURI - Deprecated
	NotificationQueueURI *string `json:"NotificationQueueUri,omitempty"`
	// DestinationStorageSubscriptionID - The subscription ID of the destination storage container.
	DestinationStorageSubscriptionID *string `json:"DestinationStorageSubscriptionId,omitempty"`
	// DestinationStorageLocationID - The location ID of the destination storage container.
	DestinationStorageLocationID *string `json:"DestinationStorageLocationId,omitempty"`
	// DestinationAccountID - The name of destination storage account.
	DestinationAccountID *string `json:"DestinationAccountId,omitempty"`
}

// ApplicationInsightsComponentListResult describes the list of Application Insights Resources.
type ApplicationInsightsComponentListResult struct {
	autorest.Response `json:"-"`
	// Value - List of Application Insights component definitions.
	Value *[]ApplicationInsightsComponent `json:"value,omitempty"`
	// NextLink - The URI to get the next set of Application Insights component defintions if too many components where returned in the result set.
	NextLink *string `json:"nextLink,omitempty"`
}

// ApplicationInsightsComponentListResultIterator provides access to a complete listing of ApplicationInsightsComponent
// values.
type ApplicationInsightsComponentListResultIterator struct {
	i    int
	page ApplicationInsightsComponentListResultPage
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *ApplicationInsightsComponentListResultIterator) Next() error {
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err := iter.page.Next()
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter ApplicationInsightsComponentListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter ApplicationInsightsComponentListResultIterator) Response() ApplicationInsightsComponentListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter ApplicationInsightsComponentListResultIterator) Value() ApplicationInsightsComponent {
	if !iter.page.NotDone() {
		return ApplicationInsightsComponent{}
	}
	return iter.page.Values()[iter.i]
}

// IsEmpty returns true if the ListResult contains no values.
func (aiclr ApplicationInsightsComponentListResult) IsEmpty() bool {
	return aiclr.Value == nil || len(*aiclr.Value) == 0
}

// applicationInsightsComponentListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (aiclr ApplicationInsightsComponentListResult) applicationInsightsComponentListResultPreparer() (*http.Request, error) {
	if aiclr.NextLink == nil || len(to.String(aiclr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(aiclr.NextLink)))
}

// ApplicationInsightsComponentListResultPage contains a page of ApplicationInsightsComponent values.
type ApplicationInsightsComponentListResultPage struct {
	fn    func(ApplicationInsightsComponentListResult) (ApplicationInsightsComponentListResult, error)
	aiclr ApplicationInsightsComponentListResult
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *ApplicationInsightsComponentListResultPage) Next() error {
	next, err := page.fn(page.aiclr)
	if err != nil {
		return err
	}
	page.aiclr = next
	return nil
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page ApplicationInsightsComponentListResultPage) NotDone() bool {
	return !page.aiclr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page ApplicationInsightsComponentListResultPage) Response() ApplicationInsightsComponentListResult {
	return page.aiclr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page ApplicationInsightsComponentListResultPage) Values() []ApplicationInsightsComponent {
	if page.aiclr.IsEmpty() {
		return nil
	}
	return *page.aiclr.Value
}

// ApplicationInsightsComponentProperties properties that define an Application Insights component resource.
type ApplicationInsightsComponentProperties struct {
	// ApplicationID - The unique ID of your application. This field mirrors the 'Name' field and cannot be changed.
	ApplicationID *string `json:"ApplicationId,omitempty"`
	// AppID - Application Insights Unique ID for your Application.
	AppID *string `json:"AppId,omitempty"`
	// ApplicationType - Type of application being monitored. Possible values include: 'Web', 'Other'
	ApplicationType ApplicationType `json:"Application_Type,omitempty"`
	// FlowType - Used by the Application Insights system to determine what kind of flow this component was created by. This is to be set to 'Bluefield' when creating/updating a component via the REST API. Possible values include: 'Bluefield'
	FlowType FlowType `json:"Flow_Type,omitempty"`
	// RequestSource - Describes what tool created this Application Insights component. Customers using this API should set this to the default 'rest'. Possible values include: 'Rest'
	RequestSource RequestSource `json:"Request_Source,omitempty"`
	// InstrumentationKey - Application Insights Instrumentation key. A read-only value that applications can use to identify the destination for all telemetry sent to Azure Application Insights. This value will be supplied upon construction of each new Application Insights component.
	InstrumentationKey *string `json:"InstrumentationKey,omitempty"`
	// CreationDate - Creation Date for the Application Insights component, in ISO 8601 format.
	CreationDate *date.Time `json:"CreationDate,omitempty"`
	// TenantID - Azure Tenant Id.
	TenantID *string `json:"TenantId,omitempty"`
	// HockeyAppID - The unique application ID created when a new application is added to HockeyApp, used for communications with HockeyApp.
	HockeyAppID *string `json:"HockeyAppId,omitempty"`
	// HockeyAppToken - Token used to authenticate communications with between Application Insights and HockeyApp.
	HockeyAppToken *string `json:"HockeyAppToken,omitempty"`
	// ProvisioningState - Current state of this component: whether or not is has been provisioned within the resource group it is defined. Users cannot change this value but are able to read from it. Values will include Succeeded, Deploying, Canceled, and Failed.
	ProvisioningState *string `json:"provisioningState,omitempty"`
	// SamplingPercentage - Percentage of the data produced by the application being monitored that is being sampled for Application Insights telemetry.
	SamplingPercentage *float64 `json:"SamplingPercentage,omitempty"`
}

// ApplicationInsightsComponentQuotaStatus an Application Insights component daily data volume cap status
type ApplicationInsightsComponentQuotaStatus struct {
	autorest.Response `json:"-"`
	// AppID - The Application ID for the Application Insights component.
	AppID *string `json:"AppId,omitempty"`
	// ShouldBeThrottled - The daily data volume cap is met, and data ingestion will be stopped.
	ShouldBeThrottled *bool `json:"ShouldBeThrottled,omitempty"`
	// ExpirationTime - Date and time when the daily data volume cap will be reset, and data ingestion will resume.
	ExpirationTime *string `json:"ExpirationTime,omitempty"`
}

// ErrorResponse error reponse indicates Insights service is not able to process the incoming request. The reason is
// provided in the error message.
type ErrorResponse struct {
	// Code - Error code.
	Code *string `json:"code,omitempty"`
	// Message - Error message indicating why the operation failed.
	Message *string `json:"message,omitempty"`
}

// ListApplicationInsightsComponentExportConfiguration ...
type ListApplicationInsightsComponentExportConfiguration struct {
	autorest.Response `json:"-"`
	Value             *[]ApplicationInsightsComponentExportConfiguration `json:"value,omitempty"`
}

// Operation CDN REST API operation
type Operation struct {
	// Name - Operation name: {provider}/{resource}/{operation}
	Name *string `json:"name,omitempty"`
	// Display - The object that represents the operation.
	Display *OperationDisplay `json:"display,omitempty"`
}

// OperationDisplay the object that represents the operation.
type OperationDisplay struct {
	// Provider - Service provider: Microsoft.Cdn
	Provider *string `json:"provider,omitempty"`
	// Resource - Resource on which the operation is performed: Profile, endpoint, etc.
	Resource *string `json:"resource,omitempty"`
	// Operation - Operation type: Read, write, delete, etc.
	Operation *string `json:"operation,omitempty"`
}

// OperationListResult result of the request to list CDN operations. It contains a list of operations and a URL link to
// get the next set of results.
type OperationListResult struct {
	autorest.Response `json:"-"`
	// Value - List of CDN operations supported by the CDN resource provider.
	Value *[]Operation `json:"value,omitempty"`
	// NextLink - URL to get the next set of operation list results if there are any.
	NextLink *string `json:"nextLink,omitempty"`
}

// OperationListResultIterator provides access to a complete listing of Operation values.
type OperationListResultIterator struct {
	i    int
	page OperationListResultPage
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *OperationListResultIterator) Next() error {
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err := iter.page.Next()
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter OperationListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter OperationListResultIterator) Response() OperationListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter OperationListResultIterator) Value() Operation {
	if !iter.page.NotDone() {
		return Operation{}
	}
	return iter.page.Values()[iter.i]
}

// IsEmpty returns true if the ListResult contains no values.
func (olr OperationListResult) IsEmpty() bool {
	return olr.Value == nil || len(*olr.Value) == 0
}

// operationListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (olr OperationListResult) operationListResultPreparer() (*http.Request, error) {
	if olr.NextLink == nil || len(to.String(olr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(olr.NextLink)))
}

// OperationListResultPage contains a page of Operation values.
type OperationListResultPage struct {
	fn  func(OperationListResult) (OperationListResult, error)
	olr OperationListResult
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *OperationListResultPage) Next() error {
	next, err := page.fn(page.olr)
	if err != nil {
		return err
	}
	page.olr = next
	return nil
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page OperationListResultPage) NotDone() bool {
	return !page.olr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page OperationListResultPage) Response() OperationListResult {
	return page.olr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page OperationListResultPage) Values() []Operation {
	if page.olr.IsEmpty() {
		return nil
	}
	return *page.olr.Value
}

// Resource an azure resource object
type Resource struct {
	// ID - Azure resource Id
	ID *string `json:"id,omitempty"`
	// Name - Azure resource name
	Name *string `json:"name,omitempty"`
	// Type - Azure resource type
	Type *string `json:"type,omitempty"`
	// Location - Resource location
	Location *string `json:"location,omitempty"`
	// Tags - Resource tags
	Tags *map[string]*string `json:"tags,omitempty"`
}

// TagsResource a container holding only the Tags for a resource, allowing the user to update the tags on a WebTest
// instance.
type TagsResource struct {
	// Tags - Resource tags
	Tags *map[string]*string `json:"tags,omitempty"`
}

// WebTest an Application Insights web test definition.
type WebTest struct {
	autorest.Response `json:"-"`
	// ID - Azure resource Id
	ID *string `json:"id,omitempty"`
	// Name - Azure resource name
	Name *string `json:"name,omitempty"`
	// Type - Azure resource type
	Type *string `json:"type,omitempty"`
	// Location - Resource location
	Location *string `json:"location,omitempty"`
	// Tags - Resource tags
	Tags *map[string]*string `json:"tags,omitempty"`
	// Kind - The kind of web test that this web test watches. Choices are ping and multistep. Possible values include: 'Ping', 'Multistep'
	Kind WebTestKind `json:"kind,omitempty"`
	// WebTestProperties - Metadata describing a web test for an Azure resource.
	*WebTestProperties `json:"properties,omitempty"`
}

// UnmarshalJSON is the custom unmarshaler for WebTest struct.
func (wt *WebTest) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	var v *json.RawMessage

	v = m["kind"]
	if v != nil {
		var kind WebTestKind
		err = json.Unmarshal(*m["kind"], &kind)
		if err != nil {
			return err
		}
		wt.Kind = kind
	}

	v = m["properties"]
	if v != nil {
		var properties WebTestProperties
		err = json.Unmarshal(*m["properties"], &properties)
		if err != nil {
			return err
		}
		wt.WebTestProperties = &properties
	}

	v = m["id"]
	if v != nil {
		var ID string
		err = json.Unmarshal(*m["id"], &ID)
		if err != nil {
			return err
		}
		wt.ID = &ID
	}

	v = m["name"]
	if v != nil {
		var name string
		err = json.Unmarshal(*m["name"], &name)
		if err != nil {
			return err
		}
		wt.Name = &name
	}

	v = m["type"]
	if v != nil {
		var typeVar string
		err = json.Unmarshal(*m["type"], &typeVar)
		if err != nil {
			return err
		}
		wt.Type = &typeVar
	}

	v = m["location"]
	if v != nil {
		var location string
		err = json.Unmarshal(*m["location"], &location)
		if err != nil {
			return err
		}
		wt.Location = &location
	}

	v = m["tags"]
	if v != nil {
		var tags map[string]*string
		err = json.Unmarshal(*m["tags"], &tags)
		if err != nil {
			return err
		}
		wt.Tags = &tags
	}

	return nil
}

// WebTestGeolocation geo-physical location to run a web test from. You must specify one or more locations for the test
// to run from.
type WebTestGeolocation struct {
	// Location - Location ID for the webtest to run from.
	Location *string `json:"Id,omitempty"`
}

// WebTestListResult a list of 0 or more Application Insights web test definitions.
type WebTestListResult struct {
	autorest.Response `json:"-"`
	// Value - Set of Application Insights web test definitions.
	Value *[]WebTest `json:"value,omitempty"`
	// NextLink - The link to get the next part of the returned list of web tests, should the return set be too large for a single request. May be null.
	NextLink *string `json:"nextLink,omitempty"`
}

// WebTestListResultIterator provides access to a complete listing of WebTest values.
type WebTestListResultIterator struct {
	i    int
	page WebTestListResultPage
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *WebTestListResultIterator) Next() error {
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err := iter.page.Next()
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter WebTestListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter WebTestListResultIterator) Response() WebTestListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter WebTestListResultIterator) Value() WebTest {
	if !iter.page.NotDone() {
		return WebTest{}
	}
	return iter.page.Values()[iter.i]
}

// IsEmpty returns true if the ListResult contains no values.
func (wtlr WebTestListResult) IsEmpty() bool {
	return wtlr.Value == nil || len(*wtlr.Value) == 0
}

// webTestListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (wtlr WebTestListResult) webTestListResultPreparer() (*http.Request, error) {
	if wtlr.NextLink == nil || len(to.String(wtlr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(wtlr.NextLink)))
}

// WebTestListResultPage contains a page of WebTest values.
type WebTestListResultPage struct {
	fn   func(WebTestListResult) (WebTestListResult, error)
	wtlr WebTestListResult
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *WebTestListResultPage) Next() error {
	next, err := page.fn(page.wtlr)
	if err != nil {
		return err
	}
	page.wtlr = next
	return nil
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page WebTestListResultPage) NotDone() bool {
	return !page.wtlr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page WebTestListResultPage) Response() WebTestListResult {
	return page.wtlr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page WebTestListResultPage) Values() []WebTest {
	if page.wtlr.IsEmpty() {
		return nil
	}
	return *page.wtlr.Value
}

// WebTestProperties metadata describing a web test for an Azure resource.
type WebTestProperties struct {
	// SyntheticMonitorID - Unique ID of this WebTest. This is typically the same value as the Name field.
	SyntheticMonitorID *string `json:"SyntheticMonitorId,omitempty"`
	// WebTestName - User defined name if this WebTest.
	WebTestName *string `json:"Name,omitempty"`
	// Description - Purpose/user defined descriptive test for this WebTest.
	Description *string `json:"Description,omitempty"`
	// Enabled - Is the test actively being monitored.
	Enabled *bool `json:"Enabled,omitempty"`
	// Frequency - Interval in seconds between test runs for this WebTest. Default value is 300.
	Frequency *int32 `json:"Frequency,omitempty"`
	// Timeout - Seconds until this WebTest will timeout and fail. Default value is 30.
	Timeout *int32 `json:"Timeout,omitempty"`
	// WebTestKind - The kind of web test this is, valid choices are ping and multistep. Possible values include: 'Ping', 'Multistep'
	WebTestKind WebTestKind `json:"Kind,omitempty"`
	// RetryEnabled - Allow for retries should this WebTest fail.
	RetryEnabled *bool `json:"RetryEnabled,omitempty"`
	// Locations - A list of where to physically run the tests from to give global coverage for accessibility of your application.
	Locations *[]WebTestGeolocation `json:"Locations,omitempty"`
	// Configuration - An XML configuration specification for a WebTest.
	Configuration *WebTestPropertiesConfiguration `json:"Configuration,omitempty"`
	// ProvisioningState - Current state of this component, whether or not is has been provisioned within the resource group it is defined. Users cannot change this value but are able to read from it. Values will include Succeeded, Deploying, Canceled, and Failed.
	ProvisioningState *string `json:"provisioningState,omitempty"`
}

// WebTestPropertiesConfiguration an XML configuration specification for a WebTest.
type WebTestPropertiesConfiguration struct {
	// WebTest - The XML specification of a WebTest to run against an application.
	WebTest *string `json:"WebTest,omitempty"`
}