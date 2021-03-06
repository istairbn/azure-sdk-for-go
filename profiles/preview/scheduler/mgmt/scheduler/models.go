// +build go1.9

// Copyright 2018 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package scheduler

import original "github.com/Azure/azure-sdk-for-go/services/scheduler/mgmt/2016-03-01/scheduler"

type DayOfWeek = original.DayOfWeek

const (
	Friday    DayOfWeek = original.Friday
	Monday    DayOfWeek = original.Monday
	Saturday  DayOfWeek = original.Saturday
	Sunday    DayOfWeek = original.Sunday
	Thursday  DayOfWeek = original.Thursday
	Tuesday   DayOfWeek = original.Tuesday
	Wednesday DayOfWeek = original.Wednesday
)

func PossibleDayOfWeekValues() []DayOfWeek {
	return original.PossibleDayOfWeekValues()
}

type JobActionType = original.JobActionType

const (
	HTTP            JobActionType = original.HTTP
	HTTPS           JobActionType = original.HTTPS
	ServiceBusQueue JobActionType = original.ServiceBusQueue
	ServiceBusTopic JobActionType = original.ServiceBusTopic
	StorageQueue    JobActionType = original.StorageQueue
)

func PossibleJobActionTypeValues() []JobActionType {
	return original.PossibleJobActionTypeValues()
}

type JobCollectionState = original.JobCollectionState

const (
	Deleted   JobCollectionState = original.Deleted
	Disabled  JobCollectionState = original.Disabled
	Enabled   JobCollectionState = original.Enabled
	Suspended JobCollectionState = original.Suspended
)

func PossibleJobCollectionStateValues() []JobCollectionState {
	return original.PossibleJobCollectionStateValues()
}

type JobExecutionStatus = original.JobExecutionStatus

const (
	Completed JobExecutionStatus = original.Completed
	Failed    JobExecutionStatus = original.Failed
	Postponed JobExecutionStatus = original.Postponed
)

func PossibleJobExecutionStatusValues() []JobExecutionStatus {
	return original.PossibleJobExecutionStatusValues()
}

type JobHistoryActionName = original.JobHistoryActionName

const (
	ErrorAction JobHistoryActionName = original.ErrorAction
	MainAction  JobHistoryActionName = original.MainAction
)

func PossibleJobHistoryActionNameValues() []JobHistoryActionName {
	return original.PossibleJobHistoryActionNameValues()
}

type JobScheduleDay = original.JobScheduleDay

const (
	JobScheduleDayFriday    JobScheduleDay = original.JobScheduleDayFriday
	JobScheduleDayMonday    JobScheduleDay = original.JobScheduleDayMonday
	JobScheduleDaySaturday  JobScheduleDay = original.JobScheduleDaySaturday
	JobScheduleDaySunday    JobScheduleDay = original.JobScheduleDaySunday
	JobScheduleDayThursday  JobScheduleDay = original.JobScheduleDayThursday
	JobScheduleDayTuesday   JobScheduleDay = original.JobScheduleDayTuesday
	JobScheduleDayWednesday JobScheduleDay = original.JobScheduleDayWednesday
)

func PossibleJobScheduleDayValues() []JobScheduleDay {
	return original.PossibleJobScheduleDayValues()
}

type JobState = original.JobState

const (
	JobStateCompleted JobState = original.JobStateCompleted
	JobStateDisabled  JobState = original.JobStateDisabled
	JobStateEnabled   JobState = original.JobStateEnabled
	JobStateFaulted   JobState = original.JobStateFaulted
)

func PossibleJobStateValues() []JobState {
	return original.PossibleJobStateValues()
}

type RecurrenceFrequency = original.RecurrenceFrequency

const (
	Day    RecurrenceFrequency = original.Day
	Hour   RecurrenceFrequency = original.Hour
	Minute RecurrenceFrequency = original.Minute
	Month  RecurrenceFrequency = original.Month
	Week   RecurrenceFrequency = original.Week
)

func PossibleRecurrenceFrequencyValues() []RecurrenceFrequency {
	return original.PossibleRecurrenceFrequencyValues()
}

type RetryType = original.RetryType

const (
	Fixed RetryType = original.Fixed
	None  RetryType = original.None
)

func PossibleRetryTypeValues() []RetryType {
	return original.PossibleRetryTypeValues()
}

type ServiceBusAuthenticationType = original.ServiceBusAuthenticationType

const (
	NotSpecified    ServiceBusAuthenticationType = original.NotSpecified
	SharedAccessKey ServiceBusAuthenticationType = original.SharedAccessKey
)

func PossibleServiceBusAuthenticationTypeValues() []ServiceBusAuthenticationType {
	return original.PossibleServiceBusAuthenticationTypeValues()
}

type ServiceBusTransportType = original.ServiceBusTransportType

const (
	ServiceBusTransportTypeAMQP         ServiceBusTransportType = original.ServiceBusTransportTypeAMQP
	ServiceBusTransportTypeNetMessaging ServiceBusTransportType = original.ServiceBusTransportTypeNetMessaging
	ServiceBusTransportTypeNotSpecified ServiceBusTransportType = original.ServiceBusTransportTypeNotSpecified
)

func PossibleServiceBusTransportTypeValues() []ServiceBusTransportType {
	return original.PossibleServiceBusTransportTypeValues()
}

type SkuDefinition = original.SkuDefinition

const (
	Free       SkuDefinition = original.Free
	P10Premium SkuDefinition = original.P10Premium
	P20Premium SkuDefinition = original.P20Premium
	Standard   SkuDefinition = original.Standard
)

func PossibleSkuDefinitionValues() []SkuDefinition {
	return original.PossibleSkuDefinitionValues()
}

type Type = original.Type

const (
	TypeActiveDirectoryOAuth Type = original.TypeActiveDirectoryOAuth
	TypeBasic                Type = original.TypeBasic
	TypeClientCertificate    Type = original.TypeClientCertificate
	TypeHTTPAuthentication   Type = original.TypeHTTPAuthentication
)

func PossibleTypeValues() []Type {
	return original.PossibleTypeValues()
}

type BasicAuthentication = original.BasicAuthentication
type ClientCertAuthentication = original.ClientCertAuthentication
type BasicHTTPAuthentication = original.BasicHTTPAuthentication
type HTTPAuthentication = original.HTTPAuthentication
type HTTPRequest = original.HTTPRequest
type JobAction = original.JobAction
type JobCollectionDefinition = original.JobCollectionDefinition
type JobCollectionListResult = original.JobCollectionListResult
type JobCollectionListResultIterator = original.JobCollectionListResultIterator
type JobCollectionListResultPage = original.JobCollectionListResultPage
type JobCollectionProperties = original.JobCollectionProperties
type JobCollectionQuota = original.JobCollectionQuota
type JobCollectionsDeleteFuture = original.JobCollectionsDeleteFuture
type JobCollectionsDisableFuture = original.JobCollectionsDisableFuture
type JobCollectionsEnableFuture = original.JobCollectionsEnableFuture
type JobDefinition = original.JobDefinition
type JobErrorAction = original.JobErrorAction
type JobHistoryDefinition = original.JobHistoryDefinition
type JobHistoryDefinitionProperties = original.JobHistoryDefinitionProperties
type JobHistoryFilter = original.JobHistoryFilter
type JobHistoryListResult = original.JobHistoryListResult
type JobHistoryListResultIterator = original.JobHistoryListResultIterator
type JobHistoryListResultPage = original.JobHistoryListResultPage
type JobListResult = original.JobListResult
type JobListResultIterator = original.JobListResultIterator
type JobListResultPage = original.JobListResultPage
type JobMaxRecurrence = original.JobMaxRecurrence
type JobProperties = original.JobProperties
type JobRecurrence = original.JobRecurrence
type JobRecurrenceSchedule = original.JobRecurrenceSchedule
type JobRecurrenceScheduleMonthlyOccurrence = original.JobRecurrenceScheduleMonthlyOccurrence
type JobStateFilter = original.JobStateFilter
type JobStatus = original.JobStatus
type OAuthAuthentication = original.OAuthAuthentication
type RetryPolicy = original.RetryPolicy
type ServiceBusAuthentication = original.ServiceBusAuthentication
type ServiceBusBrokeredMessageProperties = original.ServiceBusBrokeredMessageProperties
type ServiceBusMessage = original.ServiceBusMessage
type ServiceBusQueueMessage = original.ServiceBusQueueMessage
type ServiceBusTopicMessage = original.ServiceBusTopicMessage
type Sku = original.Sku
type StorageQueueMessage = original.StorageQueueMessage

func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}

type JobCollectionsClient = original.JobCollectionsClient

func NewJobCollectionsClient(subscriptionID string) JobCollectionsClient {
	return original.NewJobCollectionsClient(subscriptionID)
}
func NewJobCollectionsClientWithBaseURI(baseURI string, subscriptionID string) JobCollectionsClient {
	return original.NewJobCollectionsClientWithBaseURI(baseURI, subscriptionID)
}

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}

type JobsClient = original.JobsClient

func NewJobsClient(subscriptionID string) JobsClient {
	return original.NewJobsClient(subscriptionID)
}
func NewJobsClientWithBaseURI(baseURI string, subscriptionID string) JobsClient {
	return original.NewJobsClientWithBaseURI(baseURI, subscriptionID)
}
